package middlewares

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {
		logrus.WithField("middleware", "Request").Infof(
			"[%s] \"%s %s %s\"",
			req.RemoteAddr,
			req.Method,
			req.URL,
			req.Proto,
		)

		next.ServeHTTP(w, req)
	})
}
