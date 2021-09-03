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
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, tables[0], "userCredentials")
}

func TestCreateUser_Valid(t *testing.T) {
	deleteTestDatabase()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	username := "usernameOne"
	password := "passwordOne"
	err := userService.addUser(username, password)
	if err != nil {
		println(err.Error())
		t.Logf("failed to add user")
		t.Fail()
	}
	users, err := userService.getAllUsers()
	assert.Equal(t, users[0].Username, username)
	assert.Equal(t, users[0].Password, password)
}

func TestUpdateUserBearerToken_Valid(t *testing.T) {
	deleteTestDatabase()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	username := "usernameOne"
	password := "passwordOne"
	err := userService.addUser(username, password)
	if err != nil {
		t.Logf("failed to add user")
		t.Fail()
	}
	bearerToken := "bearerTokenOne"
	userService.updateUserBearerToken(username, password, bearerToken)
	user, err := userService.getUser(username, password)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	assert.Equal(t, user.Username, username)
	assert.Equal(t, user.Password, password)
	assert.Equal(t, user.BearerToken, bearerToken)

	userExists, err := userService.userExists(&User{Username: username, Password: password})
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	assert.True(t, userExists)

	bearerTokenExists := userService.BearerTokenExists(bearerToken)
	assert.True(t, bearerTokenExists)
}