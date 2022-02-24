package lib

import (
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

	return fmt.Sprintf("%s\n%s\n\n%s\n",
		l.Header,
		strings.Repeat("=", len(l.Header)),
		strings.Join(logStrings, "\n"),
	)
}

func (l LogEntry) ToString() string {
	return fmt.Sprintf("- %s: %s",
		l.Date.Local().Format("15:04"),
		l.Text,
	)
}
