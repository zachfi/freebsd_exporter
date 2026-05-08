// Copyright © 2021 Zach Leslie <code@zleslie.info>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/grafana/dskit/flagext"
	"github.com/prometheus/client_golang/prometheus"
	versioncollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	yaml "gopkg.in/yaml.v2"

	"github.com/zachfi/zkit/pkg/tracing"

	"github.com/zachfi/freebsd_exporter/pkg/nfs"
	"github.com/zachfi/freebsd_exporter/pkg/poudriere"
)

const appName = "freebsd_exporter"

// Build info — set via -ldflags -X main.Version=...
var (
	Version  string
	Branch   string
	Revision string
)

func init() {
	version.Version = Version
	version.Branch = Branch
	version.Revision = Revision
	prometheus.MustRegister(versioncollector.NewCollector(appName))
}

// Config is the top-level config struct, populated from -config.file (YAML)
// then overridden by command-line flags.
type Config struct {
	ListenAddress string         `yaml:"listen_address"`
	LogLevel      string         `yaml:"log_level"`
	Tracing       tracing.Config `yaml:"tracing,omitempty"`
}

// RegisterFlagsAndApplyDefaults wires fields up to f and seeds defaults.
// Called twice during loadConfig: once into a discard FlagSet to find the
// config file path, and once into flag.CommandLine for the real parse.
func (c *Config) RegisterFlagsAndApplyDefaults(prefix string, f *flag.FlagSet) {
	f.StringVar(&c.ListenAddress, "listen", ":9100", "address for the metrics HTTP listener")
	f.StringVar(&c.LogLevel, "log.level", "info", "log level: debug, info, warn, error")

	c.Tracing.RegisterFlagsAndApplyDefaults("tracing", f)
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %v\n", err)
		os.Exit(1)
	}

	level := new(slog.LevelVar)
	if err := setLogLevel(level, cfg.LogLevel); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	logger.Info("starting", "app", appName, "version", Version, "revision", Revision)

	shutdownTracer, err := tracing.InstallOpenTelemetryTracer(&cfg.Tracing, logger, appName, Version)
	if err != nil {
		logger.Error("error initialising tracer", "err", err)
		os.Exit(1)
	}
	defer shutdownTracer()

	// Register collectors. Constructors that fail to initialise (missing
	// binaries, etc.) are skipped with a warning rather than aborting startup
	// so a partial set of metrics is still exposed.
	if exp, err := nfs.NewExporter(logger); err != nil {
		logger.Warn("nfs exporter disabled", "err", err)
	} else {
		prometheus.MustRegister(exp)
	}
	if exp, err := poudriere.NewExporter(logger); err != nil {
		logger.Warn("poudriere exporter disabled", "err", err)
	} else {
		prometheus.MustRegister(exp)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(appName + "\n"))
	})

	srv := &http.Server{
		Addr:              cfg.ListenAddress,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Graceful shutdown on SIGINT/SIGTERM so in-flight scrapes finish.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errCh := make(chan error, 1)
	go func() {
		logger.Info("listening", "addr", cfg.ListenAddress)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			logger.Error("server failed", "err", err)
			os.Exit(1)
		}
	case <-ctx.Done():
		logger.Info("shutting down")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			logger.Error("shutdown error", "err", err)
		}
	}
}

// loadConfig follows the dskit two-pass pattern used by streamgo/weigh: peel
// -config.file off args first, load the YAML, then re-register flags so the
// CLI overrides YAML.
func loadConfig() (*Config, error) {
	const configFileOption = "config.file"

	var configFile string
	args := os.Args[1:]
	cfg := &Config{}

	// First pass: discover -config.file. ContinueOnError + io.Discard keeps
	// flag.Parse from bailing on the unknown app flags.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.StringVar(&configFile, configFileOption, "", "")
	for len(args) > 0 {
		_ = fs.Parse(args)
		args = args[1:]
	}

	cfg.RegisterFlagsAndApplyDefaults("", flag.CommandLine)

	if configFile != "" {
		path, err := filepath.Abs(configFile)
		if err != nil {
			return nil, fmt.Errorf("resolve config path %s: %w", configFile, err)
		}
		buff, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read config %s: %w", path, err)
		}
		if err := yaml.UnmarshalStrict(buff, cfg); err != nil {
			return nil, fmt.Errorf("parse config %s: %w", path, err)
		}
	}

	// Second pass: register the (now-known) -config.file flag so it doesn't
	// trip the real parse, then let CLI flags override YAML.
	flagext.IgnoredFlag(flag.CommandLine, configFileOption, "Configuration file to load")
	flag.Parse()

	return cfg, nil
}

func setLogLevel(lv *slog.LevelVar, name string) error {
	switch strings.ToLower(name) {
	case "debug":
		lv.Set(slog.LevelDebug)
	case "info", "":
		lv.Set(slog.LevelInfo)
	case "warn", "warning":
		lv.Set(slog.LevelWarn)
	case "error":
		lv.Set(slog.LevelError)
	default:
		return fmt.Errorf("unknown log level %q (want debug|info|warn|error)", name)
	}
	return nil
}
