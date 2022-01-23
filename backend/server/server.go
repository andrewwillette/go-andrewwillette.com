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
const updateSoundcloudUrlsEndpoint = "/update-soundcloud-urls"
const port = 9099

type UserService interface {
	Login(username, password string) (success bool, willetteToken string)
	WilletteTokenExists(willetteToken string) bool
}

type SoundcloudUrlService interface {
	GetAllSoundcloudUrls() ([]persistence.SoundcloudUrl, error)
	AddSoundcloudUrl(string) error
	DeleteSoundcloudUrl(string) error
	UpdateSoundcloudUiOrders([]persistence.SoundcloudUrl) error
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

/**
Headers to add to all responses. Hacky, one-size-fits-all, but CORs is a pain and I don't have the style yet.
*/
func addDefaultRequestHeaders(w http.ResponseWriter, r *http.Request) {
	originWhiteList := []string{"http://localhost:3000", "http://andrewwillette.com"}
	for _, originUrl := range originWhiteList {
		if r.Header.Get("origin") == originUrl {
			w.Header().Set("Access-Control-Allow-Origin", originUrl)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func (u *WilletteAPIServer) addSoundcloudUrl(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("addSoundcloudUrl called.")
	addDefaultRequestHeaders(w, r)
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
	if u.userService.WilletteTokenExists(r.Header.Get("Authorization")) {
		logging.GlobalLogger.Info().Msg("WilletteToken is valid.")
		err := u.soundcloudUrlService.AddSoundcloudUrl(soundcloudData.Url)
		if err != nil {
			logging.GlobalLogger.Err(err).Msg("Error when adding soundcloud url to service layer.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Success adding soundcloud url. url: %s", soundcloudData.Url))
		return
	} else {
		logging.GlobalLogger.Info().Msg("WilletteToken is invalid.")
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
	if r.Method == "OPTIONS" {
		addDefaultRequestHeaders(w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var soundcloudData SoundcloudUrlJson
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		logging.GlobalLogger.Info().Msg("Failed to decode soundcloud data in delete.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if u.userService.WilletteTokenExists(r.Header.Get("Authorization")) {
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
			Msg(fmt.Sprintf("deleteSoundcloudUrl called unauthorized for item: %s, WilletteToken: %s", soundcloudData.Url, r.Header.Get("Authorization")))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

/**
Update soundcloud url uiOrder values.
*/
func (u *WilletteAPIServer) updateSoundcloudUrlUiOrders(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("updateSoundcloudUrlUiOrders called.")
	addDefaultRequestHeaders(w, r)
	if r.Method == "OPTIONS" {
		return
	}
	if r.Method != "PUT" {
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Update soundcloudUrls must be PUT method. Method provided: %s", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var urls []SoundcloudUrlUiOrderJson
	//b, _ := io.ReadAll(r.Body)
	//println(b)
	if err := json.NewDecoder(r.Body).Decode(&urls); err != nil {
		logging.GlobalLogger.Info().Msg("Error decoding soundcloud urls in update soundcloud urls.")
		http.Error(w, "Error decoding soundcloud urls.", http.StatusBadRequest)
		return
	}
	var persistenceUrls []persistence.SoundcloudUrl
	for _, v := range urls {
		persistenceUrls = append(persistenceUrls, persistence.SoundcloudUrl{Url: v.Url, UiOrder: v.UiOrder})
	}
	if err := u.soundcloudUrlService.UpdateSoundcloudUiOrders(persistenceUrls); err != nil {
		http.Error(w, "Error decoding soundcloud urls.", http.StatusBadRequest)
		logging.GlobalLogger.Err(err).Msg("Error updating soundcloud urls.")
	}
}

func (u *WilletteAPIServer) login(w http.ResponseWriter, r *http.Request) {
	logging.GlobalLogger.Info().Msg("Login called.")
	addDefaultRequestHeaders(w, r)
	if r.Method == "OPTIONS" {
		return
	}
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
	loginSuccessful, willetteToken := u.userService.Login(user.Username, user.Password)
	if loginSuccessful {
		if err := json.NewEncoder(w).Encode(willetteToken); err != nil {
			logging.GlobalLogger.Err(err).Msg("Failed to encode WilletteToken after successful authentication.")
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
	http.Handle(getSoundcloudAllEndpoint, http.HandlerFunc(u.getAllSoundcloudUrls))
	http.Handle(addSoundcloudEndpoint, http.HandlerFunc(u.addSoundcloudUrl))
	http.Handle(deleteSoundcloudEndpoint, http.HandlerFunc(u.deleteSoundcloudUrlPost))
	http.Handle(loginEndpoint, http.HandlerFunc(u.login))
	http.Handle(updateSoundcloudUrlsEndpoint, http.HandlerFunc(u.updateSoundcloudUrlUiOrders))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		logging.GlobalLogger.Err(err).Msg(fmt.Sprintf("Failed to run WilletteAPIServer on port: %d", port))
		logging.GlobalLogger.Fatal().Msg("Server failed to start. Exiting application.")
		return
	}
}
