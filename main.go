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

	serverMux := http.NewServeMux()

	serverMux.HandleFunc(configs.ROUTE_HOME, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})
	serverMux.HandleFunc(configs.ROUTE_ACCOUNT, handlers.CreateAccountHandler)
	serverMux.HandleFunc(configs.ROUTE_LOGIN, handlers.LoginHandler)
	serverMux.HandleFunc(configs.ROUTE_LOGOUT, handlers.LogoutHandler)

	fmt.Println(configs.MSG_START_SERVER, port)

	server := &http.Server{
		Addr: ":" + port,
		Handler: middlewares.ChainMiddlewares(
			serverMux,
			middlewares.RateLimiterMiddleware,
			middlewares.AuthMiddleware,
		),
	}

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(configs.ERR_START_SERVER, err)
		return
	}
}
