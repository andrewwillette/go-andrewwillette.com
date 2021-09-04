package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserService struct {
	UsersRegistered       []User
	LoginFunc             func(username string, password string) (success bool, bearerToken string, err error)
	BearerTokenExistsFunc func(bearerToken string) bool
}

func (m *MockUserService) Login(username, password string) (success bool, bearerToken string, err error) {
	return m.LoginFunc(username, password)
}

func (m *MockUserService) BearerTokenExists(bearerToken string) bool {
	return m.BearerTokenExistsFunc(bearerToken)
}

type MockSoundcloudUrlService struct {
	GetAllSoundcloudUrlsFunc func() ([]string, error)
	AddSoundcloudUrlsFunc    func(s string) error
	DeleteSoundcloudUrls     func(s string) error
	SoundcloudUrls           []string
}

func (m *MockSoundcloudUrlService) GetAllSoundcloudUrls() ([]string, error) {
	return m.GetAllSoundcloudUrlsFunc()
}

func (m *MockSoundcloudUrlService) AddSoundcloudUrl(s string) error {
	return m.AddSoundcloudUrlsFunc(s)
}

func (m MockSoundcloudUrlService) DeleteSoundcloudUrl(s string) error {
	return m.DeleteSoundcloudUrls(s)
}

func TestLogin(t *testing.T) {
	t.Run("invalid user login", func(t *testing.T) {
		response := httptest.NewRecorder()
		body := User{Username: "hello", Password: "passwordWorld"}
		var users []User
		userService := &MockUserService{
			UsersRegistered: users,
			LoginFunc: func(username, password string) (success bool, bearerToken string, err error) {
				return false, "", nil
			},
			BearerTokenExistsFunc: func(bearerToken string) bool {
				return true
			},
		}
		var soundcloudUrls []string
		soundcloudUrlService := &MockSoundcloudUrlService{
			SoundcloudUrls: soundcloudUrls,
		}
		server := NewWilletteAPIServer(userService, soundcloudUrlService)
		request := httptest.NewRequest(http.MethodPost, loginEndpoint, userToJSON(body))
		server.login(response, request)
		assert.Equal(t, 401, response.Code)
	})

	t.Run("valid user login", func(t *testing.T) {
		response := httptest.NewRecorder()
		body := User{Username: "hello", Password: "passwordWorld"}
		var users []User
		bearertoken := "testBearerToken"
		userService := &MockUserService{
			UsersRegistered: users,
			LoginFunc: func(username, password string) (success bool, bearerToken string, err error) {
				return true, bearertoken, nil
			},
			BearerTokenExistsFunc: func(bearerToken string) bool {
				return true
			},
		}
		var soundcloudUrls []string
		soundcloudUrlService := &MockSoundcloudUrlService{
			SoundcloudUrls: soundcloudUrls,
		}
		server := NewWilletteAPIServer(userService, soundcloudUrlService)
		request := httptest.NewRequest(http.MethodPost, loginEndpoint, userToJSON(body))
		server.login(response, request)
		assert.Equal(t, 200, response.Code)
		assert.Contains(t, response.Body.String(), bearertoken)
	})
}
func TestAddSoundcloudUrl(t *testing.T) {

	t.Run("valid bearer token", func(t *testing.T) {

	})
}

func userToJSON(user User) io.Reader {
	marshalledUser, _ := json.Marshal(user)
	return bytes.NewReader(marshalledUser)
}
