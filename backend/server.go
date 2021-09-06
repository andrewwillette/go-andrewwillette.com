package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const deleteSoundcloudEndpoint = "/delete-soundcloud-url"
const loginEndpoint = "/login"
const port = 9099

type UserService interface {
	Login(username, password string) (success bool, bearerToken string, err error)
	BearerTokenExists(bearerToken string) bool
}

type SoundcloudUrlService interface {
	GetAllSoundcloudUrls() ([]string, error)
	AddSoundcloudUrl(string) error
	DeleteSoundcloudUrl(string) error
}

type WilletteAPIServer struct {
	userService UserService
	soundcloudUrlService SoundcloudUrlService
}

func NewWilletteAPIServer(userService UserService, soundcloudUrlService SoundcloudUrlService) *WilletteAPIServer {
	return &WilletteAPIServer{userService: userService, soundcloudUrlService: soundcloudUrlService}
}

func (u *WilletteAPIServer) getAllSoundcloudUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	urls, err := u.soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var soundcloudUrls []SoundcloudUrl
	for i := 0; i < len(urls); i++ {
		soundcloudUrls = append(soundcloudUrls, SoundcloudUrl{Url: urls[i]})
	}
	err = json.NewEncoder(w).Encode(soundcloudUrls)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func (u *WilletteAPIServer) addSoundcloudUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var soundcloudData AuthenticatedSoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if u.userService.BearerTokenExists(soundcloudData.BearerToken) {
		err := u.soundcloudUrlService.AddSoundcloudUrl(soundcloudData.Url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) deleteSoundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var soundcloudData AuthenticatedSoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if u.userService.BearerTokenExists(soundcloudData.BearerToken) {
		err = u.soundcloudUrlService.DeleteSoundcloudUrl(soundcloudData.Url)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var userCredentials User
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		http.Error(w, fmt.Sprintf("cloud not decode user payload: %v", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user := User{Username: userCredentials.Username, Password: userCredentials.Password}
	loginSuccessful, bearerToken, err := u.userService.Login(user.Username, user.Password)
	if loginSuccessful {
		if err = json.NewEncoder(w).Encode(bearerToken); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) runServer() {
	getAllSoundcloudUrlsHandler := http.HandlerFunc(u.getAllSoundcloudUrls)
	http.Handle(getSoundcloudAllEndpoint, getAllSoundcloudUrlsHandler)

	putHandler := http.HandlerFunc(u.addSoundcloudUrl)
	http.Handle(addSoundcloudEndpoint, putHandler)

	deleteHandler := http.HandlerFunc(u.deleteSoundcloudUrlPost)
	http.Handle(deleteSoundcloudEndpoint, deleteHandler)

	loginHandler := http.HandlerFunc(u.login)
	http.Handle(loginEndpoint, loginHandler)

	fmt.Println("running server")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		println(err.Error())
		log.Fatal(fmt.Sprintf("Failed to listen and serve to port %d", port))
		return
	}
}
