package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/amaraliou/trackr-core/internal/response"
	"github.com/amaraliou/trackr-core/pkg/logger"
)

// CreateApplication ...
func (handler *Handler) CreateApplication(writer http.ResponseWriter, request *http.Request) {

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

	application := model.Application{}
	err = json.Unmarshal(body, &application)
	if err != nil {
		log.Warnf("Couldn't marshal JSON body")
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	err = application.Validate("create")
	if err != nil {
		log.Warnf(err.Error())
		response.ERROR(writer, http.StatusUnprocessableEntity, err)
		return
	}

	applicationCreated, err := pgRepo.CreateApplication(application)
	if err != nil {
		response.ERROR(writer, http.StatusInternalServerError, err)
		return
	}

	log.Infof("Successfully created application.")
	writer.Header().Set("Location", fmt.Sprintf("%s%s/%s", request.Host, request.RequestURI, applicationCreated.ID.String()))
	response.JSON(writer, http.StatusCreated, map[string]interface{}{"application": applicationCreated})
}
