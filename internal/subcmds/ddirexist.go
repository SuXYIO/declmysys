package subcmds

import (
	"fmt"
	"os"

	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/urfave/cli/v3"
)

// getDDir gets decldir to use, decldir set in cmd overrides globconf, also checks if decldir exists
func getDDir(gc globconf.Globconf, cmd *cli.Command) (string, error) {
	// override ddir
	ddir := gc.DDir
	if cmd.IsSet("decldir") {
		ddir = cmd.String("decldir")
	}

	if ddirExist(ddir) {
		if dir, err := os.ReadDir(ddir); err != nil {
			return "", fmt.Errorf("failed to read ddir at %s: %v", ddir, err)
		} else {
			if len(dir) > 0 {
				return "", fmt.Errorf("ddir at %s already exists and is not empty", ddir)
			}
		}
	}

	return ddir, nil
}

func ddirExist(path string) bool {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}
