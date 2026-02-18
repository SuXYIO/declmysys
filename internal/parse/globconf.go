package parse

import (
	"github.com/BurntSushi/toml"
)

type Globconf struct {
	DDDir    string `toml:"dotdecldir"`
	GlyphSet string `toml:"glyphset"`
}

func ParseGlobconf(dat []byte) (Globconf, error) {
	var gc Globconf

	_, err := toml.Decode(string(dat), &gc)
	if err != nil {
		return gc, err
	}

	return gc, nil
}
