package subcmds

import (
	"bytes"
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
	// ddir doesn't exist
	if !DDirExist(mopts.DDir) {
		utils.Panic(fmt.Sprintf("ddir %s does not exist, you can create via \"init\" subcommand", mopts.DDir), nil, exitcode.FileError)
	}

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
	var w bytes.Buffer
	var priority *uint
	if opts.Args.Priority == 0 {
		fmt.Fprintf(&w, "Listing %s:\n", gc.DDir)
		priority = nil
	} else {
		fmt.Fprintf(&w, "Listing %s (priority %d):\n", gc.DDir, opts.Args.Priority)
		priority = &opts.Args.Priority
	}
	if err := declss.List(&w, decls.DeclsListOpts{
		Indent:         1,
		FilterPriority: priority,
	}); err != nil {
		utils.Panic("error listing decls", err, exitcode.ExecError)
	}

	if err := utils.AutoPager(w.Bytes()); err != nil {
		utils.Panic("error autopaging", err, exitcode.ExecError)
	}
}
