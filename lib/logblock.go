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

const LOGBLOCK_HEADER_TTY = COLOR_SECONDARY + "## %s" + COLOR_RESET + "\n\n%s\n"
const LOGBLOCK_HEADER_NOTTY = "## %s\n\n%s\n"
const LOGENTRY_TTY = COLOR_TERTIARY + "- %s:" + COLOR_RESET + " %s"
const LOGENTRY_NOTTY = "- %s: %s"

func (l LogBlock) ToString() string {
	var logStrings []string

	for _, v := range l.LogEntries {
		logStrings = append(logStrings, v.ToString())
	}

	if len(l.LogEntries) == 0 {
		return ""
	}

	return fmt.Sprintf(ttyCheck(LOGBLOCK_HEADER_TTY, LOGBLOCK_HEADER_NOTTY),
		l.Header,
		strings.Join(logStrings, "\n"),
	)
}

func (l LogEntry) ToString() string {
	return fmt.Sprintf(ttyCheck(LOGENTRY_TTY, LOGENTRY_NOTTY),
		l.Date.Local().Format("15:04"),
		l.Text,
	)
}

func GetLogBlock(db *sql.DB, name string, header string, time time.Time) (LogBlock, error) {
	entries, err := getForDate(db, name, time)

	return LogBlock{
		Header:     header,
		LogEntries: entries,
	}, err
}

func getForDate(db *sql.DB, name string, date time.Time) ([]LogEntry, error) {
	stmt, err := db.Prepare(fmt.Sprintf(`
SELECT id, date, text
  FROM %s
 WHERE datetime(date) >= datetime(?)
   AND datetime(date) < datetime(?)
 ORDER BY datetime(date)`, name),
	)
	if err != nil {
		return nil, err
	}

	// time.Truncate does not take into account time zone; we need to do this ourselves
	startDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Local().Location())
	endDate := date.AddDate(0, 0, 1)

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
