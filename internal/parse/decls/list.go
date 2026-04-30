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

		fmt.Fprintf(w, "%s%s[%s]", prestr, decl.Preset, decl.Name)
		if mode == ToStringModeList {
			fmt.Fprintf(w, " [workdir: %s]", decl.Pwd)
			fmt.Fprint(w, preset.ToString(decl))
		}
		fmt.Fprintln(w)
		return nil
	}
}
