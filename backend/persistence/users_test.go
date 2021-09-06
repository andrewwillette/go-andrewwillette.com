package persistence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const testDatabaseFile = "testDatabase.db"

type UsersTestSuite struct {
	suite.Suite
}

func deleteDatabase() {
	os.Remove(testDatabaseFile)
}

func (suite *UsersTestSuite) SetupTest() {
	deleteDatabase()
}

func (suite *UsersTestSuite) TearDownSuite() {
	deleteDatabase()
}

func (suite *UsersTestSuite)TestCreateUserTable() {
	//deleteTestDatabase()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		suite.T().Fail()
	}
	assert.Equal(suite.T(), tables[0], "userCredentials")
}

func (suite *UsersTestSuite)TestCreateUser_Valid() {
	//deleteTestDatabase()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	username := "usernameOne"
	password := "passwordOne"
	err := userService.addUser(username, password)
	if err != nil {
		println(err.Error())
		suite.T().Logf("failed to add user")
		suite.T().Fail()
	}
	users, err := userService.getAllUsers()
	assert.Equal(suite.T(), users[0].Username, username)
	assert.Equal(suite.T(), users[0].Password, password)
}

func (suite *UsersTestSuite)TestUpdateUserBearerToken_Valid() {
	//deleteTestDatabase()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	username := "usernameOne"
	password := "passwordOne"
	err := userService.addUser(username, password)
	if err != nil {
		suite.T().Logf("failed to add user")
		suite.T().Fail()
	}
	bearerToken := "bearerTokenOne"
	userService.updateUserBearerToken(username, password, bearerToken)
	user, err := userService.getUser(username, password)
	if err != nil {
		suite.T().Log(err)
		suite.T().Fail()
	}
	assert.Equal(suite.T(), user.Username, username)
	assert.Equal(suite.T(), user.Password, password)
	assert.Equal(suite.T(), user.BearerToken, bearerToken)

	userExists, err := userService.userExists(&User{Username: username, Password: password})
	if err != nil {
		suite.T().Log(err)
		suite.T().Fail()
	}
	assert.True(suite.T(), userExists)

	bearerTokenExists := userService.BearerTokenExists(bearerToken)
	assert.True(suite.T(), bearerTokenExists)
}

func TestUsersSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}