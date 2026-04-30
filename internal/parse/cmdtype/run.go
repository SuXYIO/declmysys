package cmdtype

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/suxyio/declmysys/internal/parse/metadata"
)

// CmdRunOptions defines options for running a command

type CmdRunOptions struct {
	RedirectStdin  io.Reader // default os.Stdin
	RedirectStdout io.Writer // default os.Stdout
	RedirectStderr io.Writer // default os.Stderr
	AppendedArgs   []string  // useful for package install append package spec, []string{} or nil for none
	WorkingDir     string    // change working directory, "" for use default, will do subs for it
	DryRun         io.Writer // won't run, only prints command to run to writer, nil for normal run
	DryPrintPrestr string    // prestr for dry run print, appended before every line
}

// Run runs a command
func (cmd Cmd) Run(opts CmdRunOptions) error {
	if len(cmd) == 0 {
		return nil
	}

	var c *exec.Cmd

	// do subs for cmd
	err := cmd.subs(&opts)
	if err != nil {
		return err
	}
	// also for workdir
	workdir, err := metadata.SubsApply(opts.WorkingDir)
	if err != nil {
		return err
	}

	if len(opts.AppendedArgs) > 0 {
		// append args
		c = exec.Command(cmd[0], append(cmd[1:], opts.AppendedArgs...)...)
	} else {
		c = exec.Command(cmd[0], cmd[1:]...)
	}

	if opts.DryRun != nil {
		// dry run
		fmt.Fprint(opts.DryRun, opts.DryPrintPrestr)
		_, err := fmt.Fprintf(opts.DryRun, "[workdir: %s] %v", workdir, c)
		fmt.Fprintln(opts.DryRun)
		return err
	}

	// change dir
	if workdir != "" {
		c.Dir = workdir
	}
	// redirect
	if opts.RedirectStdin != nil {
		c.Stdin = opts.RedirectStdin
	} else {
		c.Stdin = os.Stdin
	}
	if opts.RedirectStdout != nil {
		c.Stdout = opts.RedirectStdout
	} else {
		c.Stdout = os.Stdout
	}
	if opts.RedirectStderr != nil {
		c.Stderr = opts.RedirectStderr
	} else {
		c.Stderr = os.Stderr
	}

	// run
	if err := c.Run(); err != nil {
		return err
	}

	return nil
}

func (cmd *Cmd) subs(opts *CmdRunOptions) error {
	// paths&cmds for cmd
	for i := range *cmd {
		if elem, err := metadata.SubsApply((*cmd)[i]); err != nil {
			return err
		} else {
			(*cmd)[i] = elem
		}
	}

	// and also for appended args, changes opts directly
	for i := range opts.AppendedArgs {
		if elem, err := metadata.SubsApply(opts.AppendedArgs[i]); err != nil {
			return err
		} else {
			opts.AppendedArgs[i] = elem
		}
	}

	return nil
}
