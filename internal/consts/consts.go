package consts

import "github.com/suxyio/declmysys/internal/parse"

const (
	// program metadata
	Name    string = "declmysys"
	Desc    string = "DeclMySys (Declare My System), the simple, declarative, system config manager."
	Version string = "0.0a1"
	Author  string = "SuXYIO"

	// defaults
	defaultGlobconfPath string = "{CONF}/declmysys/config.toml"
	defaultDDDirPath    string = "{HOME}/Dotdecl"
)

func DefaultGlobconfPath() (string, error) {
	return parse.ApplyDefaultFCSubs(defaultGlobconfPath)
}

func DefaultDDDirPath() (string, error) {
	return parse.ApplyDefaultFCSubs(defaultDDDirPath)
}
