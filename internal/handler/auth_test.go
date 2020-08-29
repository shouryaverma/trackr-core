package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amaraliou/trackr-v2/internal/model"
	"github.com/amaraliou/trackr-v2/internal/storage/mock"
	"gopkg.in/go-playground/assert.v1"
)

func TestLogin_200(t *testing.T) {

	userLogin := model.User{
		Email:    "random@gmail.com",
		Password: "random",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &userLogin,
		IsError:      false,
	}

	jsonByte, err := json.Marshal(userLogin)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/auth/login' request")
	}

	rr := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(handler.Login)
	loginHandler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)
}

func TestLogin_422_JSON(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      false,
	}

	jsonByte, err := json.Marshal("bruh}")
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/auth/login' request")
	}

	rr := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(handler.Login)
	loginHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 422)
	assert.Equal(t, responseMap["error"], "json: cannot unmarshal string into Go value of type model.User")
}

func TestLogin_422_Validation(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      false,
	}

	cases := []struct {
		inputJSON    string
		errorMessage string
	}{
		{
			inputJSON:    `{"password": "random"}`,
			errorMessage: "Required Email",
		},
		{
			inputJSON:    `{"email": "random@email.com"}`,
			errorMessage: "Required Password",
		},
		{
			inputJSON:    `{"email": "randomemail.com", "password": "random"}`,
			errorMessage: "Invalid Email",
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBufferString(c.inputJSON))
		if err != nil {
			t.Error("Failed to create 'POST: /api/v1/auth/login' request")
		}

		rr := httptest.NewRecorder()
		loginHandler := http.HandlerFunc(handler.Login)
		loginHandler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, 422)
		assert.Equal(t, responseMap["error"], c.errorMessage)
	}
}

func TestLogin_500(t *testing.T) {

	userLogin := model.User{
		Email:    "random@gmail.com",
		Password: "random",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      true,
		ErrorMessage: "User not found",
	}

	jsonByte, err := json.Marshal(userLogin)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/auth/login' request")
	}

	rr := httptest.NewRecorder()
	loginHandler := http.HandlerFunc(handler.Login)
	loginHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "User not found")
}
