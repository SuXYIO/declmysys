package decls

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
)

func (decls *Decls) Load(ddir string) error {
	declspath := filepath.Join(ddir, "decls")

	declsEnt, err := os.ReadDir(declspath)
	if err != nil {
		return err
	}

	for _, ent := range declsEnt {
		if ent.IsDir() {
			var decl Decl
			err := decl.Load(filepath.Join(declspath, ent.Name()))
			if err != nil {
				return err
			}
			*decls = append(*decls, decl)
		}
	}

	return nil
}

func (dot *Decl) Load(path string) error {
	// desc.toml
	descdata, err := os.ReadFile(filepath.Join(path, "desc.toml"))
	if err != nil {
		return err
	}
	err = dot.Desc.Load(descdata)
	if err != nil {
		return err
	}

	// data/
	dot.Data = filepath.Join(path, "data")

	return nil
}

func (desc *Desc) Load(data []byte) error {
	metadat, err := toml.Decode(string(data), desc)
	if err != nil {
		return err
	}

	// check default fields
	if !metadat.IsDefined("name") {
		return fmt.Errorf("must specify name for dot desc")
	}
	if !metadat.IsDefined("preset") {
		return fmt.Errorf("must specify preset for dot desc")
	}
	if !metadat.IsDefined("priority") {
		desc.Priority = consts.DefaultDeclsPriority
	}
	// check custom rundat fields for presets
	if desc.RunDat == nil {
		desc.RunDat = make(map[string]any)
	}

	// match preset
	preset, exists := Presets[desc.Preset]
	if !exists {
		return fmt.Errorf("preset not found for preset name: %s", desc.Preset)
	}
	if preset.DescIsValidFunc == nil {
		return fmt.Errorf("DescIsValidFunc not defined for preset %s", desc.Preset)
	}
	err = preset.DescIsValidFunc(*desc, metadat)
	if err != nil {
		return err
	}

	// subs
	if preset.DescIsValidFunc == nil {
		return fmt.Errorf("DescSubsFunc not defined for preset %s", desc.Preset)
	}
	err = preset.DescSubsFunc(desc)
	if err != nil {
		return err
	}

	return nil
}
