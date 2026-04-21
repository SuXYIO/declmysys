package subcmds

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/utils"
)

type VersionOpts struct{}

func Version(opts VersionOpts) {
	// get version
	version, err := getVer()
	if err != nil {
		utils.Panic("error getting version", err, exitcode.Unknown)
	}

	fmt.Println(
		consts.Name,
		version,
		runtime.GOOS+"/"+runtime.GOARCH)
}

func getVer() (string, error) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "", fmt.Errorf("error reading build info")
	}
	version := info.Main.Version
	return version, nil
}
