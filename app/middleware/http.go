package middleware

import (
	"net"
	"net/http"

	"github.com/shashankbiet/rate-limiter/pkg/logger"
)

func HttpRequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		logger.Info("HTTP Request", "ip", ip, "url", r.URL, "body", r.Body)
		next.ServeHTTP(w, r)
	})
}
