package packagestoml

import (
	"fmt"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

// PacksSpec.Run installs the packages in PacksSpec
func (ps PacksSpec) Run() error {
	m := ps.Manager
	if m.Preset != "" {
		// preset manager name
		cmd, err := GetManagerPreset(m.Preset)
		if err != nil {
			return err
		}

		err = cmd.Run(cmdtype.CmdRunOptions{
			AppendedArgs: ps.Packs,
		})
		if err != nil {
			return err
		}
	} else {
		// self defined manager command
		if err := m.CustomCmd.Run(cmdtype.CmdRunOptions{
			AppendedArgs: ps.Packs,
		}); err != nil {
			return err
		}
	}
	return nil
}

// Pkgs.Run installs all packages in a pkg
func (p Pkgs) Run() error {
	for _, ps := range p.Packages {
		if err := ps.Run(); err != nil {
			return err
		}
	}
	return nil
}

// GetManagerPreset turns manager preset name to actual command
func GetManagerPreset(prename string) (cmdtype.Cmd, error) {
	switch prename {
	case "apt":
		return cmdtype.Cmd{"sudo", "apt", "install", "-y"}, nil
	case "flatpak-user-flathub":
		return cmdtype.Cmd{"flatpak", "install", "flathub", "--noninteractive", "-y", "--user"}, nil
	case "flatpak-system-flathub":
		return cmdtype.Cmd{"flatpak", "install", "flathub", "--noninteractive", "-y", "--system"}, nil
	default:
		return cmdtype.Cmd{}, fmt.Errorf("manager preset not found for %v", prename)
	}
}
