package consts

import (
	"github.com/suxyio/declmysys/internal/parse/priority"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

const (
	// program metadata
	Name    string = "declmysys"
	Desc    string = "DeclMySys (Declare My System), the simple, declarative, system config manager."
	Version string = "0.0.1-alpha"
	Author  string = "SuXYIO"

	// default paths
	defaultGlobconfPath string = "{CONF}/declmysys/config.toml"
	defaultDDDirPath    string = "{HOME}/Dotdecl"

	// default priorities
	DefaultPackagesPriority priority.Priority = 200
	DefaultDotsPriority     priority.Priority = 100
	DefaultActionsPriority  priority.Priority = 50

	// default misc
	DefaultDotsDataDir string = "data"
)

func DefaultGlobconfPath() (string, error) {
	return subs.ApplyDefaultPC(defaultGlobconfPath)
}

func DefaultDDDirPath() (string, error) {
	return subs.ApplyDefaultPC(defaultDDDirPath)
}
