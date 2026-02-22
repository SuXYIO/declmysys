package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/suxyio/declmysys/internal/parse"
	"github.com/suxyio/declmysys/internal/templates"
)

func GetGlobconf(gcpath string) (parse.Globconf, error) {
	// parse path with paths&cmds subs first
	gcpath, err := parse.ApplyDefaultFCSubs(gcpath)
	if err != nil {
		return parse.Globconf{}, err
	}

	// Get global config, create one if not exist
	var gc parse.Globconf
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
			return parse.DefaultGlobconf, nil
		}

		// create dir
		dir := filepath.Dir(gcpath)
		if dir != "." && dir != "" {
			fmt.Printf("Missing top directories found, creating all: %s\n", dir)
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return gc, err
			}
		}

		// copy file
		err = os.WriteFile(gcpath, templates.Globconf, 0644)
		if err != nil {
			return gc, err
		}
	}
	// read it
	dat, err := os.ReadFile(gcpath)
	if err != nil {
		return gc, err
	}
	gc, err = parse.LoadGlobconf(dat)
	if err != nil {
		return gc, err
	}

	return gc, nil
}
