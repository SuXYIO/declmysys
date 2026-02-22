package subcmd

import "github.com/suxyio/declmysys/internal/parse/globconf"

type InitOpts struct {
	NoGit bool `long:"no-git" description:"Won't create the .git directory (via git init)"`
}

func Init(gc globconf.Globconf, opts *InitOpts) error {
	// TODO: Impl
	return nil
}
