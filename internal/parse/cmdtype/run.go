package cmdtype

import (
	"os"
	"os/exec"
)

// CmdRunOptions defines options for running a command
type CmdRunOptions struct {
	RedirectStdout bool     // if to redirect cmd.Stdout to os.Stdout
	RedirectStderr bool     // if to redirect cmd.Stderr to os.Stderr
	AppendedArgs   []string // useful for package install append package spec, []string{} for none
	WorkingDir     string   // change working directory, "" for use default
}

// Run runs a command
func (cmd Cmd) Run(opts CmdRunOptions) error {
	if len(cmd) == 0 {
		return nil
	}

	var c *exec.Cmd

	if len(opts.AppendedArgs) > 0 {
		// append args
		c = exec.Command(cmd[0], append(cmd[1:], opts.AppendedArgs...)...)
	} else {
		c = exec.Command(cmd[0], cmd[1:]...)
	}

	// change dir
	if opts.WorkingDir != "" {
		c.Dir = opts.WorkingDir
	}

	// redirect
	if opts.RedirectStdout {
		c.Stdout = os.Stdout
	}
	if opts.RedirectStderr {
		c.Stderr = os.Stderr
	}

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}
