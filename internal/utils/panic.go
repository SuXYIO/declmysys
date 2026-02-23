package utils

import (
	"github.com/mitchellh/colorstring"
	"os"
)

func Panic(msg string, err error, exitcode int) {
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
			Panic("some error", err, exitcode.SomeError)
		}
	*/
	if err != nil {
		colorstring.Fprintf(os.Stderr, "[red]"+msg+": %v\n", err)
	} else {
		colorstring.Fprint(os.Stderr, "[red]"+msg+"\n")
	}
	os.Exit(exitcode)
}
