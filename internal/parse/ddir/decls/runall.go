package decls

import (
	"sort"
	"strings"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

type Priority = uint

type DeclsRunOpts struct {
	Cmdopts  cmdtype.CmdRunOptions
	DoPrint  bool
	Indent   uint
	Priority *Priority
}

// RunDecls runs the decls by their priority, from high to low.
// priority filters out other priorities, use nil for no filter
func (decls Decls) Run(opts DeclsRunOpts) error {
	if len(decls) == 0 {
		return nil
	}

	sort.Slice(decls, func(i, j int) bool {
		// ">" for high to low sort
		return decls[i].Priority > decls[j].Priority
	})

	indent := strings.Repeat("\t", int(opts.Indent))
	for _, d := range decls {
		// filter by priority
		if opts.Priority == nil || (opts.Priority != nil && d.Priority == *opts.Priority) {
			if opts.DoPrint {
				d.List(ToStringModeRun, indent)
			}
			if err := d.Run(opts.Cmdopts); err != nil {
				return err
			}
		}
	}

	return nil
}
