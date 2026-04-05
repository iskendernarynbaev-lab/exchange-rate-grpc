package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	RequestsTotal *prometheus.CounterVec
	Duration      *prometheus.HistogramVec
	StoreErrors   prometheus.Counter
}

func New(reg prometheus.Registerer) *Metrics {
	m := &Metrics{
		RequestsTotal: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "grinex",
			Subsystem: "rates",
			Name:      "requests_total",
			Help:      "Total GetRates requests by status.",
		}, []string{"status"}),
		Duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "grinex",
			Subsystem: "rates",
			Name:      "request_duration_seconds",
			Help:      "Duration of GetRates requests.",
			Buckets:   prometheus.DefBuckets,
		}, []string{"status"}),
		StoreErrors: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "grinex",
			Subsystem: "storage",
			Name:      "store_errors_total",
			Help:      "Total number of repository store errors.",
		}),
	}

	reg.MustRegister(m.RequestsTotal, m.Duration, m.StoreErrors)
	return m
}
