package lib

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./diary.sqlite3")

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

func GetLogBlockForToday(db *sql.DB, name string, header string) (LogBlock, error) {
	entries, err := getForToday(db, name)

	return LogBlock{
		Header:     header,
		LogEntries: entries,
	}, err
}

func GetLogBlock(db *sql.DB, name string, header string, time time.Time) (LogBlock, error) {
	entries, err := getForDate(db, name, time)

	return LogBlock{
		Header:     header,
		LogEntries: entries,
	}, err
}

func getForToday(db *sql.DB, name string) ([]LogEntry, error) {
	return getForDate(db, name, time.Now().Local())
}

func getForDate(db *sql.DB, name string, date time.Time) ([]LogEntry, error) {
	stmt, err := db.Prepare(
		fmt.Sprintf("SELECT id, date, text FROM %s WHERE date >= ? AND date < ? ORDER BY date", name),
	)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(date.Truncate(time.Hour*24), date.AddDate(0, 0, 1).Truncate(time.Hour*24))
	if err != nil {
		return nil, err
	} else {
		var result []LogEntry
		for rows.Next() {
			var row LogEntry

			err = rows.Scan(&row.Id, &row.Date, &row.Text)
			if err != nil {
				return nil, err
			}

			result = append(result, row)
		}

		return result, nil
	}
}
