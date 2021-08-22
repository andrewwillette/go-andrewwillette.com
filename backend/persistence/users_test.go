package persistence

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const testDatabaseFile = "testDatabase.db"

func deleteTestDatabase() {
	os.Remove(testDatabaseFile)
}

func TestCreateUserTable(t *testing.T) {
	deleteTestDatabase()
	createUserTable(testDatabaseFile)
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, tables[0], "userCredentials")
}

func TestCreateUser_Valid(t *testing.T) {
	deleteTestDatabase()
	createUserTable(testDatabaseFile)
	username := "usernameOne"
	password := "passwordOne"
	err := AddUser(username, password, testDatabaseFile)
	if err != nil {
		println(err.Error())
		t.Logf("failed to add user")
		t.Fail()
	}
	users, err := getAllUsers(testDatabaseFile)
	assert.Equal(t, users[0].Username, username)
	assert.Equal(t, users[0].Password, password)
}

func TestUpdateUserBearerToken_Valid(t *testing.T) {
	deleteTestDatabase()
	createUserTable(testDatabaseFile)
	username := "usernameOne"
	password := "passwordOne"
	err := AddUser(username, password, testDatabaseFile)
	if err != nil {
		t.Logf("failed to add user")
		t.Fail()
	}
	bearerToken := "bearerTokenOne"
	UpdateUserBearerToken(username, password, bearerToken, testDatabaseFile)
	user, err := GetUser(username, password, testDatabaseFile)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Password, password)
	assert.Equal(t, user.BearerToken, bearerToken)

	userExists := UserExists(User{Username: username, Password: password}, testDatabaseFile)
	assert.True(t, userExists)

	bearerTokenExists := BearerTokenExists(bearerToken, testDatabaseFile)
	assert.True(t, bearerTokenExists)
}