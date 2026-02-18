package subcmd

import "github.com/suxyio/declmysys/internal/parse"

type ListOpts struct {
	Positional struct {
		Priority int `positional-name:"priority" required:"false" description:"Show the specific procedures for a certain priority" long-description:"Show the specific procedures for a certain priority, shows them more verbosely by default"`
	}
}

// TODO: Impl
func List(gc parse.Globconf, opts *ListOpts) error {
	return nil
}
