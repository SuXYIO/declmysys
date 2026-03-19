package decls

import (
	"sort"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

type Priority = int

// RunDecls runs the decls by their priority, from high to low
func RunDecls(decls []Decl, opts cmdtype.CmdRunOptions) error {
	//TODO: Add printing, prints the info during run
	if len(decls) == 0 {
		return nil
	}

	sort.Slice(decls, func(i, j int) bool {
		// ">" for high to low sort
		return decls[i].Priority > decls[j].Priority
	})

	for _, d := range decls {
		if err := d.Run(opts); err != nil {
			return err
		}
	}

	return nil
}
