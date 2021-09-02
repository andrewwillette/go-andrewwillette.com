package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//func TestLoginPost_Invalid(t *testing.T) {
//	response := httptest.NewRecorder()
//	body := User {Username: "hello", Password: "passwordWorld"}
//	request := httptest.NewRequest(http.MethodPost, loginEndpoint, userToJSON(body))
//	loginPost(response, request)
//	assert.Equal(t, 401, response.Code)
//}

type MockUserService struct {
	Login       func(user User) (bool, string, error)
	UsersRegistered []User
}

func (m *MockUserService) Login(user User) (success bool, bearerToken string, err error) {
	return m.LoginFunc(user)
}

func (m *MockUserService) AddUser(user User) (err error) {
	m.UsersRegistered = append(m.UsersRegistered, user)
	return nil
}

func (m *MockUserService) UpdateUserBearerToken(user User, bearerToken string) (err error) {
	return nil
}

type MockSoundcloudUrlService struct {

}

func TestLoginPost_Valid(t *testing.T) {
	t.Run("can login a valid user", func(t *testing.T) {
		expectedInsertedBearerToken := "bearerTokenOne"

		userService := &MockUserService {
			LoginFunc: func(user User) (bool, string, error) {
				return true, expectedInsertedBearerToken, nil
			},
		}
		server := NewWilletteAPIServer(userService)

		user := User{Username: "usernameOne", Password: "passwordOne"}
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, loginEndpoint, userToJSON(user))

		server.Login(res, req)

		//assert.Equal(t, res.Code, http.StatusCreated)
	})
}

func userToJSON(user User) io.Reader {
	marshalledUser, _ := json.Marshal(user)
	return bytes.NewReader(marshalledUser)
}