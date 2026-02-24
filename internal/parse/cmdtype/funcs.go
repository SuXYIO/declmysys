package cmdtype

import (
	"os"
	"os/exec"
)

type CmdRunOptions struct {
	RedirectStdout bool      // if to redirect cmd.Stdout to os.Stdout
	RedirectStderr bool      // if to redirect cmd.Stderr to os.Stderr
	AppendedArgs   *[]string // useful for package install append package spec
}

// Run runs a command
func (cmd Cmd) Run(opts CmdRunOptions) error {
	var c *exec.Cmd
	if opts.AppendedArgs != nil {
		// append args
		c = exec.Command(cmd[0], append(cmd[1:], *opts.AppendedArgs...)...)
	} else {
		c = exec.Command(cmd[0], cmd[1:]...)
	}

	// redirect
	if opts.RedirectStdout {
		c.Stdout = os.Stdout
	}
	if opts.RedirectStderr {
		c.Stderr = os.Stderr
	}

	err := c.Run()
	if err != nil {
		return err
	}

	return nil
}
