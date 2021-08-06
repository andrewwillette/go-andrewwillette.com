package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"

// InitDatabase Creates database if it doesn't currently exist.
func InitDatabase() {
	if _, err := os.Stat(SqlLiteDatabaseFileName); os.IsNotExist(err) {
		log.Println("Creating sqlite-database.db...")
		file, err := os.Create(SqlLiteDatabaseFileName) // Create SQLite file
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
		sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
		defer sqliteDatabase.Close()
		//createSoundcloudUrlTable(sqliteDatabase)
	}
}

