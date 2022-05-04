package middleware

import (
	"log"
	"net/http"
	"time"
)

func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		defer accesslog(r, now)

		cors(w)

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func accesslog(r *http.Request, start time.Time) {
	elapsed := time.Since(start)
	log.Printf("%s %s %s\n", r.Method, r.URL, elapsed)
}

func cors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
}
