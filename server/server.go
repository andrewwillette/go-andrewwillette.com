package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/andrewwillette/willette_api/config"
	"github.com/andrewwillette/willette_api/logging"
	"github.com/andrewwillette/willette_api/persistence"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const (
	getSoundcloudAllEndpoint     = "/get-soundcloud-urls"
	addSoundcloudEndpoint        = "/add-soundcloud-url"
	deleteSoundcloudEndpoint     = "/delete-soundcloud-url"
	loginEndpoint                = "/login"
	updateSoundcloudUrlsEndpoint = "/update-soundcloud-urls"
)

// userService manages logging users in and authenticating tokens
type userService interface {
	Login(username, password string) (success bool, authToken string)
	IsAuthorized(authToken string) bool
}

type soundcloudUrlService interface {
	GetAllSoundcloudUrls() ([]persistence.SoundcloudUrl, error)
	AddSoundcloudUrl(string) error
	DeleteSoundcloudUrl(string) error
	UpdateSoundcloudUiOrders([]persistence.SoundcloudUrl) error
}

type webServices struct {
	userService          userService
	soundcloudUrlService soundcloudUrlService
}

type router struct {
	server webServices
}

func newWebServices(userService userService, soundcloudUrlService soundcloudUrlService) *webServices {
	return &webServices{
		userService:          userService,
		soundcloudUrlService: soundcloudUrlService,
	}
}

func getNewRelicApp() *newrelic.Application {
	logging.GlobalLogger.Info().Msg("Starting up new relic")
	newrelicLicense := os.Getenv("NEW_RELIC_LICENSE")
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("go-andrewwillette"),
		newrelic.ConfigLicense(newrelicLicense),
		// newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		logging.GlobalLogger.Error().Msgf("Failed to start new relic app, newrelic key: %s", newrelicLicense)
	}
	return app
}

func RunServer() {
	databaseFile := config.GetDatabaseFile()
	persistence.InitDatabaseIdempotent(databaseFile)
	userService := &persistence.UserService{SqliteDbFile: databaseFile}
	soundcloudUrlService := &persistence.SoundcloudUrlService{SqliteFile: databaseFile}

	websiteServices := newWebServices(userService, soundcloudUrlService)
	e := echo.New()
	e.Use(nrecho.Middleware(getNewRelicApp()))
	e.GET(getSoundcloudAllEndpoint, websiteServices.getAllSoundcloudUrlsEcho)
	e.POST(loginEndpoint, websiteServices.getAllSoundcloudUrlsEcho)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

	// router := router{server: *websiteServer}
	// if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), &router); err != nil {
	// 	logging.GlobalLogger.Err(err).Msg(fmt.Sprintf("Failed to run willetteAPIServer on port: %d", config.Port))
	// 	logging.GlobalLogger.Fatal().Msg("Server failed to start. Exiting application.")
	// 	return
	// }
}

func (u *webServices) getAllSoundcloudUrlsEcho(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application-json")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	defReqHeaders(c)
	urls, err := u.soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		const errMsg = "Failed to get soundcloud urls from service."
		logging.GlobalLogger.Err(err).Msg(errMsg)
		return c.String(http.StatusInternalServerError, errMsg)
	}
	var soundcloudUrls = []SoundcloudUrlUiOrderJson{}
	for _, url := range urls {
		soundcloudUrls = append(soundcloudUrls, SoundcloudUrlUiOrderJson{Url: url.Url, UiOrder: url.UiOrder})
	}

	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(soundcloudUrls)
}

// func (u *webServices) getAllSoundcloudUrls(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application-json")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	urls, err := u.soundcloudUrlService.GetAllSoundcloudUrls()
// 	addDefaultRequestHeaders(w, r)
// 	if err != nil {
// 		logging.GlobalLogger.Err(err).Msg("Failed to get soundcloud urls from service.")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	var soundcloudUrls []SoundcloudUrlUiOrderJson
// 	for _, url := range urls {
// 		soundcloudUrls = append(soundcloudUrls, SoundcloudUrlUiOrderJson{Url: url.Url, UiOrder: url.UiOrder})
// 	}
// 	if err = json.NewEncoder(w).Encode(soundcloudUrls); err != nil {
// 		logging.GlobalLogger.Err(err).Msg("Failed to encode soundcloud urls in http response.")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	return
// }

func (u *webServices) addSoundcloudUrl(w http.ResponseWriter, r *http.Request) {
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
	if u.userService.IsAuthorized(r.Header.Get("Authorization")) {
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

func (u *webServices) deleteSoundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	addDefaultRequestHeaders(w, r)
	if r.Method == "OPTIONS" {
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
	if u.userService.IsAuthorized(r.Header.Get("Authorization")) {
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
func (u *webServices) updateSoundcloudUrlUiOrders(w http.ResponseWriter, r *http.Request) {
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

func (u *webServices) loginEcho(c echo.Context) error {
	defReqHeaders(c)
	if c.Request().Method == "OPTIONS" {
		return c.String(http.StatusOK, "Allowing OPTIONS because of prior failed handshaking.")
	}
	var userCredentials UserJson
	if err := json.NewDecoder(c.Request().Body).Decode(&userCredentials); err != nil {
		const errMsg = "Error decoding user credentials from request body."
		logging.GlobalLogger.Info().Msg(errMsg)

		return c.String(http.StatusInternalServerError, errMsg)
	}
	c.Response().Header().Set("Content-Type", "application-json")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	user := UserJson{Username: userCredentials.Username, Password: userCredentials.Password}
	loginSuccessful, authToken := u.userService.Login(user.Username, user.Password)
	if loginSuccessful {
		if err := json.NewEncoder(c.Response()).Encode(authToken); err != nil {
			const errMsg = "Failed to encode authToken after successful authentication."
			logging.GlobalLogger.Err(err).Msg(errMsg)
			return c.String(http.StatusUnauthorized, errMsg)
		}
		logging.GlobalLogger.Info().Msg("Login Successful.")
		return c.String(http.StatusOK, "")
	} else {
		logging.GlobalLogger.Info().Msg(fmt.Sprintf("Login failed with username: %s, password: %s", user.Username, user.Password))
		return c.String(http.StatusUnauthorized, "Login Failed.")
	}
}
func (u *webServices) login(w http.ResponseWriter, r *http.Request) {
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

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logging.GlobalLogger.Info().Msg(fmt.Sprintf("HTTP Request received. Path: %s", req.URL.Path))
	// switch req.URL.Path {
	// case getSoundcloudAllEndpoint:
	// 	r.server.getAllSoundcloudUrls(w, req)
	// case addSoundcloudEndpoint:
	// 	r.server.addSoundcloudUrl(w, req)
	// case deleteSoundcloudEndpoint:
	// 	r.server.deleteSoundcloudUrlPost(w, req)
	// case loginEndpoint:
	// 	r.server.login(w, req)
	// case updateSoundcloudUrlsEndpoint:
	// 	r.server.updateSoundcloudUrlUiOrders(w, req)
	// default:
	// 	http.Error(w, "404 not found", http.StatusNotFound)
	// }
}

/**
Headers to add to all responses.
*/
func addDefaultRequestHeaders(w http.ResponseWriter, r *http.Request) {
	for _, originUrl := range config.GetCorsWhiteList() {
		if r.Header.Get("origin") == originUrl {
			w.Header().Set("Access-Control-Allow-Origin", originUrl)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func defReqHeaders(c echo.Context) {
	c.Response().Header().Set("Content-Type", "application-json")
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	for _, originUrl := range config.GetCorsWhiteList() {
		if c.Request().Header.Get("origin") == originUrl {
			c.Response().Header().Set("Access-Control-Allow-Origin", originUrl)
		}
	}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
