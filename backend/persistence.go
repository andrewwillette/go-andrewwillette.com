package main

import (
	"./models"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"
const soundcloudTable = "soundcloudUrl"
const userTable = "userCredentials"

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
		createSoundcloudUrlTable(sqliteDatabase)
	}
}

func addSoundcloudUrl(url string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
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

func deleteSoundcloudUrlDb(url string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	deleteSoundcloudStatement := fmt.Sprintf("DELETE FROM %s WHERE url = (?)", soundcloudTable)
	deleteSoundcloudPreparedStatement, err := db.Prepare(deleteSoundcloudStatement) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = deleteSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createUserTable() {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqliteDatabase.Close()
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s (" +
		"\"username\" TEXT NOT NULL," +
		"\"password\" TEXT NOT NULL" +
		")", userTable)
	log.Println("Creating user credentials table.")
	statement, err := sqliteDatabase.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("User credential table created.")
}

func addUserCredentials(username string, password string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	insertUrlStatement := fmt.Sprintf("INSERT INTO %s(username, password) VALUES (?, ?)", soundcloudTable)
	addSoundcloudPreparedStatement, err := db.Prepare(insertUrlStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = addSoundcloudPreparedStatement.Exec(username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func createSoundcloudUrlTable(db *sql.DB) {
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s (" +
		"\"id\" integer NOT NULL PRIMARY KEY AUTOINCREMENT," +
		"\"url\" TEXT" +
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

// get all soundcloud urls in database
func getAllSoundcloudUrls() []string {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
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

func getAdminUsernamePassword() models.UserCredentials {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	selectCredentialsStatement := fmt.Sprintf("SELECT * FROM %s", userTable)
	row, err := db.Query(selectCredentialsStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer row.Close()
	var credentialModel models.UserCredentials
	for row.Next() { // Iterate and fetch the records from result cursor
		var username string
		var password string
		err := row.Scan(&username, &password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		credentialModel = models.UserCredentials{Username: username, Password: password}
		break
	}
	return credentialModel
}