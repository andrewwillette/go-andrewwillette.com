package main

import (
	"./models"
	"encoding/json"
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
	urls := getAllSoundcloudUrls()
	var toReturn []models.SoundcloudUrl
	for i := 0; i < len(urls); i++ {
		toReturn = append(toReturn, models.SoundcloudUrl{Url: urls[i]})
	}
	err := json.NewEncoder(w).Encode(toReturn)
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
	addSoundcloudUrl(soundcloudData.Url)
}

func soundcloudUrlDelete(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var soundcloudData models.SoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	deleteSoundcloudUrlDb(soundcloudData.Url)
}

func loginPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userCredentials models.UserCredentials
	err := decoder.Decode(&userCredentials)
	if err != nil {
		log.Fatalf(err.Error())
	}
	// how to handle checking if username / password are correct?
	println(userCredentials.Username)
	println(userCredentials.Password)
}
