package packages

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

// LoadPkgsToml parses the packages.toml data
func (pkgs *Pkgs) Load(data []byte, sd subs.SubsDef) error {
	// toml decode
	metadat, err := toml.Decode(string(data), pkgs)
	if err != nil {
		return err
	}

	// check required fields
	// packages defined
	if !metadat.IsDefined("packages") {
		return fmt.Errorf("must specify packages (can be empty list though)")
	}
	// priority defined
	if !metadat.IsDefined("priority") {
		return fmt.Errorf("must specify priority")
	}
	for i, ps := range pkgs.Packages {
		// manager defined
		if len(ps.Manager.CustomCmd) == 0 && ps.Manager.Preset == "" {
			return fmt.Errorf("must specify manager for every packsspec, missing for packages[%d]", i)
		}
		// at least one pack
		if len(ps.Packs) < 1 {
			return fmt.Errorf("must specify at least one pack in all packsspec, missing for packages[%d]", i)
		}
	}

	// subs
	err = pkgs.Subs(sd)
	if err != nil {
		return err
	}

	return nil
}

func (pkgs *Pkgs) Subs(sd subs.SubsDef) error {
	for _, ps := range pkgs.Packages {
		// global subs for all
		for i := range ps.Packs {
			tmp, err := sd.ApplyG(ps.Packs[i])
			if err != nil {
				return err
			}
			ps.Packs[i] = tmp
		}

		// paths&cmds & global for self defined cmd
		for i := range ps.Manager.CustomCmd {
			subed, err := sd.ApplyPC(ps.Manager.CustomCmd[i])
			if err != nil {
				return err
			}
			ps.Manager.CustomCmd[i] = subed
		}
	}

	return nil
}

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
			RedirectStdout: true,
			RedirectStderr: true,
			AppendedArgs:   &ps.Packs,
		})
		if err != nil {
			return err
		}
	} else {
		// self defined manager command
		err := m.CustomCmd.Run(cmdtype.CmdRunOptions{
			RedirectStdout: true,
			RedirectStderr: true,
			AppendedArgs:   &ps.Packs,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Pkgs.Run installs all packages in a pkg
func (p Pkgs) Run() error {
	for _, ps := range p.Packages {
		err := ps.Run()
		if err != nil {
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
