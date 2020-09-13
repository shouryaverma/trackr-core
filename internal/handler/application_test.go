package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amaraliou/trackr-core/internal/model"
	"github.com/amaraliou/trackr-core/internal/storage/mock"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateApplication_201(t *testing.T) {

	applicationToCreate := model.Application{
		JobTitle: "Software Engineer Intern",
		Company:  "GoCardless",
		UserID:   uuid.NewV4(),
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &applicationToCreate,
		IsError:      false,
	}

	jsonByte, err := json.Marshal(applicationToCreate)
	if err != nil {
		t.Error("Failed to marshal Application struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/applications", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/applications' request")
	}

	rr := httptest.NewRecorder()
	createApplicationHandler := http.HandlerFunc(handler.CreateApplication)
	createApplicationHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	createdApplication := responseMap["application"].(map[string]interface{})

	assert.Equal(t, rr.Code, 201)
	assert.Equal(t, createdApplication["job_title"], applicationToCreate.JobTitle)
}

func TestCreateApplication_422_JSON(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.Application{},
		IsError:      false,
	}

	jsonByte, err := json.Marshal("bruh}")
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/applications", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/applications' request")
	}

	rr := httptest.NewRecorder()
	createApplicationHandler := http.HandlerFunc(handler.CreateApplication)
	createApplicationHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 422)
	assert.Equal(t, responseMap["error"], "json: cannot unmarshal string into Go value of type model.Application")
}

func TestCreateApplication_422_Validation(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.Application{},
		IsError:      false,
	}

	cases := []struct {
		inputJSON    string
		errorMessage string
	}{
		{
			inputJSON:    `{"company": "GoCardless"}`,
			errorMessage: "Required Job Title",
		},
		{
			inputJSON:    `{"job_title": "Software Engineer Intern"}`,
			errorMessage: "Required Company",
		},
		{
			inputJSON:    `{"job_title": "Software Engineer Intern", "company": "GoCardless"}`,
			errorMessage: "Required User ID",
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest("POST", "/api/v1/applications", bytes.NewBufferString(c.inputJSON))
		if err != nil {
			t.Error("Failed to create 'POST: /api/v1/applications' request")
		}

		rr := httptest.NewRecorder()
		createApplicationHandler := http.HandlerFunc(handler.CreateApplication)
		createApplicationHandler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, 422)
		assert.Equal(t, responseMap["error"], c.errorMessage)
	}
}

func TestCreateApplication_500(t *testing.T) {

	applicationToCreate := model.Application{
		JobTitle: "Software Engineer Intern",
		Company:  "GoCardless",
		UserID:   uuid.NewV4(),
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.Application{},
		IsError:      true,
		ErrorMessage: "Table 'applications' doesn't exist",
	}

	jsonByte, err := json.Marshal(applicationToCreate)
	if err != nil {
		t.Error("Failed to marshal Application struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/applications", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/applications' request")
	}

	rr := httptest.NewRecorder()
	createApplicationHandler := http.HandlerFunc(handler.CreateApplication)
	createApplicationHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "Table 'applications' doesn't exist")
}
