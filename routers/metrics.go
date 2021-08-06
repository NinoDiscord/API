package routers

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"nino.sh/api/metrics"
)

func NewMetricsRouter(m *metrics.Metrics) chi.Router {
	router := chi.NewRouter()
	router.Use(m.Middleware)
	router.Mount("/", promhttp.Handler())

	return router
}
