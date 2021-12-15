package main

import (
	"github.com/andrewwillette/willette_api/logging"
	"github.com/andrewwillette/willette_api/persistence"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"

func main() {
	logging.GlobalLogger.Info().Msg("Starting application")
	persistence.InitDatabaseIdempotent(SqlLiteDatabaseFileName)
	userService := &persistence.UserService{SqliteDbFile: SqlLiteDatabaseFileName}
	soundcloudUrlService := &persistence.SoundcloudUrlService{Sqlite: SqlLiteDatabaseFileName}
	server := NewWilletteAPIServer(userService, soundcloudUrlService)
	server.runServer()
}
