package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var start time.Time = time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_ip", r.RemoteAddr,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}
