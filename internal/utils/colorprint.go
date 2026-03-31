package utils

import (
	"fmt"
	"os"
)

const (
	COLOR_BLACK   string = "\033[30m"
	COLOR_RED     string = "\033[31m"
	COLOR_GREEN   string = "\033[32m"
	COLOR_YELLOW  string = "\033[33m"
	COLOR_BLUE    string = "\033[34m"
	COLOR_MAGENTA string = "\033[35m"
	COLOR_CYAN    string = "\033[36m"
	COLOR_WHITE   string = "\033[37m"

	COLOR_RESET string = "\033[0m"
)

func ColorSprintf(format string, colorEsc string, a ...any) string {
	return fmt.Sprintf("%s%s%s", colorEsc, fmt.Sprintf(format, a...), COLOR_RESET)
}

func ErrorPrintf(format string, a ...any) (int, error) {
	return fmt.Fprint(os.Stderr, ColorSprintf(format, COLOR_RED, a...))
}

func WarnPrintf(format string, a ...any) (int, error) {
	return fmt.Fprint(os.Stderr, ColorSprintf(format, COLOR_YELLOW, a...))
}
