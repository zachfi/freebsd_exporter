package cmd

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xaque208/freebsd_exporter/exporter"
)

var rootCmd = &cobra.Command{
	Use:   "freebsd_exporter",
	Short: "Export FreeBSD stats to Pometheus",
	Long:  "",
	Run:   run,
}

// nfsstat -E --libxo=json

var (
	verbose       bool
	listenAddress string
	interval      int
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Increase verbosity")
	rootCmd.PersistentFlags().StringVarP(&listenAddress, "listen", "L", ":9100", "The listen address (default is :9100")
	rootCmd.PersistentFlags().IntVarP(&interval, "interval", "i", 30, "The interval at which to update the data")

	viper.SetDefault("interval", 30)
}

// initConfig reads in the config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv()
}

func run(cmd *cobra.Command, args []string) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.WithFields(log.Fields{
		"url": listenAddress,
	}).Info("starting metrics listener")
	go exporter.StartMetricsServer(listenAddress)

	tick := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-tick.C:
			log.Debugf("scraping exporter")
			err := exporter.Scrape()
			if err != nil {
				log.Error(err)
			}
		}
	}
}
