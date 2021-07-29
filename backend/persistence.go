package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"
const soundcloudTable = "soundcloudUrl"

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
		createTable(sqliteDatabase)
	}
}

func addSoundcloudUrlDb(url string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	insertSoundcloudSQL := `INSERT INTO soundcloudUrl(url) VALUES (?)`
	addSoundcloudPreparedStatement, err := db.Prepare(insertSoundcloudSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = addSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func deleteSoundcloudUrlDb(url string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	insertSoundcloudSQL := `DELETE FROM soundcloudUrl WHERE url = (?)`
	deleteSoundcloudPreparedStatement, err := db.Prepare(insertSoundcloudSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = deleteSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createTable(db *sql.DB) {
	createSoundcloudTableSQL := `CREATE TABLE soundcloudUrl (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url" TEXT
	 );`

	log.Println("Create soundcloud table...")
	println("got here 1")
	statement, err := db.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	println("got here 2")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
		return
	} // Execute SQL Statements
	log.Println("soundcloud table created")
}

// get all soundcloud urls in database
func getAllSoundcloudUrls() []string {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		fmt.Println("error reading sqlite database in getAllSoundcloudUrls")
		fmt.Println(err.Error())
		return nil
	}
	defer db.Close()
	row, err := db.Query("SELECT * FROM soundcloudUrl;")
	if err != nil {
		fmt.Println("error selecting all from soundcloudUrl table")
		fmt.Println(err.Error())
		return nil
	}
	defer row.Close()
	var soundcloudUrls []string
	for row.Next() { // Iterate and fetch the records from result cursor
		var url string
		var urlTwo string
		err := row.Scan(&url, &urlTwo)
		if err != nil {
			fmt.Println("error reading url")
			fmt.Println(err.Error())
			return nil
		}
		soundcloudUrls = append(soundcloudUrls, urlTwo)
		log.Println("SoundcloudUrl: ", urlTwo)
	}
	log.Println(soundcloudUrls)
	return soundcloudUrls
}
