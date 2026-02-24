package globconf

import (
	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

type Globconf struct {
	DDDir string `toml:"dotdecldir"`
}

var DefaultGlobconf Globconf = Globconf{
	DDDir: "~/Dotdecl",
}

// LoadGlobconf decodes toml []byte to Globconf
func LoadGlobconfToml(dat []byte) (Globconf, error) {
	var gc Globconf

	// toml decode
	_, err := toml.Decode(string(dat), &gc)
	if err != nil {
		return gc, err
	}

	// subs
	err = gc.SubsGlobconf()
	if err != nil {
		return gc, err
	}

	return gc, nil
}

// SubsGlobconf substitudes necessary stuff for globconf
func (gc *Globconf) SubsGlobconf() error {
	dddir, err := subs.ApplyDefaultPCSubs(gc.DDDir)
	if err != nil {
		return err
	}
	gc.DDDir = dddir

	return nil
}
