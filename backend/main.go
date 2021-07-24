package main

import (
	"./models"
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

/**
basic response with test values in json form
 */
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

func handleGetSoundcloudUrls(w http.ResponseWriter, r *http.Request) {
	soundcloudUrls := models.SoundCloudUpload{
		Url: "soundcloud.com",
	}
	json.NewEncoder(w).Encode(soundcloudUrls)
}

func main() {
	//handleRequests()
	handler := http.HandlerFunc(handleGetSoundcloudUrls)
	http.Handle("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to listen and serve to port 8080")
		return
	}
}
