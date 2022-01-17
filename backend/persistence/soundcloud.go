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

func (u *SoundcloudUrlService) AddSoundcloudUrl(url string) error {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in AddSoundcloudUrl")
		return err
	}
	defer db.Close()
	insertUrlStatement := fmt.Sprintf("INSERT INTO %s(url,uiOrder) VALUES (?, 0)", soundcloudTable)
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

func (u *SoundcloudUrlService) modifyTableWithUiOrder() {
	db, err := sql.Open("sqlite3", u.SqliteFile)
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error opening database in createSoundcloudUrl table")
		return
	}
	modifyUiOrderSQL := fmt.Sprintf("ALTER TABLE %s"+
		" ADD \"uiOrder\" integer", soundcloudTable)
	statement, err := db.Prepare(modifyUiOrderSQL) // Prepare SQL Statement
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error preparing modify soundcloudurl statement")
		return
	}
	_, err = statement.Exec()
	if err != nil {
		logging.GlobalLogger.Err(err).Msg("Error executing modify soundcloudurl table sql")
		return
	}
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

// GetAllSoundcloudUrls get all soundcloud urls in database
func (u *SoundcloudUrlService) GetAllSoundcloudUrls() ([]string, error) {
	db, err := sql.Open("sqlite3", u.SqliteFile)
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
		var id string
		var soundcloudUrl string
		var uiOrder int
		err := row.Scan(&id, &soundcloudUrl, &uiOrder)
		if err != nil {
			logging.GlobalLogger.Err(err).Msg("Error scanning sql row in get all soundcloud urls")
			return nil, err
		}
		soundcloudUrls = append(soundcloudUrls, soundcloudUrl)
	}
	return soundcloudUrls, nil
}
