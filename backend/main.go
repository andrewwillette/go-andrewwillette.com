package main

import (
	"github.com/andrewwillette/willette_api/persistence"
)

// local sqlite database filename
const SqlLiteDatabaseFileName = "sqlite-database.db"

func main() {
	persistence.InitDatabaseIdempotent(SqlLiteDatabaseFileName)
	userService := &persistence.UserService{Sqlite: SqlLiteDatabaseFileName}
	soundcloudUrlService := &persistence.SoundcloudUrlService{Sqlite: SqlLiteDatabaseFileName}
	server := NewWilletteAPIServer(userService, soundcloudUrlService)
	server.runServer()
}
