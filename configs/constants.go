package configs

import "time"

const (
	HEADER_AUTH         = "Authorization"
	ROOT_KEY_REDIS_AUTH = "auth"
	AUTH_REDIS_DURATION = 8 * time.Hour
	DB_ADDRESS          = "test.db"
	REDIS_ADDR          = "localhost:6379"
	REDIS_PASS          = ""
	REDIS_DB            = 0
	MSG_INVALID_BODY    = "Invalid request body!"
	MSG_USER_NOT_FOUND  = "User not found!"
	MSG_INTERNAL_ERROR  = "Internal server error!"
	JWT_SECRET          = "secret"
	TOKEN_EXPIRATION    = 7 * 24 * time.Hour
)
