package consts

import (
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
	"github.com/suxyio/declmysys/internal/utils"
)

const (
	// program metadata
	Name   string = "declmysys"
	Desc   string = "DeclMySys (Declare My System), the simple, declarative, system config manager."
	Author string = "SuXYIO"

	// default priorities
	// gonna use int here, since Priority is just an alias,
	// avoids cycle-import shit
	DefaultDeclsPriority int = 100

	// default misc
	DefaultDeclsDataDir string = "data"
	NotYourFault        string = "If you, as a user, see this message, it's probably not your fault. Please open a GitHub issue or contact the author for help."
)

var (
	// default paths, must be processed by ApplyDefaultPC
	DefaultGlobconfPath string = "{CONF}/declmysys/config.toml"
	DefaultDDirPath     string = "{HOME}/Decl"
)

func init() {
	// initializes the two vars, to adapt directories
	if ddirpath, err := subs.ApplyDefaultPC(DefaultDDirPath); err != nil {
		utils.Panic("unable to parse DefaultDDirPath via subs.ApplyDefaultPC", err, exitcode.SetupError)
	} else {
		DefaultDDirPath = ddirpath
	}
	if gcpath, err := subs.ApplyDefaultPC(DefaultGlobconfPath); err != nil {
		utils.Panic("unable to parse DefaultGlobconfPath via subs.ApplyDefaultPC", err, exitcode.SetupError)
	} else {
		DefaultGlobconfPath = gcpath
	}
}
