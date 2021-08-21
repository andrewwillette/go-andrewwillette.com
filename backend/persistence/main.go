package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

// Creates sqlite database with filename
func createDatabase(databaseFile string) error {
	file, err := os.Create(databaseFile)
	if err != nil {
		return err
	}
	file.Close()
	sqliteDatabase, err := sql.Open("sqlite3", fmt.Sprintf("./%s", databaseFile))
	if err != nil {
		return err
	}
	err = sqliteDatabase.Close()
	if err != nil {
		return err
	}
	return nil
}

func getAllTables(databaseFile string) string {
	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", databaseFile))
	row := sqliteDatabase.QueryRow("SELECT name FROM sqlite_master WHERE type='table';")
	var success string
	err := row.Scan(&success)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return success
}

// InitDatabaseIdempotent - Creates database if it doesn't currently exist.
func InitDatabaseIdempotent() {
	if _, err := os.Stat(SqlLiteDatabaseFileName); os.IsNotExist(err) {
		err = createDatabase(SqlLiteDatabaseFileName)
		if err != nil {
			panic("failed to create database")
		}
		createSoundcloudUrlTable(SqlLiteDatabaseFileName)
		createUserTable(SqlLiteDatabaseFileName)
	}
}
