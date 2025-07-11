package middlewares

import (
	"net/http"
	"time"

	"github.com/go-redis/redis"

	"github.com/natanfds/vtt_odisseia/configs"
	"github.com/natanfds/vtt_odisseia/services"
)

func RateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path
		ip := r.RemoteAddr
		key := configs.ROOT_KEY_REDIS_RATE_LIMIT + ":" + route + ":" + ip
		route_limits := configs.ROUTE_LIMITS()

		if amount, exists := route_limits[route]; !exists {
			next.ServeHTTP(w, r)
			return
		} else {
			_, err := services.RedisService.Get(key)

			if err == redis.Nil {
				err = services.RedisService.Set(key, "", time.Second*time.Duration(amount))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				next.ServeHTTP(w, r)
				return
			} else if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
		}
	})
}
