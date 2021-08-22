package persistence

import (
	"database/sql"
	"fmt"
	"log"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"
const soundcloudTable = "soundcloudUrl"

func AddSoundcloudUrl(url, sqliteFile string) {
	db, err := sql.Open("sqlite3", sqliteFile)
	defer db.Close()
	insertUrlStatement := fmt.Sprintf("INSERT INTO %s(url) VALUES (?)", soundcloudTable)
	addSoundcloudPreparedStatement, err := db.Prepare(insertUrlStatement) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = addSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func DeleteSoundcloudUrl(url, sqliteFile string) error {
	db, err := sql.Open("sqlite3", sqliteFile)
	defer db.Close()
	deleteSoundcloudStatement := fmt.Sprintf("DELETE FROM %s WHERE url = (?)", soundcloudTable)
	deleteSoundcloudPreparedStatement, err := db.Prepare(deleteSoundcloudStatement) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		return err
	}
	_, err = deleteSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		return err
	}
	return nil
}

// Creates database table for soundcloudUrls
func createSoundcloudUrlTable(sqliteFile string) {
	db, err := sql.Open("sqlite3", sqliteFile)
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s ("+
		"\"id\" integer NOT NULL PRIMARY KEY AUTOINCREMENT,"+
		"\"url\" TEXT"+
		")", soundcloudTable)
	statement, err := db.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// GetAllSoundcloudUrls get all soundcloud urls in database
func GetAllSoundcloudUrls(sqliteFile string) []string {
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	selectAllSoundcloudStatement := fmt.Sprintf("SELECT * FROM %s;", soundcloudTable)
	row, err := db.Query(selectAllSoundcloudStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer row.Close()
	var soundcloudUrls []string
	for row.Next() {
		var url string
		var urlTwo string
		err := row.Scan(&url, &urlTwo)
		if err != nil {
			log.Fatalln(err.Error())
		}
		soundcloudUrls = append(soundcloudUrls, urlTwo)
	}
	return soundcloudUrls
}
