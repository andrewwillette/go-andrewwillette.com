package persistence

import (
	"database/sql"
	"fmt"
	"github.com/andrewwillette/willette_api/models"
	"log"
)

const userTable = "userCredentials"

func createUserTable(sqliteFile string) {
	db, err := sql.Open("sqlite3", sqliteFile)
	defer db.Close()
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s ("+
		"\"username\" TEXT NOT NULL, "+
		"\"password\" TEXT NOT NULL, "+
		"\"bearerToken\" BLOB"+
		")", userTable)
	statement, err := db.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func AddUser(username, password, sqliteFile string) error {
	//"INSERT INTO userCredentials(username, password) VALUES('$username', '$password');"
	addUserSqlStatement := fmt.Sprintf("INSERT INTO %s(username, password) "+
		"VALUES('%s', '%s');", userTable, username, password)
	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		return err
	}
	addUserStatement, err := db.Prepare(addUserSqlStatement)
	if err != nil {
		return err
	}
	_, err = addUserStatement.Exec()
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserBearerToken Adds the provided bearerToken to the username/password
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

func getAllUsers(databaseFile string) ([]models.User, error) {
	db, err := sql.Open("sqlite3", databaseFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	selectAllUsers := fmt.Sprintf("SELECT * FROM %s", userTable)
	selectAllUsersPrepared, err := db.Prepare(selectAllUsers)
	if err != nil {
		return nil, err
	}
	rows, err := selectAllUsersPrepared.Query()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	var username, password, bearerToken sql.NullString
	var users []models.User
	for rows.Next() {
		err := rows.Scan(&username, &password, &bearerToken)
		user := models.User{Username: username.String, Password: password.String}
		users = append(users, user)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
	return users, nil
}

// UserExists Checks database if username, password exists
func UserExists(credentials models.User) bool {
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
