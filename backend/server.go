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

func runServer() {
	getHandler := http.HandlerFunc(soundcloudUrlsGet)
	http.Handle(getUrl, getHandler)

	putHandler := http.HandlerFunc(soundcloudUrlPost)
	http.Handle(putUrl, putHandler)

	deleteHandler := http.HandlerFunc(soundcloudUrlDelete)
	http.Handle(deleteUrl, deleteHandler)

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
	json.NewEncoder(w).Encode(toReturn)
}

func soundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var soundcloudData models.SoundcloudUrl
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	addSoundcloudUrlDb(soundcloudData.Url)
}

func soundcloudUrlDelete(w http.ResponseWriter, r *http.Request) {

}
