package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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
