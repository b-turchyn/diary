package lib

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./diary.sqlite3")

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS fuckups(
    id INTEGER PRIMARY KEY,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    text TEXT
	)`)
	db.Exec("CREATE INDEX IF NOT EXISTS fuckups_date_ix ON fuckups(date)")

	return db
}
