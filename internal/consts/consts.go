package consts

import "github.com/suxyio/declmysys/internal/parse/subs"

const (
	// program metadata
	Name    string = "declmysys"
	Desc    string = "DeclMySys (Declare My System), the simple, declarative, system config manager."
	Version string = "0.0a1"
	Author  string = "SuXYIO"

	// default paths
	defaultGlobconfPath string = "{CONF}/declmysys/config.toml"
	defaultDDDirPath    string = "{HOME}/Dotdecl"

	// default priorities
	DefaultPackagesPriority int = 200
	DefaultDotsPriority     int = 100
	DefaultActionsPriority  int = 50
)

func DefaultGlobconfPath() (string, error) {
	return subs.ApplyDefaultPC(defaultGlobconfPath)
}

func DefaultDDDirPath() (string, error) {
	return subs.ApplyDefaultPC(defaultDDDirPath)
}
