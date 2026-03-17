package packages

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
)

// Load parses the packages.toml data
func (pkgs *Pkgs) Load(data []byte) error {
	// toml decode
	md, err := toml.Decode(string(data), pkgs)
	if err != nil {
		return err
	}

	if err := pkgs.isValid(); err != nil {
		return err
	}

	if err := pkgs.setDefaults(md); err != nil {
		return err
	}

	return nil
}

// isValid checks if loaded pkgs is valid
func (pkgs Pkgs) isValid() error {
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

	return nil
}

// setDefaults sets the default values for loaded pkgs
func (pkgs *Pkgs) setDefaults(md toml.MetaData) error {
	// replace default fields
	if !md.IsDefined("packages") {
		pkgs.Packages = []PacksSpec{}
	}
	if !md.IsDefined("priority") {
		pkgs.Priority = consts.DefaultPackagesPriority
	}

	return nil
}
