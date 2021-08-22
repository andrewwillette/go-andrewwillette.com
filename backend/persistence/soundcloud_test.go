package persistence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateSoundcloudUrlTable(t *testing.T){
	deleteTestDatabase()
	createSoundcloudUrlTable(testDatabaseFile)
	createUserTable(testDatabaseFile)
	tables, err := getAllTables(testDatabaseFile)
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, tables[0], "soundcloudUrl")
	soundcloudUrl := "soundcloud.com/example"
	AddSoundcloudUrl(soundcloudUrl, testDatabaseFile)
	soundcloudUrls := GetAllSoundcloudUrls(testDatabaseFile)
	assert.Contains(t, soundcloudUrls, soundcloudUrl)

	soundcloudUrlTwo := "soundcloud.com/numbertwo"
	AddSoundcloudUrl(soundcloudUrlTwo, testDatabaseFile)
	soundcloudUrls = GetAllSoundcloudUrls(testDatabaseFile)
	assert.Contains(t, soundcloudUrls, soundcloudUrlTwo)
}

func TestAddSoundcloudUrl(t *testing.T) {
	deleteTestDatabase()
	createSoundcloudUrlTable(testDatabaseFile)
}