package parse

import (
	"github.com/BurntSushi/toml"
)

type Globconf struct {
	DDDir    string `toml:"dotdecldir"`
	GlyphSet string `toml:"glyphset"`
}

var DefaultGlobconf Globconf = Globconf{
	DDDir:    "~/Dotdecl",
	GlyphSet: "ascii",
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
	res, err := ApplyDefaultPCSubs(gc.DDDir)
	if err != nil {
		return gc, err
	}
	gc.DDDir = res

	return gc, nil
}

// SubsGlobconf substitudes necessary stuff for globconf
func SubsGlobconf(gc Globconf) (Globconf, error) {
	dddir, err := ApplyDefaultPCSubs(gc.DDDir)
	if err != nil {
		return gc, err
	}

	gc.DDDir = dddir
	return gc, nil
}
