package middleware

import (
	"net/http"
	"time"

	"log/slog"

	"github.com/go-chi/chi/v5/middleware"
)

// logger middleware for access logs
func Logger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// gathers metrics from the upstream handlers
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			st := time.Now()

			next.ServeHTTP(ww, r)

			et := time.Since(st)

			//print log and metrics
			logger.Info(
				"handled request",
				slog.String("method", r.Method),
				slog.String("uri", r.URL.RequestURI()),
				slog.String("user_agent", r.Header.Get("User-Agent")),
				slog.String("ip", r.RemoteAddr),
				slog.Int("code", ww.Status()),
				slog.Int("bytes", ww.BytesWritten()),
				slog.String("request_time_pretty", et.String()),
				slog.Int64("request_time_ns", et.Nanoseconds()),
			)
		})
	}
}
