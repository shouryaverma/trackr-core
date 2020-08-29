package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/amaraliou/trackr-core/internal/response"
	"github.com/amaraliou/trackr-core/pkg/logger"
	"github.com/go-chi/chi"
)

// CreateUser ...
func (handler *Handler) CreateUser(writer http.ResponseWriter, request *http.Request) {

	pgRepo := handler.pgRepo
	log := handler.logger
	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Warnf("Couldn't read request body")
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Warnf("Couldn't marshal JSON body")
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = user.Validate("create")
	if err != nil {
		log.Warnf(err.Error())
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := pgRepo.CreateUser(user)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log.Infof("Successfully created user.")
	writer.Header().Set("Location", fmt.Sprintf("%s%s/%s", request.Host, request.RequestURI, userCreated.ID.String()))
	response.JSON(writer, http.StatusCreated, map[string]interface{}{"user": userCreated})
}

// GetAllUsers ...
func (handler *Handler) GetAllUsers(writer http.ResponseWriter, request *http.Request) {

	pgRepo := handler.pgRepo
	log := handler.logger

	// Check role/token

	users, err := pgRepo.AllUsers()
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})
	log.Infof("Successfully retrieved all users")
	response.JSON(writer, http.StatusOK, map[string]interface{}{"users": users})
}

// GetUser ...
func (handler *Handler) GetUser(writer http.ResponseWriter, request *http.Request) {

	pgRepo := handler.pgRepo
	log := handler.logger

	// Check role/token
	userID := chi.URLParam(request, "id")

	user, err := pgRepo.GetUser(userID)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})
	log.Infof("Successfully retrieved the user")
	response.JSON(writer, http.StatusOK, map[string]interface{}{"user": user})
}

// UpdateUser ...
func (handler *Handler) UpdateUser(writer http.ResponseWriter, request *http.Request) {

	pgRepo := handler.pgRepo
	log := handler.logger

	// Check role/token
	userID := chi.URLParam(request, "id")

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

	updatedUser, err := pgRepo.UpdateUser(user, userID)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})
	log.Infof("Successfully updated the user")
	response.JSON(writer, http.StatusOK, map[string]interface{}{"user": updatedUser})
}

// DeleteUser ...
func (handler *Handler) DeleteUser(writer http.ResponseWriter, request *http.Request) {

	pgRepo := handler.pgRepo
	log := handler.logger

	// Check role/token
	userID := chi.URLParam(request, "id")

	_, err := pgRepo.DeleteUser(userID)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log = log.WithFields(logger.Fields{
		"method": request.Method,
		"host":   request.Host,
		"path":   request.URL.Path,
	})
	log.Infof("Successfully deleted the user")
	writer.Header().Set("Entity", userID)
	response.JSON(writer, http.StatusNoContent, "")
}
