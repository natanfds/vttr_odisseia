package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/natanfds/vtt_odisseia/dtos"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/utils"
)

func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	var loginData dtos.CreateUser

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body!"))
		return
	}

	if err := utils.Validate.Struct(loginData); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Invalid request body!"))
		return
	}

	err := repositories.UserRepository.CreateUser(loginData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account created successfully!"))
}
