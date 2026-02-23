package utils

import (
	"fmt"
	"os"

	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/parse/subs"
	"github.com/suxyio/declmysys/internal/templates"
)

func GetGlobconf(gcpath string) (globconf.Globconf, error) {
	// parse path with paths&cmds subs first
	gcpath, err := subs.ApplyDefaultPCSubs(gcpath)
	if err != nil {
		return globconf.Globconf{}, err
	}

	// Get global config, create one if not exist
	var gc globconf.Globconf
	if _, err := os.Stat(gcpath); err == nil {
		// exists, no action
	} else if !os.IsNotExist(err) {
		// other error
		return gc, err
	} else {
		// not exist, create one
		fmt.Printf("Global config file not found at %s. Will load default config if not created.\n", gcpath)
		docreate := AskYN("create default there?")
		if !docreate {
			return globconf.DefaultGlobconf, nil
		}
		err = CreateFile(gcpath, templates.Globconf)
		if err != nil {
			return gc, err
		}
	}
	// read & parse
	dat, err := os.ReadFile(gcpath)
	if err != nil {
		return gc, err
	}
	gc, err = globconf.LoadGlobconf(dat)
	if err != nil {
		return gc, err
	}

	// and run subs
	err = globconf.SubsGlobconf(&gc)
	if err != nil {
		return gc, err
	}
	return gc, nil
}
