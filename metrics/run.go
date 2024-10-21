package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vaticyai/syncpocalypse/logger"
)

var (
	// TotalBackups Counts all requests arrived to the API Gateway.
	TotalBackups = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "syncpocalypse_total_backups",
			Help: "Total number of backups executed",
		},
		[]string{"status"},
	)
)

// Run Starts serving the Prometheus metrics.
//
// Example:
//
//	go Run()
func Run() {
	logger.Logger.Debug("Exporting metrics on port '9090' at route: '/metrics'")

	prometheus.MustRegister(TotalBackups)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)
}
