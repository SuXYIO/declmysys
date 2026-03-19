package subcmds

import (
	"fmt"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/ddir/decls"
	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/utils"
)

type ListOpts struct {
	Args struct {
		Priority uint `positional-arg-name:"priority" description:"Show the specific procedures for a certain priority" long-description:"Show the specific procedures for a certain priority"`
	} `positional-args:"yes"`
}

func List(gc globconf.Globconf, mopts MainOpts, opts ListOpts) {
	// subs.toml
	if err := subs.GetSubsToml(mopts.DDir); err != nil {
		utils.Panic("error getting subs.toml", err, exitcode.ConfigError)
	}

	// decls
	// can't use "decls" as var name since decls package took it damn it
	declss, err := decls.GetDecls(mopts.DDir)
	if err != nil {
		utils.Panic("error getting decls", err, exitcode.ConfigError)
	}

	// list
	// BUG: since go-flags doesn't support default for positionals yet,
	// can only assume that default value 0 is not user input
	var priority *uint
	if opts.Args.Priority == 0 {
		fmt.Printf("Listing %s:\n", gc.DDir)
		priority = nil
	} else {
		fmt.Printf("Listing %s (priority %d):\n", gc.DDir, opts.Args.Priority)
		priority = &opts.Args.Priority
	}
	if err := declss.List(decls.DeclsListOpts{
		Indent:   1,
		Priority: priority,
	}); err != nil {
		utils.Panic("error listing decls", err, exitcode.ExecError)
	}
}
