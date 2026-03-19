package decls

import (
	"fmt"
	"sort"
	"strings"

	"github.com/suxyio/declmysys/internal/consts"
)

type DeclsListOpts struct {
	Indent   int
	Priority *Priority
}

func (decls Decls) List(opts DeclsListOpts) error {
	if len(decls) == 0 {
		return nil
	}

	sort.Slice(decls, func(i, j int) bool {
		return decls[i].Priority > decls[j].Priority
	})

	// put same priority decls into one group
	var priGroups = [][]Decl{}
	currentgrp := []Decl{}
	prevPri := decls[0].Priority
	for _, d := range decls {
		if d.Priority != prevPri {
			priGroups = append(priGroups, currentgrp)
			prevPri = d.Priority
		}
		currentgrp = append(currentgrp, d)
	}
	priGroups = append(priGroups, currentgrp)

	indent := strings.Repeat(consts.Indent, opts.Indent)
	for _, grp := range priGroups {
		if opts.Priority == nil || (opts.Priority != nil && grp[0].Priority == *opts.Priority) {
			fmt.Printf("%s(%d)\n", indent, grp[0].Priority)
			for _, d := range grp {
				d.List(ToStringModeList, indent+consts.Indent)
			}
		}
	}

	return nil
}
