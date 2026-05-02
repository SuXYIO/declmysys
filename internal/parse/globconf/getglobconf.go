package globconf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/metadata"
	"github.com/suxyio/declmysys/internal/templates"
	"github.com/suxyio/declmysys/internal/utils"
)

// GetGlobconf gets global config.
func GetGlobconf() (Globconf, error) {
	// choose, envvar first, default second
	gcpath := consts.DefaultGlobconfPath
	if envGcpath := os.Getenv(consts.GlobconfPathEnvvar); envGcpath != "" {
		gcpath = envGcpath
	}

	// parse path with paths&cmds subs first
	gcp, err := metadata.ApplyDefaultsSubs(gcpath)
	if err != nil {
		return Globconf{}, err
	}
	gcfp := filepath.Join(gcp, "config.toml")

	// Get global config, create one if not exist
	var gc Globconf
	if _, err := os.Stat(gcfp); err == nil {
		// exists, no action
	} else if !os.IsNotExist(err) {
		// other error
		return gc, err
	} else {
		// not exist, create one
		fmt.Printf("Global config file not found at %s. Will load default config if not created.\n", gcfp)
		docreate := utils.AskYN("create default there?")
		if !docreate {
			return DefaultGlobconf, nil
		}
		err = utils.CreateFile(gcfp, templates.Globconf)
		if err != nil {
			return gc, err
		}
		fmt.Println("Global config file created")
	}
	// read & parse
	dat, err := os.ReadFile(gcfp)
	if err != nil {
		return gc, err
	}
	err = gc.Load(dat)
	if err != nil {
		return gc, err
	}

	return gc, nil
}
