package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/andrewwillette/willette_api/persistence"
	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	UsersRegistered  []UserJson
	LoginFunc        func(username string, password string) (success bool, bearerToken string)
	IsAuthorizedFunc func(bearerToken string) bool
}

func (m *MockUserService) Login(username, password string) (success bool, bearerToken string) {
	return m.LoginFunc(username, password)
}

func (m *MockUserService) IsAuthorized(bearerToken string) bool {
	return m.IsAuthorizedFunc(bearerToken)
}

type MockSoundcloudUrlService struct {
	GetAllSoundcloudUrlsFunc     func() ([]persistence.SoundcloudUrl, error)
	AddSoundcloudUrlsFunc        func(s string) error
	DeleteSoundcloudUrlFunc      func(s string) error
	UpdateSoundcloudUiOrdersFunc func([]persistence.SoundcloudUrl) error
	SoundcloudUrls               []persistence.SoundcloudUrl
}

func (m *MockSoundcloudUrlService) GetAllSoundcloudUrls() ([]persistence.SoundcloudUrl, error) {
	return m.GetAllSoundcloudUrlsFunc()
}

func (m *MockSoundcloudUrlService) AddSoundcloudUrl(s string) error {
	return m.AddSoundcloudUrlsFunc(s)
}

func (m MockSoundcloudUrlService) DeleteSoundcloudUrl(s string) error {
	return m.DeleteSoundcloudUrlFunc(s)
}

func (m MockSoundcloudUrlService) UpdateSoundcloudUiOrders(urls []persistence.SoundcloudUrl) error {
	return m.UpdateSoundcloudUiOrdersFunc(urls)
}

func TestLogin(t *testing.T) {
	t.Run("invalid user login", func(t *testing.T) {
		response := httptest.NewRecorder()
		// body := UserJson{Username: "hello", Password: "passwordWorld"}
		// var users []UserJson
		// userService := &MockUserService{
		// 	UsersRegistered: users,
		// 	LoginFunc: func(username, password string) (success bool, bearerToken string) {
		// 		return false, ""
		// 	},
		// 	IsAuthorizedFunc: func(bearerToken string) bool {
		// 		return true
		// 	},
		// }
		// var soundcloudUrls []persistence.SoundcloudUrl
		// soundcloudUrlService := &MockSoundcloudUrlService{
		// 	SoundcloudUrls: soundcloudUrls,
		// }
		// server := newWebServices(userService, soundcloudUrlService)
		// request := httptest.NewRequest(http.MethodPost, loginEndpoint, userToJSON(body))
		// server.login(response, request)
		// assert.Equal(t, 401, response.Code)
	})

	t.Run("valid user login", func(t *testing.T) {
		// var users []UserJson
		// testBearerToken := "testBearerToken"
		// userService := &MockUserService{
		// 	UsersRegistered: users,
		// 	LoginFunc: func(username, password string) (success bool, bearerToken string) {
		// 		return true, testBearerToken
		// 	},
		// 	IsAuthorizedFunc: func(bearerToken string) bool {
		// 		return true
		// 	},
		// }
		// var soundcloudUrls []persistence.SoundcloudUrl
		// soundcloudUrlService := &MockSoundcloudUrlService{
		// 	SoundcloudUrls: soundcloudUrls,
		// }
		// server := newWebServices(userService, soundcloudUrlService)
		// response := httptest.NewRecorder()
		// body := UserJson{Username: "hello", Password: "passwordWorld"}
		// request := httptest.NewRequest(http.MethodPost, loginEndpoint, userToJSON(body))
		// server.login(response, request)
		// assert.Equal(t, 200, response.Code)
		// assert.Contains(t, response.Body.String(), testBearerToken)
	})

	t.Run("GET returns 405", func(t *testing.T) {
		// var users []UserJson
		// testBearerToken := "testBearerToken"
		// userService := &MockUserService{
		// 	UsersRegistered: users,
		// 	LoginFunc: func(username, password string) (success bool, bearerToken string) {
		// 		return true, testBearerToken
		// 	},
		// 	IsAuthorizedFunc: func(bearerToken string) bool {
		// 		return true
		// 	},
		// }
		// var soundcloudUrls []persistence.SoundcloudUrl
		// soundcloudUrlService := &MockSoundcloudUrlService{
		// 	SoundcloudUrls: soundcloudUrls,
		// }
		// server := newWebServices(userService, soundcloudUrlService)
		// response := httptest.NewRecorder()
		// body := UserJson{Username: "hello", Password: "passwordWorld"}
		// request := httptest.NewRequest(http.MethodGet, loginEndpoint, userToJSON(body))
		// server.login(response, request)
		// assert.Equal(t, 405, response.Code)
	})
}
func TestAddSoundcloudUrl(t *testing.T) {
	t.Run("valid bearer token", func(t *testing.T) {
		// var users []UserJson
		// testBearerToken := "testBearerToken"
		// userService := &MockUserService{
		// 	UsersRegistered: users,
		// 	LoginFunc: func(username, password string) (success bool, bearerToken string) {
		// 		return true, testBearerToken
		// 	},
		// 	IsAuthorizedFunc: func(bearerToken string) bool {
		// 		return true
		// 	},
		// }
		// soundcloudUrls := []persistence.SoundcloudUrl{{Url: "urlone.com", UiOrder: 1, Id: 1},
		// 	{Url: "urltwo.com", Id: 2, UiOrder: 3}}
		// soundcloudUrlService := &MockSoundcloudUrlService{
		// 	SoundcloudUrls: soundcloudUrls,
		// 	GetAllSoundcloudUrlsFunc: func() ([]persistence.SoundcloudUrl, error) {
		// 		return soundcloudUrls, nil
		// 	},
		// 	AddSoundcloudUrlsFunc: func(s string) error {
		// 		soundcloudUrl := persistence.SoundcloudUrl{Url: s}
		// 		soundcloudUrls = append(soundcloudUrls, soundcloudUrl)
		// 		return nil
		// 	},
		// 	DeleteSoundcloudUrlFunc: func(s string) error {
		// 		return nil
		// 	},
		// }
		// server := newWebServices(userService, soundcloudUrlService)
		// response := httptest.NewRecorder()
		// newSoundcloudUrl := "testsoundcloudurl.com"
		// body := SoundcloudUrlJson{Url: newSoundcloudUrl}
		// request := httptest.NewRequest(http.MethodPost, loginEndpoint, authenticatedSoundcloudUrlToJSON(body))
		// server.addSoundcloudUrl(response, request)
		// responseTwo := httptest.NewRecorder()
		// requestTwo := httptest.NewRequest(http.MethodGet, getSoundcloudAllEndpoint, nil)
		// server.getAllSoundcloudUrls(responseTwo, requestTwo)
		// fmt.Printf("new soundcloud url is %s\n", responseTwo.Body.String())
		// decoder := json.NewDecoder(responseTwo.Body)
		// var soundcloudData []persistence.SoundcloudUrl
		// err := decoder.Decode(&soundcloudData)
		// if err != nil {
		// 	t.Log("failed to decode soundcloud data")
		// 	t.Fail()
		// }
		// assert.ElementsMatch(t, soundcloudData, []persistence.SoundcloudUrl{
		// 	{Id: 0, Url: "urlone.com", UiOrder: 1},
		// 	{Id: 0, Url: "urltwo.com", UiOrder: 3},
		// 	{Id: 0, Url: "testsoundcloudurl.com", UiOrder: 0}})
	})
}

func authenticatedSoundcloudUrlToJSON(url SoundcloudUrlJson) io.Reader {
	marshalledUser, _ := json.Marshal(url)
	return bytes.NewReader(marshalledUser)
}

func userToJSON(user UserJson) io.Reader {
	marshalledUser, _ := json.Marshal(user)
	return bytes.NewReader(marshalledUser)
}
