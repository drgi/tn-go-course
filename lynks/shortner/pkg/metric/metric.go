package metric

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Metric struct {
	reqCount prometheus.Counter
	reqTime  *prometheus.HistogramVec
}

func New() *Metric {
	reqCount := promauto.NewCounter(prometheus.CounterOpts{
		Name: "total_request",
		Help: "The total number of http request",
	})

	reqTime := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "request_duration_seconds",
		Help:    "Request duration distribution",
		Buckets: []float64{0.25, 0.5, 1, 1.5, 2, 3, 4, 5, 7.5, 10, 20},
	},
		[]string{"endpoint"})

	return &Metric{
		reqCount: reqCount,
		reqTime:  reqTime,
	}
}

func (m *Metric) Handler() http.Handler {
	return promhttp.Handler()
}

func (m *Metric) IncrementRequest(url string, startTime time.Time) {
	m.reqCount.Inc()
	m.reqTime.WithLabelValues(url).Observe(float64(time.Since(startTime).Seconds()))
}
