package packagestoml

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/ddir/substoml"
)

// Load parses the packages.toml data
func (pkgs *Pkgs) Load(data []byte) error {
	// toml decode
	metadat, err := toml.Decode(string(data), pkgs)
	if err != nil {
		return err
	}

	// replace default fields
	if !metadat.IsDefined("packages") {
		pkgs.Packages = []PacksSpec{}
	}
	if !metadat.IsDefined("priority") {
		pkgs.Priority = consts.DefaultPackagesPriority
	}
	for i, ps := range pkgs.Packages {
		// manager defined
		if len(ps.Manager.CustomCmd) == 0 && ps.Manager.Preset == "" {
			return fmt.Errorf("must specify manager for every packsspec, missing for packages[%d]", i)
		}
		// at least one pack
		if len(ps.Packs) < 1 {
			return fmt.Errorf("must specify at least one pack in all packsspec, missing for packages[%d]", i)
		}
	}

	// subs
	err = pkgs.subs()
	if err != nil {
		return err
	}

	return nil
}

func (pkgs *Pkgs) subs() error {
	for _, ps := range pkgs.Packages {
		// paths&cmds & global for self defined cmd
		for i := range ps.Manager.CustomCmd {
			if cmdelem, err := substoml.ApplyPC(ps.Manager.CustomCmd[i]); err != nil {
				return err
			} else {
				ps.Manager.CustomCmd[i] = cmdelem
			}
		}
	}

	return nil
}
