package lib

import (
	"os"

	isatty "github.com/mattn/go-isatty"
	"github.com/spf13/viper"
)

const COLOR_RESET = "\033[0m"
const COLOR_PRIMARY = "\033[1;31m"
const COLOR_SECONDARY = "\033[1;32m"
const COLOR_TERTIARY = "\033[1;33m"

func ttyCheck(yes string, no string) string {
	if viper.GetBool("color") && isatty.IsTerminal(os.Stdout.Fd()) {
		return yes
	} else {
		return no
	}
}

func PrimaryText(input string) string {
	return ttyCheck(COLOR_PRIMARY+input+COLOR_RESET, input)
}
