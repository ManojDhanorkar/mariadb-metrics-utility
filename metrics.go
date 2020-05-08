package main

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

var (
	MetricMariadbActiveConnections = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "monitoring",
			Name:      "mariadb_active_connections",
			Help:      "Active connections to MariaDB server",
		})
)

func setTestMetrics() {
	MetricMariadbActiveConnections.Add(100)
}

// UpdateMetrics updates all the metrics every n minutes
func UpdateMetrics(min int) {

	d := time.Duration(min) * time.Minute
	for range time.Tick(d) {
		UpdateConnectionsCount()

	}
}

// Update connection count
func UpdateConnectionsCount() {

	count := GetMetricMariadbActiveConnections()
	log.Debug(fmt.Sprintf("Active connections: %d", count))
	MetricMariadbActiveConnections.Add(float64(count))
}
