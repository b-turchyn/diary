package lib

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", viper.GetString("database"))

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	createGenericTextTable(db, "fuckups")
	createGenericTextTable(db, "log")
	createGenericTextTable(db, "learn")

	return db
}

func createGenericTextTable(db *sql.DB, name string) {
	db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id INTEGER PRIMARY KEY,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    text TEXT
	)`, name))
	db.Exec(fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_date_ix ON %s(date)", name, name))
}

func InsertGenericText(db *sql.DB, name string, text string) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s(text) VALUES(?)", name), text)

	return err
}
