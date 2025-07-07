package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func cacheMiddleware(w http.ResponseWriter, r *http.Request) {
	key := r.URL.String()
	fmt.Println(key)
}

func CacheMiddleware(tll time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(cacheMiddleware)
}
