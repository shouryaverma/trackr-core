// +build !integration

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
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateUser_201(t *testing.T) {

	userToCreate := model.User{
		Email:     "random@gmail.com",
		Password:  "random",
		FirstName: "John",
		LastName:  "Doe",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &userToCreate,
		IsError:      false,
	}

	jsonByte, err := json.Marshal(userToCreate)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/users' request")
	}

	rr := httptest.NewRecorder()
	createUserHandler := http.HandlerFunc(handler.CreateUser)
	createUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	createdUser := responseMap["user"].(map[string]interface{})

	assert.Equal(t, rr.Code, 201)
	assert.Equal(t, createdUser["first_name"], userToCreate.FirstName)
}

func TestCreateUser_422_JSON(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      false,
	}

	jsonByte, err := json.Marshal("bruh}")
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/users' request")
	}

	rr := httptest.NewRecorder()
	createUserHandler := http.HandlerFunc(handler.CreateUser)
	createUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 422)
	assert.Equal(t, responseMap["error"], "json: cannot unmarshal string into Go value of type model.User")
}

func TestCreateUser_422_Validation(t *testing.T) {

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
			inputJSON:    `{"email": "random@gmail.com"}`,
			errorMessage: "Required Password",
		},
		{
			inputJSON:    `{"email": "randomgmail.com", "password": "random"}`,
			errorMessage: "Invalid Email",
		},
		{
			inputJSON:    `{"email": "random@gmail.com", "password": "random"}`,
			errorMessage: "Required First Name",
		},
		{
			inputJSON:    `{"email": "random@gmail.com", "password": "random", "first_name": "John"}`,
			errorMessage: "Required Last Name",
		},
	}

	for _, c := range cases {
		req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBufferString(c.inputJSON))
		if err != nil {
			t.Error("Failed to create 'POST: /api/v1/users' request")
		}

		rr := httptest.NewRecorder()
		createUserHandler := http.HandlerFunc(handler.CreateUser)
		createUserHandler.ServeHTTP(rr, req)

		responseMap := make(map[string]interface{})
		err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
		if err != nil {
			fmt.Printf("Cannot convert to json: %v", err)
		}

		assert.Equal(t, rr.Code, 422)
		assert.Equal(t, responseMap["error"], c.errorMessage)
	}
}

func TestCreateUser_500(t *testing.T) {

	userToCreate := model.User{
		Email:     "random@gmail.com",
		Password:  "random",
		FirstName: "John",
		LastName:  "Doe",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      true,
		ErrorMessage: "Table 'users' doesn't exist",
	}

	jsonByte, err := json.Marshal(userToCreate)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'POST: /api/v1/users' request")
	}

	rr := httptest.NewRecorder()
	createUserHandler := http.HandlerFunc(handler.CreateUser)
	createUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "Table 'users' doesn't exist")
}

func TestGetAllUsers_200(t *testing.T) {

	usersToGet := []model.User{
		{
			Email:     "john@gmail.com",
			Password:  "random",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			Email:     "mario@gmail.com",
			Password:  "random",
			FirstName: "Mario",
			LastName:  "Draghi",
		},
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &usersToGet,
		IsError:      false,
	}

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'GET: /api/v1/users' request")
	}

	rr := httptest.NewRecorder()
	getAllUsersHandler := http.HandlerFunc(handler.GetAllUsers)
	getAllUsersHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	users := responseMap["users"].([]interface{})
	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, len(users), len(usersToGet))
}

func TestGetAllUsers_500(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &[]model.User{},
		IsError:      true,
		ErrorMessage: "Table 'users' doesn't exist",
	}

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'GET: /api/v1/users' request")
	}

	rr := httptest.NewRecorder()
	getAllUsersHandler := http.HandlerFunc(handler.GetAllUsers)
	getAllUsersHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "Table 'users' doesn't exist")
}

func TestGetUser_200(t *testing.T) {

	userToGet := model.User{
		Email:     "john@gmail.com",
		Password:  "random",
		FirstName: "John",
		LastName:  "Doe",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &userToGet,
		IsError:      false,
	}

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'GET: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	getUserHandler := http.HandlerFunc(handler.GetUser)
	getUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	user := responseMap["user"].(map[string]interface{})
	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, user["first_name"], userToGet.FirstName)
}

func TestGetUser_500(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      true,
		ErrorMessage: "User not found",
	}

	req, err := http.NewRequest("GET", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'GET: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	getUserHandler := http.HandlerFunc(handler.GetUser)
	getUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "User not found")
}

func TestUpdateUser_200(t *testing.T) {

	userUpdate := model.User{
		FirstName: "Mario",
		LastName:  "Draghi",
	}

	updatedUser := model.User{
		Email:     "john@gmail.com",
		Password:  "random",
		FirstName: "Mario",
		LastName:  "Draghi",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &updatedUser,
		IsError:      false,
	}

	jsonByte, err := json.Marshal(&userUpdate)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("PUT", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'PUT: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	updateUserHandler := http.HandlerFunc(handler.UpdateUser)
	updateUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	user := responseMap["user"].(map[string]interface{})
	assert.Equal(t, rr.Code, 200)
	assert.Equal(t, user["first_name"], updatedUser.FirstName)
}

func TestUpdateUser_422(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      false,
	}

	jsonByte, err := json.Marshal("{bruh")
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("PUT", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'PUT: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	updateUserHandler := http.HandlerFunc(handler.UpdateUser)
	updateUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 422)
	assert.Equal(t, responseMap["error"], "json: cannot unmarshal string into Go value of type model.User")
}

func TestUpdateUser_500(t *testing.T) {

	userUpdate := model.User{
		FirstName: "Mario",
		LastName:  "Draghi",
	}

	handler.pgRepo = &mock.Repository{
		ReturnObject: &model.User{},
		IsError:      true,
		ErrorMessage: "User not found",
	}

	jsonByte, err := json.Marshal(&userUpdate)
	if err != nil {
		t.Error("Failed to marshal User struct")
	}

	req, err := http.NewRequest("PUT", "/api/v1/users", bytes.NewBuffer(jsonByte))
	if err != nil {
		t.Error("Failed to create 'PUT: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	updateUserHandler := http.HandlerFunc(handler.UpdateUser)
	updateUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "User not found")
}

func TestDeleteUser_204(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: int64(1),
		IsError:      false,
	}

	req, err := http.NewRequest("DELETE", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'DELETE: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	deleteUserHandler := http.HandlerFunc(handler.DeleteUser)
	deleteUserHandler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 204)
}

func TestDeleteUser_500(t *testing.T) {

	handler.pgRepo = &mock.Repository{
		ReturnObject: int64(0),
		IsError:      true,
		ErrorMessage: "User not found",
	}

	req, err := http.NewRequest("DELETE", "/api/v1/users", nil)
	if err != nil {
		t.Error("Failed to create 'DELETE: /api/v1/users/{id}' request")
	}

	userID := uuid.NewV4()
	req = mux.SetURLVars(req, map[string]string{"id": userID.String()})
	rr := httptest.NewRecorder()
	deleteUserHandler := http.HandlerFunc(handler.DeleteUser)
	deleteUserHandler.ServeHTTP(rr, req)

	responseMap := make(map[string]interface{})
	err = json.Unmarshal(rr.Body.Bytes(), &responseMap)
	if err != nil {
		fmt.Printf("Cannot convert to json: %v", err)
	}

	assert.Equal(t, rr.Code, 500)
	assert.Equal(t, responseMap["error"], "User not found")
}
