package globconf

import (
	"fmt"
	"os"

	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
	"github.com/suxyio/declmysys/internal/templates"
	"github.com/suxyio/declmysys/internal/utils"
)

func GetGlobconf(gcpath string) (Globconf, error) {
	// parse path with paths&cmds subs first
	if gcp, err := subs.ApplyDefaultPC(gcpath); err != nil {
		return Globconf{}, err
	} else {
		gcpath = gcp
	}

	// Get global config, create one if not exist
	var gc Globconf
	if _, err := os.Stat(gcpath); err == nil {
		// exists, no action
	} else if !os.IsNotExist(err) {
		// other error
		return gc, err
	} else {
		// not exist, create one
		fmt.Printf("Global config file not found at %s. Will load default config if not created.\n", gcpath)
		docreate := utils.AskYN("create default there?")
		if !docreate {
			return DefaultGlobconf, nil
		}
		err = utils.CreateFile(gcpath, templates.Globconf)
		if err != nil {
			return gc, err
		}
		fmt.Println("Global config file created")
	}
	// read & parse
	dat, err := os.ReadFile(gcpath)
	if err != nil {
		return gc, err
	}
	err = gc.Load(dat)
	if err != nil {
		return gc, err
	}

	return gc, nil
}
