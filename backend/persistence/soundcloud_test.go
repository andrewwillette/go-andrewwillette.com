package persistence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSoundcloudUrlTable(t *testing.T){
	deleteTestDatabase()
	soundcloudUrlService := &SoundcloudUrlService{Sqlite: testDatabaseFile}
	soundcloudUrlService.createSoundcloudUrlTable()
	userService := &UserService{Sqlite: testDatabaseFile}
	userService.createUserTable()
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, tables[0], "soundcloudUrl")
	soundcloudUrl := "soundcloud.com/example"
	soundcloudUrlService.AddSoundcloudUrl(soundcloudUrl)
	soundcloudUrls, err := soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		t.Fail()
	}
	assert.Contains(t, soundcloudUrls, soundcloudUrl)

	soundcloudUrlTwo := "soundcloud.com/numbertwo"
	soundcloudUrlService.AddSoundcloudUrl(soundcloudUrlTwo)
	soundcloudUrls, err = soundcloudUrlService.GetAllSoundcloudUrls()
	if err != nil {
		t.Fail()
	}
	assert.Contains(t, soundcloudUrls, soundcloudUrlTwo)
}

//func TestAddSoundcloudUrl(t *testing.T) {
//	deleteTestDatabase()
//	createSoundcloudUrlTable(testDatabaseFile)
//}