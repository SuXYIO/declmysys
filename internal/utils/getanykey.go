package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// GetAnyKey gets single byte from terminal, without needing to press enter. It adds newline automatically
func GetAnyKey() (byte, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return 0, err
	}
	defer fmt.Println() // this defer must go before the restore defer, since it has to be executed after restoring
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	b := make([]byte, 1)
	_, _ = os.Stdin.Read(b)

	return b[0], nil
}
