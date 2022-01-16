package persistence

import (
	"database/sql"
	"fmt"
	"github.com/andrewwillette/willette_api/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

const testDatabaseFile = "testDatabase.db"

type PersistenceTestSuite struct {
	suite.Suite
}

func TestPersistenceSuite(t *testing.T) {
	suite.Run(t, new(PersistenceTestSuite))
}

func (suite *PersistenceTestSuite) SetupTest() {
	deleteDatabase()
}

func (suite *PersistenceTestSuite) TearDownSuite() {
	deleteDatabase()
}

func deleteDatabase() {
	_ = os.Remove(testDatabaseFile)
}

func TestCreateDatabase(t *testing.T) {
	err := createDatabase(testDatabaseFile)
	if err != nil {
		logging.TestLogger.Err(err)
		t.Fail()
	}
	_, err = sql.Open("sqlite3", fmt.Sprintf("./%s", testDatabaseFile))
	if err != nil {
		logging.TestLogger.Err(err)
		t.Fail()
	}
	assert.FileExists(t, testDatabaseFile)
}
