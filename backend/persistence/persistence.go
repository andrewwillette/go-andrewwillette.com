package persistence

import (
	"database/sql"
	"fmt"
	"github.com/andrewwillette/willette_api/logging"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

// Creates SqliteFile database with filename
func createDatabase(databaseFile string) error {
	file, err := os.Create(databaseFile)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
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
// Database creation includes creating the user and soundcloud table.
func InitDatabaseIdempotent(sqliteFile string) {
	if _, err := os.Stat(sqliteFile); os.IsNotExist(err) {
		err = createDatabase(sqliteFile)
		if err != nil {
			logging.GlobalLogger.Fatal().Msg("failed to create database")
		}

		userService := &UserService{SqliteDbFile: sqliteFile}
		userService.createUserTable()

		soundcloudUrlService := &SoundcloudUrlService{SqliteFile: sqliteFile}
		soundcloudUrlService.createSoundcloudUrlTable()
	}
}
