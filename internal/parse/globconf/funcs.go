package globconf

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

// LoadGlobconf decodes toml []byte to Globconf
// sd option will be ignored, pass whatever you want,
// subs.toml is not loaded so passing it is meaningless, the argument is just for interface compatability
func (gc *Globconf) Load(data []byte, _ subs.SubsDef) error {
	// toml decode
	metadat, err := toml.Decode(string(data), gc)
	if err != nil {
		return err
	}

	// check exist
	if !metadat.IsDefined("dotdecldir") {
		return fmt.Errorf("must specify default dotdecldir")
	}

	// subs
	err = gc.Subs()
	if err != nil {
		return err
	}

	return nil
}

// SubsGlobconf substitudes necessary stuff for globconf. No need for subs.SubsDef, since this only uses default subs
func (gc *Globconf) Subs() error {
	tmp, err := subs.ApplyDefaultPC(gc.DDDir)
	if err != nil {
		return err
	}
	gc.DDDir = tmp

	return nil
}
