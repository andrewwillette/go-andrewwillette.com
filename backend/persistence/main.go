package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Creates Sqlite database with filename
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

func getAllTables(databaseFile string) ([]string, error) {
	sqliteDatabase, _ := sql.Open("sqlite3", fmt.Sprintf("./%s", databaseFile))
	rows, err := sqliteDatabase.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil, err
	}
	var table string
	var tables []string
	for rows.Next() {
		err := rows.Scan(&table)
		tables = append(tables, table)
		if err != nil {
			return nil, err
		}
	}
	return tables, nil
}

// InitDatabaseIdempotent - Creates database if it doesn't currently exist.
func InitDatabaseIdempotent(sqlite string) {
	if _, err := os.Stat(sqlite); os.IsNotExist(err) {
		err = createDatabase(sqlite)
		if err != nil {
			panic("failed to create database")
		}

		userService := &UserService{SqliteDbFile: sqlite}
		userService.createUserTable()

		soundcloudUrlService := &SoundcloudUrlService{Sqlite: sqlite}
		soundcloudUrlService.createSoundcloudUrlTable()
	}
}
