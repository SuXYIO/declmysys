package subcmds

import (
	"fmt"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/decls"
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
	declss := getDeclsData(mopts.DDir)

	// run
	// BUG: since go-flags doesn't support default for positionals yet,
	// can only assume that default value 0 is not user input
	var priority *uint
	if opts.Args.Priority == 0 {
		if opts.Dry {
			fmt.Printf("Running %s (dry):\n", gc.DDir)
		} else {
			fmt.Printf("Running %s:\n", gc.DDir)
			priority = nil
		}
	} else {
		if opts.Dry {
			fmt.Printf("Running %s (dry, priority %d):\n", gc.DDir, opts.Args.Priority)
		} else {
			fmt.Printf("Running %s (priority %d):\n", gc.DDir, opts.Args.Priority)
			priority = &opts.Args.Priority
		}
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
