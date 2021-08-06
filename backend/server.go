package main

import (
	"willette_site/models"
	"willette_site/persistence"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const getUrl = "/get-soundcloud-urls"
const putUrl = "/add-soundcloud-url"
const deleteUrl = "/delete-soundcloud-url"
const loginUrl = "/login"

func runServer() {
	getHandler := http.HandlerFunc(soundcloudUrlsGet)
	http.Handle(getUrl, getHandler)

	putHandler := http.HandlerFunc(soundcloudUrlPost)
	http.Handle(putUrl, putHandler)

	deleteHandler := http.HandlerFunc(soundcloudUrlDelete)
	http.Handle(deleteUrl, deleteHandler)

	loginHandler := http.HandlerFunc(loginPost)
	http.Handle(loginUrl, loginHandler)

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

func loginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userCredentials models.UserCredentials
	err := decoder.Decode(&userCredentials)
	if err != nil {
		log.Fatalf(err.Error())
	}
	userExists := persistence.UserCredentialsExists(userCredentials)
	if userExists {
		key := NewSHA1Hash()
		var sessionKey models.SessionKey
		sessionKey.SessionKey = key
		persistence.UpdateUserSessionKey(userCredentials.Username, userCredentials.Password, key)
		err := json.NewEncoder(w).Encode(sessionKey)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	fmt.Println(userExists)
}

