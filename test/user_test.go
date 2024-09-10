package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"golang-restful-api/internal/entity"
	"golang-restful-api/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	ClearAll()
	requestBody := model.RegisterUserRequest{
		Name:     "Habibi Iberahim",
		Email:    "habibiiberahim21@gmail.com",
		Phone:    "085248458503",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Email, responseBody.Data.Email)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.Phone, responseBody.Data.Phone)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestRegisterError(t *testing.T) {
	ClearAll()
	requestBody := model.RegisterUserRequest{
		Name:     "",
		Email:    "",
		Phone:    "",
		Password: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestRegisterDuplicate(t *testing.T) {
	ClearAll()
	TestRegister(t)
	requestBody := model.RegisterUserRequest{
		Name:     "Habibi Iberahim",
		Email:    "habibiiberahim21@gmail.com",
		Phone:    "085248458503",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestLogin(t *testing.T) {
	TestRegister(t)

	requestBody := model.LoginUserRequest{
		Email:    "habibiiberahim21@gmail.com",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.NotNil(t, responseBody.Data.AccessToken)
	assert.NotNil(t, responseBody.Data.RefreshToken)

	user := new(entity.User)
	err = db.Where("email = ?", requestBody.Email).First(user).Error
	assert.Nil(t, err)
	assert.Equal(t, user.Token, responseBody.Data.RefreshToken)
}

func TestLoginWrongUsername(t *testing.T) {
	ClearAll()
	TestRegister(t)

	requestBody := model.LoginUserRequest{
		Email:    "wrong",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestLoginWrongPassword(t *testing.T) {
	ClearAll()
	TestRegister(t)

	requestBody := model.LoginUserRequest{
		Email:    "habibiiberahim21@gmail.com",
		Password: "wrong",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestLogout(t *testing.T) {
	TestRegister(t)

	//login request
	requestBody := model.LoginUserRequest{
		Email:    "habibiiberahim21@gmail.com",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	//logout requset
	request = httptest.NewRequest(http.MethodDelete, "/api/users/logout", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", responseBody.Data.AccessToken)

	response, err = app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	newResponseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, newResponseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestUpdateName(t *testing.T) {
	TestRegister(t)

	//login request
	loginRequestBody := model.LoginUserRequest{
		Email:    "habibiiberahim21@gmail.com",
		Password: "secret-password",
	}

	bodyJson, err := json.Marshal(loginRequestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/users/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	responseBytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(responseBytes, responseBody)
	assert.Nil(t, err)

	//update request
	updateRequestBody := model.UpdateUserRequest{
		Name: "Habibi Iberahim !!!",
	}

	bodyJson, err = json.Marshal(updateRequestBody)
	assert.Nil(t, err)

	request = httptest.NewRequest(http.MethodPatch, "/api/users/update", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", responseBody.Data.AccessToken)

	response, err = app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	newResponseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(bytes, newResponseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, updateRequestBody.Name, newResponseBody.Data.Name)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)

}
