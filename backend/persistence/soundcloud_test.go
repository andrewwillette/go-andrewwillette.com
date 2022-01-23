package persistence

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SoundcloudTestSuite struct {
	suite.Suite
}

func TestSoundcloudSuite(t *testing.T) {
	suite.Run(t, new(SoundcloudTestSuite))
}

func (suite *SoundcloudTestSuite) SetupTest() {
	deleteTestDatabase()
}

func (suite *SoundcloudTestSuite) TearDownSuite() {
	deleteTestDatabase()
}

func (suite *SoundcloudTestSuite) TestSoundcloudUrlService() {
	soundcloudUrlService := &SoundcloudUrlService{SqliteFile: testDatabaseFile}
	soundcloudUrlService.createSoundcloudUrlTable()
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		suite.T().Fail()
	}
	assert.Contains(suite.T(), tables, soundcloudTable)
	soundcloudUrlOne := "soundcloud.com/example"
	err = soundcloudUrlService.AddSoundcloudUrl(soundcloudUrlOne)
	if err != nil {
		suite.T().Fail()
	}
	soundcloudUrls, err := soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		suite.T().Fail()
	}
	assert.True(suite.T(), soundcloudUrlExists(soundcloudUrls, soundcloudUrlOne, 0))
	soundcloudUrlTwo := "soundcloud.com/numbertwo"
	err = soundcloudUrlService.AddSoundcloudUrl(soundcloudUrlTwo)
	if err != nil {
		suite.T().Fail()
		return
	}
	soundcloudUrls, err = soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		suite.T().Fail()
	}
	assert.True(suite.T(), soundcloudUrlExists(soundcloudUrls, soundcloudUrlTwo, 0))
	newUiOrderOne := SoundcloudUrl{Url: soundcloudUrlOne, UiOrder: 23}
	newUiOrderTwo := SoundcloudUrl{Url: soundcloudUrlTwo, UiOrder: 5}
	//newUiOrderTwo := 5
	err = soundcloudUrlService.UpdateSoundcloudUiOrders([]SoundcloudUrl{newUiOrderTwo, newUiOrderOne})
	//err = soundcloudUrlService.UpdateSoundcloudUrls_uiOrder(soundcloudUrlTwo, newUiOrderTwo)
	if err != nil {
		suite.T().Fail()
	}
	soundcloudUrls, err = soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		suite.T().Fail()
	}
	assert.True(suite.T(), soundcloudUrlExists(soundcloudUrls, newUiOrderOne))
	assert.True(suite.T(), soundcloudUrlExists(soundcloudUrls, newUiOrderTwo))
	assert.ElementsMatch(suite.T(), soundcloudUrls, []SoundcloudUrl{newUiOrderOne, newUiOrderTwo})
	//assert.True(suite.T(), soundcloudUrlExists(soundcloudUrls, soundcloudUrlTwo, newUiOrderTwo))
}

func soundcloudUrlExists(soundcloudUrls []SoundcloudUrl, url SoundcloudUrl) bool {
	for _, value := range soundcloudUrls {
		if value.Url == url && value.UiOrder == uiOrder {
			return true
		}
	}
	return false
}
