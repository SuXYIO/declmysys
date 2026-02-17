package subcmd

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type HelpOpts struct{}

// TODO: maybe add options showing help for specified subcmd?
func Help(p *flags.Parser) {
	// change active for showing main help
	// otherwise shows help for `help` subcmd
	a := p.Active
	p.Active = nil

	p.WriteHelp(os.Stdout)

	p.Active = a
}
