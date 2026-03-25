package subcmds

import (
	"fmt"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/ddir/decls"
	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/utils"
)

type RunOpts struct {
	Dry  bool `short:"d" long:"dry" description:"Dry run, only prints out the command to run"`
	Args struct {
		Priority uint `positional-arg-name:"priority" description:"Run the specific procedures for a certain priority" long-description:"Run the specific procedures for a certain priority"`
	} `positional-args:"yes"`
}

func Run(gc globconf.Globconf, mopts MainOpts, opts RunOpts) {
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

	// run
	// BUG: since go-flags doesn't support default for positionals yet,
	// can only assume that default value 0 is not user input
	var priority *uint
	if opts.Args.Priority == 0 {
		fmt.Printf("Running %s:\n", gc.DDir)
		priority = nil
	} else {
		fmt.Printf("Running %s (priority %d):\n", gc.DDir, opts.Args.Priority)
		priority = &opts.Args.Priority
	}

	declopts := decls.DeclsRunOpts{
		Indent:         1,
		FilterPriority: priority,
	}
	if opts.Dry {
		declopts.Dry = true
	}

	if err := declss.Run(declopts); err != nil {
		utils.Panic("error running decls", err, exitcode.ExecError)
	}
}
