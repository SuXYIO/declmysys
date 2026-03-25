package utils

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/term"
)

// AutoPager automatically determines whether to use a pager for long output, and outputs it
func AutoPager(data []byte) error {
	stdoutfd := int(os.Stdout.Fd())
	if !term.IsTerminal(stdoutfd) {
		return direct(data)
	}

	_, h, err := term.GetSize(stdoutfd)
	if err != nil {
		return direct(data)
	}

	linecnt := bytes.Count(data, []byte{'\n'})

	if linecnt <= h {
		// direct output
		return direct(data)
	} else {
		// pager
		return pager(data)
	}
}

func pager(data []byte) error {
	if os.Getenv("NO_PAGER") != "" {
		return direct(data)
	}
	pagerCmd := os.Getenv("PAGER")
	if pagerCmd == "" {
		// no pager set, try common
		for _, p := range []string{"less", "more"} {
			if path, err := exec.LookPath(p); err == nil {
				pagerCmd = path
				break
			}
		}
	}
	if pagerCmd == "" {
		// no common found
		return direct(data)
	}

	args := strings.Fields(pagerCmd)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = bytes.NewReader(data)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func direct(data []byte) error {
	_, err := fmt.Print(string(data))
	return err
}
