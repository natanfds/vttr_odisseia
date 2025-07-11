package handlers

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v5"

	"github.com/natanfds/vtt_odisseia/configs"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/services"
	"github.com/natanfds/vtt_odisseia/utils"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	tokenHeader := r.Header.Get(configs.HEADER_AUTH)
	tokenData, _ := utils.ValidateJWT(tokenHeader)
	userID := tokenData.Claims.(jwt.MapClaims)["user_id"].(string)
	cacheKey := configs.ROOT_KEY_REDIS_AUTH + ":" + userID

	if tokenHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err := services.RedisService.Delete(cacheKey)
	if err == redis.Nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = repositories.AuthTokenRepository.DeleteToken(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
