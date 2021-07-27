package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

const SqlLiteDatabaseFileName = "sqlite-database.db"

func InitDatabase() {
	os.Remove(SqlLiteDatabaseFileName) // I delete the file to avoid duplicated records. SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db")
	//defer sqliteDatabase.Close()
	createTable(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createSoundcloudTableSQL := `CREATE TABLE soundcloudUrl (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"url" TEXT
	 );`

	log.Println("Create soundcloud table...")
	println("got here 1")
	statement, err := db.Prepare(createSoundcloudTableSQL) // Prepare SQL Statement
	println("got here 2")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err.Error())
		return
	} // Execute SQL Statements
	log.Println("soundcloud table created")
}

func addSoundcloudUrl(url string) {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	insertSoundcloudSQL := `INSERT INTO soundcloudUrl(url) VALUES (?)`
	statement, err := db.Prepare(insertSoundcloudSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(url)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GetAllSoundcloudUrls() {
	db, err := sql.Open("sqlite3", SqlLiteDatabaseFileName)
	defer db.Close()
	row, err := db.Query("SELECT * FROM soundcloudUrl ORDER BY url")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var url string
		row.Scan(&url)
		log.Println("SoundcloudUrl: ", url)
	}
}
