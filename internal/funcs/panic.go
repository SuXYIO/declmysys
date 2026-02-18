package funcs

import (
	"fmt"
	"os"
)

/*
Probs are gonna use this in main. Shortcut, and not losing control.
Writing these damn things over and over, it's more comfortable even if only saving a single line.
Before:

	if err != nil {
		fmt.Fprintf(os.Stderr, "some error: %v\n", err)
		os.Exit(exitcode.SomeError)
	}

After:

	if err != nil {
		funcs.Panic("some error", err, exitcode.SomeError)
	}
*/
func Panic(msg string, err error, exitcode int) {
	fmt.Fprintf(os.Stderr, msg+": %v\n", err)
	os.Exit(exitcode)
}
