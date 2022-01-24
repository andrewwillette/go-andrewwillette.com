package config

const (
	SqliteFile = "sqlite-database.db"
	Port       = 9099
)

func GetCorsWhiteList() []string {
	return []string{"http://localhost:3000", "http://andrewwillette.com"}
}
