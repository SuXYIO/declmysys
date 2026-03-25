package decls

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

type Priority = uint

type DeclsRunOpts struct {
	Indent         uint
	WorkingDir     string
	FilterPriority *Priority
	Dry            bool
	RedirectStdout io.Writer
	RedirectStderr io.Writer
	noPrint        bool // convenient for testing, only works for Dry == false
}

// RunDecls runs the decls by their priority, from high to low.
// priority filters out other priorities, use nil for no filter
func (decls Decls) Run(opts DeclsRunOpts) error {
	if len(decls) == 0 {
		return nil
	}

	// sort in order
	sort.Slice(decls, func(i, j int) bool {
		// ">" for high to low sort
		return decls[i].Priority > decls[j].Priority
	})

	indent := strings.Repeat(consts.Indent, int(opts.Indent))
	for _, d := range decls {
		// filter by priority
		if opts.FilterPriority == nil || (opts.FilterPriority != nil && d.Priority == *opts.FilterPriority) {
			if opts.Dry {
				// dry run
				if err := d.List(os.Stdout, ToStringModeRun, indent); err != nil {
					return err
				}
				fmt.Println()
				fmt.Print(indent + consts.Indent)
				if err := d.Run(cmdtype.CmdRunOptions{
					WorkingDir: opts.WorkingDir,
					DryRun:     os.Stdout,
				}); err != nil {
					return err
				}
			} else {
				if !opts.noPrint {
					if err := d.List(os.Stdout, ToStringModeRun, indent); err != nil {
						return err
					}
					fmt.Println()
				}
				if err := d.Run(cmdtype.CmdRunOptions{
					RedirectStdout: opts.RedirectStdout,
					RedirectStderr: opts.RedirectStderr,
					WorkingDir:     opts.WorkingDir,
					DryRun:         nil,
				}); err != nil {
					return err
				}
				if !opts.noPrint {
					fmt.Printf("%s%sDone!\n", indent, consts.Indent)
				}
			}
		}
	}

	return nil
}
