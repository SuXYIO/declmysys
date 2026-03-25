package decls

import (
	"fmt"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

func (decl Decl) Run(opts cmdtype.CmdRunOptions) error {
	// note that opts.AppendedArgs will be ignored
	// subs
	preset, exists := presets[decl.Preset]
	if !exists {
		return fmt.Errorf("preset not found for preset name: %q", decl.Preset)
	}

	// run
	if preset.RunFunc == nil {
		return fmt.Errorf("RunFunc not defined for preset %q", decl.Preset)
	}
	if err := preset.RunFunc(decl, opts); err != nil {
		return err
	}

	return nil
}
