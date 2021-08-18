package main

import "github.com/andrewwillette/willette_api/persistence"

func main() {
	persistence.InitDatabaseIdempotent()
	runServer()
}
