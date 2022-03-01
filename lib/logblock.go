package lib

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type LogEntry struct {
	Id   int
	Date time.Time
	Text string
}

type LogBlock struct {
	Header     string
	LogEntries []LogEntry
}

type logBlock interface {
	ToString() string
}

type logEntry interface {
	ToString() string
}

func (l LogBlock) ToString() string {
	var logStrings []string

	for _, v := range l.LogEntries {
		logStrings = append(logStrings, v.ToString())
	}

	return fmt.Sprintf("## %s\n\n%s\n",
		l.Header,
		strings.Join(logStrings, "\n"),
	)
}

func (l LogEntry) ToString() string {
	return fmt.Sprintf("- %s: %s",
		l.Date.Local().Format("15:04"),
		l.Text,
	)
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
		fmt.Sprintf("SELECT id, date, text FROM %s WHERE datetime(date) >= ? AND datetime(date) < ? ORDER BY datetime(date)", name),
	)
	if err != nil {
		return nil, err
	}

	date = date.UTC()

	startDate := date.Truncate(24 * time.Hour)
	endDate := date.AddDate(0, 0, 1).Truncate(time.Hour * 24)

	rows, err := stmt.Query(startDate, endDate)
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
