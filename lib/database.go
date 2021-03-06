package lib

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", viper.GetString("database"))

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	renameTable(db, "fuckups", "mistakes")
	createGenericTextTable(db, "mistakes")
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

func renameTable(db *sql.DB, oldName string, newName string) {
	db.Exec(fmt.Sprintf(`ALTER TABLE %s RENAME TO %s`, oldName, newName))
}

func InsertGenericText(db *sql.DB, name string, t time.Time, text string) error {
	_, err := db.Exec(fmt.Sprintf("INSERT INTO %s(date, text) VALUES(?, ?)", name), t, text)

	return err
}
