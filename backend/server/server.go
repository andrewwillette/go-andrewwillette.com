package server

import (
	"encoding/json"
	"fmt"
	"github.com/andrewwillette/willette_api/logging"
	"github.com/andrewwillette/willette_api/persistence"
	"net/http"
)

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const deleteSoundcloudEndpoint = "/delete-soundcloud-url"
const loginEndpoint = "/login"
const port = 9099

type UserService interface {
	Login(username, password string) (success bool, bearerToken string)
	BearerTokenExists(bearerToken string) bool
}

type SoundcloudUrlService interface {
	GetAllSoundcloudUrls() ([]persistence.SoundcloudUrl, error)
	AddSoundcloudUrl(string) error
	DeleteSoundcloudUrl(string) error
}

type WilletteAPIServer struct {
	userService          UserService
	soundcloudUrlService SoundcloudUrlService
}

func NewWilletteAPIServer(userService UserService, soundcloudUrlService SoundcloudUrlService) *WilletteAPIServer {
	return &WilletteAPIServer{userService: userService, soundcloudUrlService: soundcloudUrlService}
}

func (u *WilletteAPIServer) getAllSoundcloudUrls(w http.ResponseWriter, _ *http.Request) {
	logging.GlobalLogger.Info().Msg("getAllSoundcloudUrls called.")
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	urls, err := u.soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Failed to get soundcloud urls from service.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var soundcloudUrls []SoundcloudUrlUiOrderJson
	for _, url := range urls {
		soundcloudUrls = append(soundcloudUrls, SoundcloudUrlUiOrderJson{Url: url.Url, UiOrder: url.UiOrder})
	}
	if err = json.NewEncoder(w).Encode(soundcloudUrls); err != nil {
		logging.GlobalLogger.Err(err).Msg("Failed to encode soundcloud urls in http response.")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	return
}

func (u *WilletteAPIServer) addSoundcloudUrl(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("addSoundcloudUrl called.")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		return
	}
	decoder := json.NewDecoder(r.Body)
	var soundcloudData SoundcloudUrlJson
	err := decoder.Decode(&soundcloudData)
	logging.GlobalLogger.Info().Msg(fmt.Sprintf("%+v", soundcloudData))
	if err != nil {
		logging.GlobalLogger.Info().Msg("Failed to decode soundcloud data.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if u.userService.BearerTokenExists(r.Header.Get("Authorization")) {
		logging.GlobalLogger.Info().Msg("Bearertoken is valid.")
		err := u.soundcloudUrlService.AddSoundcloudUrl(soundcloudData.Url)
		if err != nil {
			logging.GlobalLogger.Err(err).Msg("Error when adding soundcloud url to service layer.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Success adding soundcloud url. url: %s", soundcloudData.Url))
		return
	} else {
		logging.GlobalLogger.Info().Msg("Bearertoken is invalid.")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) deleteSoundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("deleteSoundcloudUrlPost called.")
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, x-requested-with, Authorization")
	decoder := json.NewDecoder(r.Body)
	var soundcloudData SoundcloudUrlJson
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		logging.GlobalLogger.Info().Msg("Failed to decode soundcloud data in delete.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if u.userService.BearerTokenExists(r.Header.Get("Authorization")) {
		err = u.soundcloudUrlService.DeleteSoundcloudUrl(soundcloudData.Url)
		if err != nil {
			logging.GlobalLogger.Err(err).Msg("Error deleting soundcloudUrl in service layer.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("deleteSoundcloudUrl called successfully for item: %s", soundcloudData.Url))
		return
	} else {
		logging.GlobalLogger.Info().
			Msg(fmt.Sprintf("deleteSoundcloudUrl called unauthorized for item: %s, bearerToken: %s", soundcloudData.Url, r.Header.Get("Authorization")))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) login(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("Login called.")
	if r.Method != "POST" {
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Login not POST method. Method: %s", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var userCredentials UserJson
	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		logging.GlobalLogger.Info().Msg("Error decoding user credentials from request body.")
		http.Error(w, fmt.Sprintf("Error decoding user credentials from request body. %v", err), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	user := UserJson{Username: userCredentials.Username, Password: userCredentials.Password}
	loginSuccessful, bearerToken := u.userService.Login(user.Username, user.Password)
	if loginSuccessful {
		if err := json.NewEncoder(w).Encode(bearerToken); err != nil {
			logging.GlobalLogger.Err(err).Msg("Failed to encode bearer token after successful authentication.")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		logging.GlobalLogger.Info().Msg("Login Successful.")
		return
	} else {
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Login failed with username: %s, password: %s", user.Username, user.Password))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (u *WilletteAPIServer) RunServer() {
	getAllSoundcloudUrlsHandler := http.HandlerFunc(u.getAllSoundcloudUrls)
	http.Handle(getSoundcloudAllEndpoint, getAllSoundcloudUrlsHandler)

	addSoundcloudUrlHandler := http.HandlerFunc(u.addSoundcloudUrl)
	http.Handle(addSoundcloudEndpoint, addSoundcloudUrlHandler)

	deleteSoundcloudUrlHandler := http.HandlerFunc(u.deleteSoundcloudUrlPost)
	http.Handle(deleteSoundcloudEndpoint, deleteSoundcloudUrlHandler)

	loginHandler := http.HandlerFunc(u.login)
	http.Handle(loginEndpoint, loginHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		logging.GlobalLogger.Err(err).Msg(fmt.Sprintf("Failed to listen and serve to port %d", port))
		return
	}
}
