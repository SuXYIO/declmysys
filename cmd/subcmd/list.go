package subcmd

import "github.com/suxyio/declmysys/internal/parse/globconf"

type ListOpts struct {
	Positional struct {
		Priority int `positional-arg-name:"priority" description:"Show the specific procedures for a certain priority" long-description:"Show the specific procedures for a certain priority, shows them more verbosely by default"`
	} `positional-args:"true"`
}

func List(gc globconf.Globconf, opts *ListOpts) error {
	// TODO: Impl
	return nil
}
