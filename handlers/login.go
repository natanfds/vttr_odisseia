package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/natanfds/vtt_odisseia/configs"
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
		w.Write([]byte(configs.MSG_INVALID_BODY))
		return
	}

	if err := utils.Validate.Struct(loginData); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(configs.MSG_INVALID_BODY))
		return
	}

	searchData := dtos.GetUser{
		Username: loginData.Username,
	}
	user, err := repositories.UserRepository.GetUser(searchData)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(configs.MSG_USER_NOT_FOUND))
		return
	}

	//Pegando token do banco se tiver
	userID := strconv.Itoa(int(user.ID))
	dbToken, err := repositories.AuthTokenRepository.GetTokenByID(userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.MSG_INTERNAL_ERROR))
		return
	}

	willUpdateToken := false
	if dbToken != "" {
		_, err := utils.ValidateJWT(dbToken)
		if err == nil {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(dbToken))
			return
		} else {
			willUpdateToken = true
		}
	}

	generatedToken, err := utils.GenerateJWT(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(configs.MSG_INTERNAL_ERROR))
		return
	}

	if willUpdateToken {
		err = repositories.AuthTokenRepository.UpdateToken(generatedToken, user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(configs.MSG_INTERNAL_ERROR))
			return
		}
	} else {
		err = repositories.AuthTokenRepository.CreateToken(user, generatedToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(configs.MSG_INTERNAL_ERROR))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(generatedToken))
}
