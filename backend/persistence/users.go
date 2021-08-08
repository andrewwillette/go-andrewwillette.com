package persistence

import (
	"github.com/andrewwillette/willette_api/models"
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
		"\"bearerToken\" BLOB" +
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
	fmt.Printf("sql to prepare\n%s\n", addUserWithSessionKey)
	addUserWithSessionKeyStatement, err := db.Prepare(addUserWithSessionKey)
	if err != nil {
		println("in business")
		println("error preparing bearer token update statement")
		return
	}
	result, err := addUserWithSessionKeyStatement.Exec()
	fmt.Println(result)
	if err != nil {
		println("error executing bearer token update statement")
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
	fmt.Printf("success is %s\n", success)
	if success == "1" {
		return true
	} else {
		return false
	}
}

