package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/amaraliou/trackr-core/internal/auth"
	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/amaraliou/trackr-core/internal/response"
	"github.com/amaraliou/trackr-core/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// Login -> handles POST /api/v1/auth/login
func (handler *Handler) Login(writer http.ResponseWriter, request *http.Request) {

	log := handler.logger
	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("login")
	if err != nil {
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := handler.SignIn(user.Email, user.Password)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log.Infof("Successfully logged in.")
	response.JSON(writer, http.StatusOK, token)
}

// SignIn -> retrieves user JWT token given username and password
func (handler *Handler) SignIn(email, password string) (string, error) {

	pgRepo := handler.pgRepo

	var err error
	user, err := pgRepo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = model.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}
