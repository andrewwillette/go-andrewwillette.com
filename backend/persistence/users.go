package persistence

import (
	"willette_site/models"
	"database/sql"
	"fmt"
	"log"
)

const userTable = "userCredentials"

func CreateUserTable() {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqliteDatabase.Close()
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s (" +
		"\"username\" TEXT NOT NULL, " +
		"\"password\" TEXT NOT NULL, " +
		"\"sessionKey\" BLOB" +
		")", userTable)
	println(createSoundcloudTableSQL)
	statement, err := sqliteDatabase.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

/*
To be set after logging in, sessionKey is bearer token of sorts
*/
func UpdateUserSessionKey(username string, password, bearerToken string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	addUserWithSessionKey := fmt.Sprintf("UPDATE %s SET bearerToken = %s WHERE username = %s AND password = %s)",
		userTable, bearerToken, username, password)
	addUserWithSessionKeyStatement, err := db.Prepare(addUserWithSessionKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = addUserWithSessionKeyStatement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

/**
Checks database if username, password exists
*/
func UserCredentialsExists(credentials models.UserCredentials) bool {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	userExistsStatement := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE username = "%s" AND password = "%s")`, userTable, credentials.Username, credentials.Password)
	preparedStatement, err := db.Prepare(userExistsStatement)
	if err != nil {
		log.Fatalln(err.Error())
	}
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
	fmt.Printf("success is %s\n", success)
	if success == "1" {
		return true
	} else {
		return false
	}
}

