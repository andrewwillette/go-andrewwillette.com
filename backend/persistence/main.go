package persistence

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// InitDatabaseIdempotent - Creates database if it doesn't currently exist.
func InitDatabaseIdempotent() {
	if _, err := os.Stat(SqlLiteDatabaseFileName); os.IsNotExist(err) {
		file, err := os.Create(SqlLiteDatabaseFileName) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
		defer sqliteDatabase.Close()
		createSoundcloudUrlTable(sqliteDatabase)
		createUserTable(sqliteDatabase)
	}
}
