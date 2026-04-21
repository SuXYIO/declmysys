package globconf

import (
	"github.com/suxyio/declmysys/internal/consts"
)

type Globconf struct {
	DDir string `toml:"decldir"`
}

var DefaultGlobconf Globconf = Globconf{
	DDir: consts.DefaultGlobconfPath,
}
