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
func LoadGlobconf(dat []byte) (Globconf, error) {
	var gc Globconf

	// toml decode
	_, err := toml.Decode(string(dat), &gc)
	if err != nil {
		return gc, err
	}

	// subs
	res, err := subs.ApplyDefaultPCSubs(gc.DDDir)
	if err != nil {
		return gc, err
	}
	gc.DDDir = res

	return gc, nil
}

// SubsGlobconf substitudes necessary stuff for globconf
func SubsGlobconf(gc *Globconf) error {
	dddir, err := subs.ApplyDefaultPCSubs(gc.DDDir)
	if err != nil {
		return err
	}
	gc.DDDir = dddir

	return nil
}
