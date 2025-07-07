package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/natanfds/vtt_odisseia/dtos"
	"github.com/natanfds/vtt_odisseia/repositories"
	"github.com/natanfds/vtt_odisseia/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!"))
		return
	}

	var loginData dtos.Login

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

	searchData := dtos.GetUser{
		Username: loginData.Username,
	}
	user, err := repositories.UserRepository.GetUser(searchData)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found!"))
		return
	}

	id := strconv.Itoa(int(user.ID))
	generatedToken, err := utils.GenerateJWT(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error!"))
		return
	}

	err = repositories.AuthTokenRepository.CreateToken(user, generatedToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(generatedToken))
}
