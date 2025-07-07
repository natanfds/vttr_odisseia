package main

import (
	"fmt"
	"net/http"

	"github.com/natanfds/vtt_odisseia/handlers"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/services"
)

func main() {
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

	fmt.Println("Servidor iniciado em http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error at start server:", err)
		return
	}
}
