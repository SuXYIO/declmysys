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
	defaultDDirPath     string = "{HOME}/Decl"

	// default priorities
	DefaultPackagesPriority priority.Priority = 200
	DefaultDeclsPriority    priority.Priority = 100

	// default misc
	DefaultDeclsDataDir string = "data"
)

func DefaultGlobconfPath() (string, error) {
	return subs.ApplyDefaultPC(defaultGlobconfPath)
}

func DefaultDDirPath() (string, error) {
	return subs.ApplyDefaultPC(defaultDDirPath)
}
