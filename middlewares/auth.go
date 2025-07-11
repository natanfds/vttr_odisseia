package middlewares

import (
	"net/http"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/natanfds/vtt_odisseia/configs"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/services"
	"github.com/natanfds/vtt_odisseia/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path
		if slices.Contains(configs.NON_AUTH_ROUTES(), route) {
			next.ServeHTTP(w, r)
			return
		}
		tokenHeader := r.Header.Get(configs.HEADER_AUTH)
		tokenData, err := utils.ValidateJWT(tokenHeader)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		userID := tokenData.Claims.(jwt.MapClaims)["user_id"].(string)
		expirationInt := tokenData.Claims.(jwt.MapClaims)["exp"].(float64)
		cacheKey := configs.ROOT_KEY_REDIS_AUTH + ":" + userID
		if expirationInt < float64(time.Now().Unix()) {
			services.RedisService.Delete(cacheKey)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenCache, err := services.RedisService.Get(cacheKey)
		if err == nil {
			if tokenCache == tokenHeader {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenDb, err := repositories.AuthTokenRepository.GetTokenByID(userID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if tokenDb != tokenHeader {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		authRedisPersistence := time.Duration(configs.ENV.AuthRedisDurationHour) * time.Hour
		services.RedisService.Set(cacheKey, tokenHeader, authRedisPersistence)

		next.ServeHTTP(w, r)
	})
}
