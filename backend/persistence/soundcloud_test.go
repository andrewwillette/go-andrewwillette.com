package persistence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SoundcloudTestSuite struct {
	suite.Suite
}

func (suite *SoundcloudTestSuite) SetupTest() {
	deleteDatabase()
}

func (suite *SoundcloudTestSuite) TearDownSuite() {
	deleteDatabase()
}

func (suite *SoundcloudTestSuite) TestCreateSoundcloudUrlTable() {
	soundcloudUrlService := &SoundcloudUrlService{Sqlite: testDatabaseFile}
	soundcloudUrlService.createSoundcloudUrlTable()
	userService := &UserService{SqliteDbFile: testDatabaseFile}
	userService.createUserTable()
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		suite.T().Fail()
	}
	assert.Equal(suite.T(), tables[0], "soundcloudUrl")
	soundcloudUrl := "soundcloud.com/example"
	soundcloudUrlService.AddSoundcloudUrl(soundcloudUrl)
	soundcloudUrls, err := soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		suite.T().Fail()
	}
	assert.Contains(suite.T(), soundcloudUrls, soundcloudUrl)

	soundcloudUrlTwo := "soundcloud.com/numbertwo"
	soundcloudUrlService.AddSoundcloudUrl(soundcloudUrlTwo)
	soundcloudUrls, err = soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		suite.T().Fail()
	}
	assert.Contains(suite.T(), soundcloudUrls, soundcloudUrlTwo)
}

func TestSoundcloudSuite(t *testing.T) {
	suite.Run(t, new(SoundcloudTestSuite))
}
