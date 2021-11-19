package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xaque208/freebsd_exporter/pkg/nfs"
	"github.com/xaque208/freebsd_exporter/pkg/poudriere"
	"github.com/xaque208/znet/pkg/util"
)

var rootCmd = &cobra.Command{
	Use:   "freebsd_exporter",
	Short: "Export FreeBSD stats to Pometheus",
	Long:  "",
	Run:   run,
}

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
	logger := util.NewLogger()

	nfsExporter, err := nfs.NewExporter(logger)
	if err == nil {
		prometheus.MustRegister(nfsExporter)
	}

	poudriereExporter, err := poudriere.NewExporter(logger)
	if err == nil {
		prometheus.MustRegister(poudriereExporter)
	}

	http.Handle("/metrics", promhttp.Handler())
	_ = level.Error(logger).Log(http.ListenAndServe(":8080", nil))
}
