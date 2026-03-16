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
			if err := decl.Load(filepath.Join(declspath, ent.Name())); err != nil {
				return err
			}
			*decls = append(*decls, decl)
		}
	}

	return nil
}

func (decl *Decl) Load(path string) error {
	// desc.toml
	descdata, err := os.ReadFile(filepath.Join(path, "desc.toml"))
	if err != nil {
		return err
	}
	err = decl.loadDesc(descdata)
	if err != nil {
		return err
	}

	return nil
}

func (decl *Decl) loadDesc(data []byte) error {
	metadat, err := toml.Decode(string(data), decl)
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
		decl.Priority = consts.DefaultDeclsPriority
	}
	// check custom rundat fields for presets
	if decl.RunDat == nil {
		decl.RunDat = make(map[string]any)
	}

	// match preset
	preset, exists := Presets[decl.Preset]
	if !exists {
		return fmt.Errorf("preset not found for preset name: %q", decl.Preset)
	}
	if preset.IsValidFunc == nil {
		return fmt.Errorf("IsValidFunc not defined for preset %q", decl.Preset)
	}
	err = preset.IsValidFunc(*decl, metadat)
	if err != nil {
		return err
	}

	// subs
	if preset.IsValidFunc == nil {
		return fmt.Errorf("SubsFunc not defined for preset %q", decl.Preset)
	}
	err = preset.SubsFunc(decl)
	if err != nil {
		return err
	}

	return nil
}
