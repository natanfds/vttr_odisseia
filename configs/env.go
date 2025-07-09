package configs

import "github.com/caarlos0/env/v11"

type envStruct struct {
	JwtSecret             string `env:"JWT_SECRET"`
	TokenExpirationDays   int    `env:"TOKEN_EXPIRATION_DAYS"`
	AuthRedisDurationHour int    `env:"AUTH_REDIS_DURATION_HOUR"`
	DbFilePath            string `env:"DB_FILE_PATH"`
	RedisAddr             string `env:"REDIS_ADDR"`
	RedisPass             string `env:"REDIS_PASS"`
	RedisDb               int    `env:"REDIS_DB"`
	ApiPort               string `env:"API_PORT"`
}

func (e *envStruct) Load() error {
	cfg, err := env.ParseAs[envStruct]()
	if err != nil {
		return err
	}
	ENV = cfg
	return nil
}

var ENV envStruct = envStruct{}
