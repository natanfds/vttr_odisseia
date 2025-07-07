package main

import (
	"fmt"
	"net/http"

	"github.com/natanfds/vtt_odisseia/handlers"
	"github.com/natanfds/vtt_odisseia/models"
	"github.com/natanfds/vtt_odisseia/repositories"
)

func main() {
	db, err := models.StartDatabase()
	if err != nil {
		fmt.Println("Error at start database:", err)
		return
	}
	repositories.InitRepositories(db)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🎪")
	})
	http.HandleFunc("/account", handlers.CreateAccountHandler)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
