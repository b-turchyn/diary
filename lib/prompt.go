package lib

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func PromptForInput(prompt string, args []string) (string, error) {
	var input string

	if len(args) == 0 {
		fmt.Print(prompt)
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			input = scanner.Text()
			break
		}
		if err := scanner.Err(); err != nil {
			return "", err
		}

		if len(input) == 0 {
			return "", errors.New("You didn't enter anything.")
		}
	} else {
		input = strings.Join(args, " ")
	}

	return input, nil
}

func GetTime(overrideDate string, overrideTime string) time.Time {
	var d, t time.Time
	var err error

	result := time.Now().Local()

	if overrideDate != "" {
		d, err = time.Parse("2006-01-02", overrideDate)

		if err != nil {
			log.Fatal(err)
		}

		result = time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, result.Location())
	}

	if overrideTime != "" {
		t, err = time.Parse("15:04", overrideTime)

		if err != nil {
			log.Fatal(err)
		}

		result = time.Date(result.Year(), result.Month(), result.Day(), t.Hour(), t.Minute(), t.Second(), 0, result.Location())
	}

	return result.UTC()
}
