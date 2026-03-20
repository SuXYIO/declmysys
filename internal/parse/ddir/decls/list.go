package decls

import (
	"fmt"
	"io"

	"github.com/suxyio/declmysys/internal/consts"
)

func (decl Decl) List(w io.Writer, mode int8, prestr string) error {
	if preset, exists := presets[decl.Preset]; !exists {
		return fmt.Errorf("preset %q does not exist", decl.Preset)
	} else {
		if preset.ToString == nil {
			return fmt.Errorf("ToString function does not exist for preset %q. %s", decl.Preset, consts.NotYourFault)
		}
		fmt.Fprint(w, preset.ToString(decl, mode, prestr))
		return nil
	}
}
