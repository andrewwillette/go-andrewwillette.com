package persistence

import (
	"database/sql"
	"fmt"
	"github.com/andrewwillette/willette_api/models"
	"log"
)

const userTable = "userCredentials"

func CreateUserTable() {
	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer sqliteDatabase.Close()
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s ("+
		"\"username\" TEXT NOT NULL, "+
		"\"password\" TEXT NOT NULL, "+
		"\"bearerToken\" BLOB"+
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

/*
To be set after logging in, bearer token
*/
func UpdateUserBearerToken(username string, password, bearerToken string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	addUserWithSessionKey := fmt.Sprintf("UPDATE %s SET bearerToken = '%s' WHERE username = '%s'",
		userTable, bearerToken, username)
	addUserWithSessionKeyStatement, err := db.Prepare(addUserWithSessionKey)
	if err != nil {
		return
	}
	_, err = addUserWithSessionKeyStatement.Exec()
	if err != nil {
		println("error executing bearer token update statement")
		println(err.Error())
		return
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
	if success == "1" {
		return true
	} else {
		return false
	}
}

func BearerTokenExists(bearerToken string) bool {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer db.Close()
	bearerTokenExists := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE bearerToken = "%s")`, userTable, bearerToken)
	preparedStatement, err := db.Prepare(bearerTokenExists)
	if err != nil {
		println("Error parsing ")
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
	if success == "1" {
		return true
	} else {
		return false
	}
	return true
}
