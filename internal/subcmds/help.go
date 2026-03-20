package subcmds

import (
	"bytes"

	"github.com/jessevdk/go-flags"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/utils"
)

type HelpOpts struct {
	Args struct {
		Subcommand string `positional-arg-name:"subcommand" description:"Specify the subcommand to show help to" long-description:"Specify the subcommand to show help to, if not specified or cannot find subcommand with this name, shows help for the main program"`
	} `positional-args:"yes"`
}

func Help(parser flags.Parser, opts HelpOpts) {
	// change active for showing main help
	// otherwise shows help for `help` subcmd
	// HACK: this is really awkward code, changing the values of parser and all
	// I was expecting something like someSubcmd.WriteHelp() instead of this
	parser.Active = parser.Find(opts.Args.Subcommand)

	var buf bytes.Buffer
	parser.WriteHelp(&buf)

	if err := utils.AutoPager(buf.Bytes()); err != nil {
		utils.Panic("error autopaging", err, exitcode.ExecError)
	}
}
