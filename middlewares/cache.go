package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func CacheMiddleware(tll time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.String()
		fmt.Println(key)
		next.ServeHTTP(w, r)
	})
}
