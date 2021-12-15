package persistence

import (
	"database/sql"
	"fmt"

	"github.com/andrewwillette/willette_api/logging"
)

const soundcloudTable = "soundcloudUrl"

type SoundcloudUrlService struct {
	Sqlite string
}

func (u *SoundcloudUrlService) AddSoundcloudUrl(url string) error {
	db, err := sql.Open("sqlite3", u.Sqlite)
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
	db, err := sql.Open("sqlite3", u.Sqlite)
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
	db, err := sql.Open("sqlite3", u.Sqlite)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in createSoundcloudUrl table")
		return
	}
	createSoundcloudTableSQL := fmt.Sprintf("CREATE TABLE %s ("+
		"\"id\" integer NOT NULL PRIMARY KEY AUTOINCREMENT,"+
		"\"url\" TEXT"+
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

// GetAllSoundcloudUrls get all soundcloud urls in database
func (u *SoundcloudUrlService) GetAllSoundcloudUrls() ([]string, error) {
	db, err := sql.Open("sqlite3", u.Sqlite)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in GetAllSoundcloudUrls")
		return nil, err
	}
	defer db.Close()
	// TODO : prepare this statement
	selectAllSoundcloudStatement := fmt.Sprintf("SELECT * FROM %s;", soundcloudTable)
	row, err := db.Query(selectAllSoundcloudStatement)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing get all soundcloud url sql")
		return nil, err
	}
	defer row.Close()
	var soundcloudUrls []string
	for row.Next() {
		var url string
		var urlTwo string
		err := row.Scan(&url, &urlTwo)
		if err != nil {
			logging.GlobalLogger.Err(err).Msg("Error scanning sql row in get all soundcloud urls")
			return nil, err
		}
		soundcloudUrls = append(soundcloudUrls, urlTwo)
	}
	return soundcloudUrls, nil
}
