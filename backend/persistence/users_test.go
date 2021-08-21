package persistence

import (
	"fmt"
	"os"
	"testing"
)

const testDatabaseFile = "testDatabase.db"

func cleanUp() {
	err := os.Remove(testDatabaseFile)
	if err != nil {
		fmt.Printf("failed to delete %s", testDatabaseFile)
		return
	}
}

func TestCreateUserTable(t *testing.T) {
	defer cleanUp()
	createDatabase(testDatabaseFile)
	createUserTable(testDatabaseFile)
	tables := getAllTables(testDatabaseFile)
	if tables != "userCredentials" {
		t.Logf("userCredentials table does not exist %s", tables)
		t.Fail()
	}
	return
}

func TestCreateUser(t *testing.T) {
	defer cleanUp()
	createDatabase(testDatabaseFile)
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
	if users[0].Username != username {
		t.Log("username not returned")
		t.Fail()
	}
	if users[0].Password != password {
		t.Log("password not returned")
		t.Fail()
	}
}


