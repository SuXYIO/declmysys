package globconf

import (
	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

// Load parses the global config data
func (gc *Globconf) Load(data []byte) error {
	// toml decode
	metadat, err := toml.Decode(string(data), gc)
	if err != nil {
		return err
	}

	// replace default fields
	if !metadat.IsDefined("dotdecldir") {
		defaultDDDir, err := consts.DefaultDDDirPath()
		if err != nil {
			return err
		}
		gc.DDDir = defaultDDDir
	}

	// subs
	err = gc.subs()
	if err != nil {
		return err
	}

	return nil
}

// Subs substitudes necessary stuff for globconf
func (gc *Globconf) subs() error {
	tmp, err := subs.ApplyDefaultPC(gc.DDDir)
	if err != nil {
		return err
	}
	gc.DDDir = tmp

	return nil
}
