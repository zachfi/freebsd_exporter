package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	nfsServerOperationsGetattr = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "nfs_server_operations_getattr",
	}, []string{})
)
