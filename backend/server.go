package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrewwillette/willette_api/persistence"
	"log"
	"net/http"
)

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const deleteSoundcloudEndpoint = "/delete-soundcloud-url"
const loginEndpoint = "/login"
const port = 9099

func runServer() {
	getHandler := http.HandlerFunc(soundcloudUrlsGet)
	http.Handle(getSoundcloudAllEndpoint, getHandler)

	putHandler := http.HandlerFunc(addSoundcloudUrlPost)
	http.Handle(addSoundcloudEndpoint, putHandler)

	deleteHandler := http.HandlerFunc(deleteSoundcloudUrlPost)
	http.Handle(deleteSoundcloudEndpoint, deleteHandler)

	loginHandler := http.HandlerFunc(loginPost)
	http.Handle(loginEndpoint, loginHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		println(err.Error())
		log.Fatal(fmt.Sprintf("Failed to listen and serve to port %d", port))
		return
	}
}

func soundcloudUrlsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	urls := persistence.GetAllSoundcloudUrls(persistence.SqlLiteDatabaseFileName)
	var soundcloudUrls []SoundcloudUrl
	for i := 0; i < len(urls); i++ {
		soundcloudUrls = append(soundcloudUrls, SoundcloudUrl{Url: urls[i]})
	}
	err := json.NewEncoder(w).Encode(soundcloudUrls)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func addSoundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var soundcloudData AuthenticatedSoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if persistence.BearerTokenExists(soundcloudData.BearerToken, persistence.SqlLiteDatabaseFileName) {
		persistence.AddSoundcloudUrl(soundcloudData.Url, persistence.SqlLiteDatabaseFileName)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func deleteSoundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var soundcloudData AuthenticatedSoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if persistence.BearerTokenExists(soundcloudData.BearerToken, persistence.SqlLiteDatabaseFileName) {
		persistence.DeleteSoundcloudUrl(soundcloudData.Url, persistence.SqlLiteDatabaseFileName)
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

// Checks if provided username / password exists.
// Successful authentication returns a 201 and the created bearer token.
// Failed authentication returns a 401 Status Unauthorized
func loginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userCredentials User
	err := decoder.Decode(&userCredentials)
	if err != nil {
		println("Error decoding user credentials from client")
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := persistence.User{Username: userCredentials.Username, Password: userCredentials.Password}
	userExists := persistence.UserExists(user, persistence.SqlLiteDatabaseFileName)
	if userExists {
		key := NewSHA1Hash()
		var bearerToken BearerToken
		bearerToken.BearerToken = key
		persistence.UpdateUserBearerToken(userCredentials.Username, userCredentials.Password, key, persistence.SqlLiteDatabaseFileName)
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(bearerToken)
		if err != nil {
			println("error encoding bearer token")
			return
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
