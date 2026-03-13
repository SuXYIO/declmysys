package subcmds

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/suxyio/declmysys/internal/consts"
)

type VersionOpts struct{}

func Version(opts *VersionOpts) {
	// get version
	info, ok := debug.ReadBuildInfo()
	version := info.Main.Version
	if !ok || version == "" {
		version = "(unknown)"
	}

	fmt.Println(
		consts.Name,
		version,
		runtime.GOOS+"/"+runtime.GOARCH)
}
