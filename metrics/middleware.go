package metrics

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func (m *Metrics) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, req.ProtoMajor)
		next.ServeHTTP(ww, req)

		m.IncRequest(ww.Status(), req.Method, req.URL.Path)
		m.ObserveLatency(start, ww.Status(), req.Method, req.URL.Path)
	})
}
