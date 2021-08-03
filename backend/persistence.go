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
		"\"sessionKey\" BLOB" +
		")", userTable)
	statement, err := sqliteDatabase.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func addUserCredentials(username string, password string, sessionKey []byte) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	addUserWithSessionKey := fmt.Sprintf("INSERT INTO %s(username, password, sessionKey) VALUES (%s, %s, %s)", userTable, username, password, sessionKey)
	addUserWithSessionKeyStatement, err := db.Prepare(addUserWithSessionKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = addUserWithSessionKeyStatement.Exec(username, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func addUserSessionKey(username, password string, sessionKey []byte) {

}

func userCredentialsExists(credentials UserCredentials) bool {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	userExistsStatement := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE username = "%s" AND password = "%s")`, userTable, credentials.Username, credentials.Password)
	println(userExistsStatement)
	preparedStatement, err := db.Prepare(userExistsStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//result, err := preparedStatement.Exec(sql.Named("username", credentials.Username), sql.Named("password", credentials.Password))
	rows, err := preparedStatement.Query()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	var success string
	for rows.Next() {
		err := rows.Scan(&success)
		if err != nil {
			log.Fatalln(err.Error())
		}
		break
	}
	if success == "1" {
		return true
	} else {
		return false
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

func getAdminUsernamePassword() UserCredentials {
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
	var credentialModel UserCredentials
	for row.Next() { // Iterate and fetch the records from result cursor
		var username string
		var password string
		err := row.Scan(&username, &password)
		if err != nil {
			log.Fatalln(err.Error())
		}
		credentialModel = UserCredentials{Username: username, Password: password}
		break
	}
	return credentialModel
}