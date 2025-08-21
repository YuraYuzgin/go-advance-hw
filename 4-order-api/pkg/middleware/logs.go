package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)

		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.WithFields(logrus.Fields{
			"path":       r.URL.Path,
			"StatusCode": wrapper.StatusCode,
		}).Info("HTTP request")
	})
}
