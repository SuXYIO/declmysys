package consts

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/metadata"
	"github.com/suxyio/declmysys/internal/utils"
)

const (
	// program metadata
	Name   string = "declmysys"
	Desc   string = "DeclMySys (Declare My System), the simple, declarative, system config manager."
	Author string = "SuXYIO"

	// default priorities
	// gonna use uint here, since Priority is just an alias,
	// avoids cycle-import shit
	DefaultDeclPriority uint   = 100
	NotYourFault        string = "If you, as a user, see this message, it's probably not your fault. Please open a GitHub issue or contact the author for help."
	Indent              string = "    "
)

var (
	Version   string = "" // initialized in init()
	BuildInfo string = "" // initialized in init()

	// default paths, must be processed by ApplyDefaultPC
	DefaultGlobconfPath string = "{CONF}/declmysys" // note that this is the directory path
	DefaultDDirPath     string = "{HOME}/Decl"
	GlobconfPathEnvvar  string = "DECLMYSYS_CONFIG"
)

func init() {
	if ver, err := getVer(); err != nil {
		Version = "(unknown)"
		// probably should log something, but havn't impled log yet
	} else {
		Version = ver
	}
	BuildInfo = fmt.Sprintf(
		"%s %s",
		Version,
		runtime.GOOS+"/"+runtime.GOARCH)

	// initializes the two vars, to adapt directories
	if ddirpath, err := metadata.ApplyDefaultsSubs(DefaultDDirPath); err != nil {
		utils.Panic("unable to parse DefaultDDirPath via metadata.ApplyDefaultSubs", err, exitcode.SetupError)
	} else {
		DefaultDDirPath = ddirpath
	}
	if gcpath, err := metadata.ApplyDefaultsSubs(DefaultGlobconfPath); err != nil {
		utils.Panic("unable to parse DefaultGlobconfPath via metadata.ApplyDefaultSubs", err, exitcode.SetupError)
	} else {
		DefaultGlobconfPath = gcpath
	}
}

func getVer() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("error reading build info")
	}
	version := info.Main.Version
	return version, nil
}
