package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const metricPrefix = "mongoshake"

var (
	LogsGet = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_logs_get",
		Help: "Number of logs (get)",
	}, []string{"alias", "url"})

	LogsRepl = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_logs_repl",
		Help: "Number of logs (repl)",
	}, []string{"alias", "url"})

	LogsSuccess = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_logs_success",
		Help: "Number of successful logs",
	}, []string{"alias", "url"})

	TPS = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_tps",
		Help: "Transactions per second",
	}, []string{"alias", "url"})

	ReplicationLatency = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_replication_latency",
		Help: "Replication latency in seconds",
	}, []string{"alias", "url"})

	Up = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: metricPrefix + "_up",
		Help: "Up",
	}, []string{"alias", "url"})
)
