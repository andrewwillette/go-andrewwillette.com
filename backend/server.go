package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage")
	fmt.Println("endpoint hit: homepage")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


// response "hello world"
func handleNewRequest(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusCreated)
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	resp := make(map[string]string)
	resp["song_one"] = "Status out out"
	resp["song_two"] = "Swag"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error occurred in JSON marshal. Err: %s", err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		log.Fatalf("Error occurred when writing response Err: %s", err)
		return
	}
	return
}

type soundcloudGet struct {
	Urls []string
}

func soundcloudUrlsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	soundcloudUrls := soundcloudGet {
		Urls: getAllSoundcloudUrls(),
	}

	json.NewEncoder(w).Encode(soundcloudUrls)
}

type soundcloudPost struct {
	Url string `json:"url"`
}

func soundcloudUrlPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var soundcloudData soundcloudPost
	err := decoder.Decode(&soundcloudData)
	if err != nil {
		log.Fatalf(err.Error())
	}
	addSoundcloudUrlDb(soundcloudData.Url)
}

func runServer() {
	handler := http.HandlerFunc(soundcloudUrlsGet)
	http.Handle("/get-soundcloud-urls", handler)

	soundcloudHandler := http.HandlerFunc(soundcloudUrlPost)
	http.Handle("/add-soundcloud", soundcloudHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to listen and serve to port 8080")
		return
	}
}
