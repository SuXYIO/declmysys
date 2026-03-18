package cmdtype

import (
	"os"
	"os/exec"

	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
)

// CmdRunOptions defines options for running a command
type CmdRunOptions struct {
	RedirectStdout *os.File
	RedirectStderr *os.File
	AppendedArgs   []string // useful for package install append package spec, []string{} or nil for none
	WorkingDir     string   // change working directory, "" for use default
	DoPCSubs       bool     // whether to do default PC subs
}

// Run runs a command
func (cmd Cmd) Run(opts CmdRunOptions) error {
	if len(cmd) == 0 {
		return nil
	}

	var c *exec.Cmd

	if opts.DoPCSubs {
		// do subs
		err := cmd.subs(&opts)
		if err != nil {
			return err
		}
	}

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
	if opts.RedirectStdout != nil {
		c.Stdout = opts.RedirectStdout
	}
	if opts.RedirectStderr != nil {
		c.Stderr = opts.RedirectStderr
	}

	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

func (cmd *Cmd) subs(opts *CmdRunOptions) error {
	// paths&cmds for cmd
	for i := range *cmd {
		if elem, err := subs.ApplyPC((*cmd)[i]); err != nil {
			return err
		} else {
			(*cmd)[i] = elem
		}
	}

	// and also for appended args, changes opts directly
	for i := range opts.AppendedArgs {
		if elem, err := subs.ApplyPC(opts.AppendedArgs[i]); err != nil {
			return err
		} else {
			opts.AppendedArgs[i] = elem
		}
	}

	return nil
}
