package main

import (
	"encoding/json"
	"github.com/andrewwillette/willette_api/models"
	"github.com/andrewwillette/willette_api/persistence"
	"log"
	"net/http"
)

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const deleteSoundcloudEndpoint = "/delete-soundcloud-url"
const loginEndpoint = "/login"

func runServer() {
	getHandler := http.HandlerFunc(soundcloudUrlsGet)
	http.Handle(getSoundcloudAllEndpoint, getHandler)

	putHandler := http.HandlerFunc(soundcloudUrlPost)
	http.Handle(addSoundcloudEndpoint, putHandler)

	deleteHandler := http.HandlerFunc(soundcloudUrlDelete)
	http.Handle(deleteSoundcloudEndpoint, deleteHandler)

	loginHandler := http.HandlerFunc(loginPost)
	http.Handle(loginEndpoint, loginHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to listen and serve to port 8080")
		return
	}
}

func soundcloudUrlsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	urls := persistence.GetAllSoundcloudUrls()
	var soundcloudUrls []models.SoundcloudUrl
	for i := 0; i < len(urls); i++ {
		soundcloudUrls = append(soundcloudUrls, models.SoundcloudUrl{Url: urls[i]})
	}
	err := json.NewEncoder(w).Encode(soundcloudUrls)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func soundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var soundcloudData models.SoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	persistence.AddSoundcloudUrl(soundcloudData.Url)
}

func soundcloudUrlDelete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var soundcloudData models.SoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	persistence.DeleteSoundcloudUrlDb(soundcloudData.Url)
}

/**
Logs user in, returns generated bearer token
 */
func loginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userCredentials models.UserCredentials
	err := decoder.Decode(&userCredentials)
	if err != nil {
		println("Error decoding user credentials from client")
		return
	}
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	userExists := persistence.UserCredentialsExists(userCredentials)
	if userExists {
		key := NewSHA1Hash()
		var bearerToken models.BearerToken
		bearerToken.BearerToken = key
		persistence.UpdateUserBearerToken(userCredentials.Username, userCredentials.Password, key)
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(bearerToken)
		if err != nil {
			println("error encoding bearer token")
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write(nil)
	}
}
