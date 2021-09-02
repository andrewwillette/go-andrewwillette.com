package main

import (
	"encoding/json"
	"github.com/andrewwillette/willette_api/persistence"
	"net/http"
)


type persistenceSoundcloudUrlService struct {}

func main() {
	persistence.InitDatabaseIdempotent()
	userService := &persistence.UserService{}
	server := NewWilletteAPIServer(userService, persistenceSoundcloudUrlService)
	runServer()
}
