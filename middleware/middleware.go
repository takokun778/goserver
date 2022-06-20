package middleware

import (
	"net/http"
	"time"

	"server/logger"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		cors(w)

		r = r.WithContext(logger.SetLogCtx(r.Context()))

		lrw := &LoggingResponseWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		log := logger.GetLogCtx(r.Context())

		if r.Method == http.MethodOptions {
			elapsed := time.Since(start)

			log.Sugar().Infof("%s %s %d %s", r.Method, r.URL, lrw.StatusCode, elapsed)
		}

		next.ServeHTTP(lrw, r)

		elapsed := time.Since(start)

		log.Sugar().Infof("%s %s %d %s", r.Method, r.URL, lrw.StatusCode, elapsed)
	})
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
}
