package process

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	Namespace      = "metris"
	ErrorCountName = "error_count"
)

var (
	kebErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: "keb",
			Name:      ErrorCountName,
			Help:      "Number of continuous errors of getting the list with runtimes from KEB since last success.",
		},
		[]string{"error_count"},
	)

	cacheErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: "cache",
			Name:      ErrorCountName,
			Help:      "Number of continuous errors of adding the subaccount to the cache since last success.",
		},
		[]string{"error_rate"},
	)

	skrErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: "skr",
			Name:      ErrorCountName,
			Help:      "Number of continuous errors of getting the metrics of the cluster since last success.",
		},
		[]string{"error_count"},
	)

	gardenerErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: "gardener",
			Name:      ErrorCountName,
			Help:      "Number of continuous errors of getting the config of the cluster since last success.",
		},
		[]string{"error_count"},
	)

	edpErrorCount = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: "edp",
			Name:      ErrorCountName,
			Help:      "Number of continuous errors of sending the metrics to EDP since last success.",
		},
		[]string{"error_count"},
	)

	clustersScraped = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: "keb",
			Name:      "number_clusters_scraped",
			Help:      "Number of clusters scraped.",
		},
		[]string{"number_clusters_scraped"},
	)

	runtimesRateChanged = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: "keb",
			Name:      "rate_changed",
			Help:      "Ratio between number of current clusters scraped to previous number of clusters scraped.",
		},
		[]string{"rate_changed"},
	)
)
