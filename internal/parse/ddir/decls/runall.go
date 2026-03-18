package decls

import (
	"context"
	"sort"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"golang.org/x/sync/errgroup"
)

type Priority = int

// RunDecls runs the decls by their priority, from high to low
func RunDecls(ctx context.Context, decls []Decl, opts cmdtype.CmdRunOptions) error {
	//TODO: Add printing, prints the info during run
	if len(decls) == 0 {
		return nil
	}

	sort.Slice(decls, func(i, j int) bool {
		// ">" for high to low sort
		return decls[i].Priority > decls[j].Priority
	})

	var group []Decl
	prevPri := decls[0].Priority
	for _, d := range decls {
		if d.Priority != prevPri {
			// unequal priority, run stored group
			if err := runGroup(ctx, group, opts); err != nil {
				return err
			}
			group = nil
			prevPri = d.Priority
		}
		group = append(group, d)
	}

	// handle last group
	return runGroup(ctx, group, opts)
}

// runGroup runs decls concurrently
func runGroup(ctx context.Context, group []Decl, opts cmdtype.CmdRunOptions) error {
	if len(group) == 0 {
		return nil
	} else if len(group) == 1 {
		return group[0].Run(opts)
	} else {
		// multiple, run concurrently
		grp, _ := errgroup.WithContext(ctx)
		for _, d := range group {
			grp.Go(func() error {
				return d.Run(opts)
			})
		}
		return grp.Wait()
	}
}
