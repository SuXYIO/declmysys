package globconf

import (
	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

// Load parses the global config data
// sd option will be ignored, pass whatever you want,
// subs.toml is not loaded so passing it is meaningless, the argument is just for interface compatability
func (gc *Globconf) Load(data []byte, _ substoml.SubsDef) error {
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
	err = gc.Subs()
	if err != nil {
		return err
	}

	return nil
}

// Subs substitudes necessary stuff for globconf. No need for subs.SubsDef, since this only uses default subs
func (gc *Globconf) Subs() error {
	tmp, err := subs.ApplyDefaultPC(gc.DDDir)
	if err != nil {
		return err
	}
	gc.DDDir = tmp

	return nil
}
