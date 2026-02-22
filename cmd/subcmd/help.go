package subcmd

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type HelpOpts struct {
	Positional struct {
		Subcommand string `positional-arg-name:"subcommand" description:"Specify the subcommand to show help to" long-description:"Specify the subcommand to show help to, if not specified or cannot find subcommand with this name, shows help for the main program"`
	} `positional-args:"true"`
}

func Help(opts *HelpOpts, p *flags.Parser) {
	// change active for showing main help
	// otherwise shows help for `help` subcmd
	a := p.Active
	p.Active = p.Find(opts.Positional.Subcommand)

	p.WriteHelp(os.Stdout)

	p.Active = a
}
