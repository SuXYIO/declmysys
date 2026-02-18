package subcmd

import (
	"fmt"
	"runtime"

	"github.com/suxyio/declmysys/internal/consts"
)

type VersionOpts struct{}

func Version() {
	fmt.Println(consts.Name, consts.Version, runtime.GOOS+"/"+runtime.GOARCH)
}
