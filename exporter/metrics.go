package exporter

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	nfsServerOperationsGetattr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_getattr",
	}, []string{})
)

func init() {
	prometheus.MustRegister(
		nfsServerOperationsGetattr,
	)
}

func StartMetricsServer(bindAddr string) {
	d := http.NewServeMux()
	d.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(bindAddr, d)
	if err != nil {
		log.Fatal("Failed to start metrics server, error is:", err)
	}
}
