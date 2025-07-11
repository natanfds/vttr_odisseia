package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/natanfds/vtt_odisseia/configs"
	"github.com/natanfds/vtt_odisseia/handlers"
	"github.com/natanfds/vtt_odisseia/middlewares"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(configs.ERR_LOAD_ENV, err)
		return
	}

	err = configs.ENV.Load()
	if err != nil {
		fmt.Println(configs.ERR_LOAD_ENV, err)
		return
	}

	db, err := services.StartDatabase()

	if err != nil {
		fmt.Println(configs.ERR_START_DB, err)
		return
	}
	err = services.RedisService.Start()
	if err != nil {
		fmt.Println(configs.ERR_START_CACHE, err)
		return
	}

	repositories.InitRepositories(db)
	port := configs.ENV.ApiPort

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})
	mux.HandleFunc("/account", handlers.CreateAccountHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)

	fmt.Println(configs.MSG_START_SERVER, port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: middlewares.RateLimiterMiddleware(mux),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(configs.ERR_START_SERVER, err)
		return
	}
}
