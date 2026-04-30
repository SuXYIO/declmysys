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
		if _, err := os.ReadDir(ddir); err != nil {
			return "", fmt.Errorf("failed to read ddir at %s: %v", ddir, err)
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
