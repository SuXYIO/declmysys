package decls

import "fmt"

func (decls Decls) Run() error {
	for _, decl := range decls {
		switch decl.Desc.Preset {
		// TODO: Impl
		case "stow":
		case "gitclone":
		case "cmds":
		default:
			return fmt.Errorf("unrecognized preset name: %s", decl.Desc.Preset)
		}
	}

	return nil
}
