package persistence

import (
	"database/sql"
	"fmt"
	"github.com/andrewwillette/willette_api/logging"
)

const soundcloudTable = "soundcloudUrl"

type SoundcloudUrlService struct {
	SqliteFile string
}

type SoundcloudUrl struct {
	Id      int
	Url     string
	UiOrder int
}

func (u *SoundcloudUrlService) UpdateSoundcloudUrlUiOrder(url string, uiOrder int) error {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in UpdateSoundcloudUrlUiOrder")
		return err
	}
	updateUrlStatement := fmt.Sprintf("UPDATE %s SET uiOrder = %d WHERE url = \"%s\"", soundcloudTable, uiOrder, url)
	preparedStatement, err := db.Prepare(updateUrlStatement)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error preparing sql in UpdateSoundcloudUrlUiOrder")
		return err
	}
	_, err = preparedStatement.Exec()
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing sql in UpdateSoundcloudUrlUiOrder")
		return err
	}
	return nil
}

func (u *SoundcloudUrlService) AddSoundcloudUrl(url string) error {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in AddSoundcloudUrl")
		return err
	}
	defer db.Close()
	insertUrlStatement := fmt.Sprintf("INSERT INTO %s(url) VALUES (?)", soundcloudTable)
	addSoundcloudPreparedStatement, err := db.Prepare(insertUrlStatement)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error preparing add soundcloud url sql query")
		return err
	}
	_, err = addSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing add soundcloud url sql")
		return err
	}
	return nil
}

func (u *SoundcloudUrlService) DeleteSoundcloudUrl(url string) error {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in DeleteSoundcloudUrl")
		return err
	}
	defer db.Close()
	deleteSoundcloudStatement := fmt.Sprintf("DELETE FROM %s WHERE url = (?)", soundcloudTable)
	deleteSoundcloudPreparedStatement, err := db.Prepare(deleteSoundcloudStatement)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error preparing delete soundcloud url sql")
		return err
	}
	_, err = deleteSoundcloudPreparedStatement.Exec(url)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing delete soundcloud url sql")
		return err
	}
	return nil
}

func (u *SoundcloudUrlService) createSoundcloudUrlTable() {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in createSoundcloudUrl table")
		return
	}
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s ("+
		"\"id\" integer NOT NULL PRIMARY KEY AUTOINCREMENT,"+
		"\"url\" text,"+
		"\"uiOrder\" integer"+
		")", soundcloudTable)
	statement, err := db.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error preparing create soundCloud url table")
		return
	}
	_, err = statement.Exec()
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing create soundcloudurl table sql")
		return
	}
}

// GetAllSoundcloudUrls get all SoundcloudUrls in database
func (u *SoundcloudUrlService) GetAllSoundcloudUrls() ([]SoundcloudUrl, error) {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in GetAllSoundcloudUrls")
		return nil, err
	}
	defer db.Close()
	selectAllSoundcloudStatement := fmt.Sprintf("SELECT id, url, uiOrder FROM %s;", soundcloudTable)
	preparedStatement, err := db.Prepare(selectAllSoundcloudStatement)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Failed to prepare get all soundcloud sql statement.")
	}
	soundcloudUrlsArrayMap, err := getQueryResponseAsMap(preparedStatement)
	var soundcloudUrls []SoundcloudUrl
	for _, scUrl := range soundcloudUrlsArrayMap {
		var soundCloudUrl SoundcloudUrl
		if scUrl["url"] != nil {
			soundCloudUrl.Url = fmt.Sprint(scUrl["url"])
		}
		if scUrl["id"] != nil {
			soundCloudUrl.Id = int(scUrl["id"].(int64))
		}
		if scUrl["uiOrder"] != nil {
			soundCloudUrl.UiOrder = int(scUrl["uiOrder"].(int64))
		}
		soundcloudUrls = append(soundcloudUrls, soundCloudUrl)
	}
	return soundcloudUrls, nil
}
