package subcmd

import (
	"github.com/suxyio/declmysys/cmd"
	"github.com/suxyio/declmysys/internal/parse/globconf"
)

type RunOpts struct {
	Dry        bool `short:"d" long:"dry" description:"Dry run, only print the procedures out" long-description:"Dry run, only print the procedures out. The difference from list subcommand is that dry run just prints the named structure with command, without redundant information e.g. undo, affected"`
	Verbose    bool `short:"v" long:"verbose" description:"Verbose, print verbose information, including procedure outputs and stats"`
	Positional struct {
		Procedure string `positional-arg-name:"procedure" description:"Specify the procedure to do" long-description:"Specify the procedures to do, e.g. dots.foobar. See procedure spec in docs"`
	} `positional-args:"true"`
}

func Run(gc globconf.Globconf, mopts *cmd.MainOpts, opts *RunOpts) {
	// TODO: Impl
}
