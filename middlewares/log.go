package middlewares

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, req.ProtoMajor)

		logrus.WithField("middleware", "Request").Info(
			"[%s] \"%s %s / %s\" => %d",
			req.RemoteAddr,
			req.Method,
			req.URL,
			req.Proto,
			ww.Status(),
		)

		next.ServeHTTP(w, req)
	})
}
