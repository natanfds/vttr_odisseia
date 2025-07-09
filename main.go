package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/natanfds/vtt_odisseia/configs"
	"github.com/natanfds/vtt_odisseia/handlers"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error at load env:", err)
		return
	}

	err = configs.ENV.Load()
	if err != nil {
		fmt.Println("Error at load env:", err)
		return
	}

	db, err := services.StartDatabase()

	if err != nil {
		fmt.Println("Error at start database:", err)
		return
	}
	err = services.CacheService.Start()
	if err != nil {
		fmt.Println("Error ar start cache", err)
		return
	}

	repositories.InitRepositories(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})
	http.HandleFunc("/account", handlers.CreateAccountHandler)
	http.HandleFunc("/login", handlers.LoginHandler)

	port := configs.ENV.ApiPort
	fmt.Println("Servidor iniciado em http://localhost:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error at start server:", err)
		return
	}
}
