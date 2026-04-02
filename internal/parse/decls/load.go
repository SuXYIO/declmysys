package decls

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
)

func (decl *Decl) Load(path string) error {
	// desc.toml
	descdata, err := os.ReadFile(filepath.Join(path, "desc.toml"))
	if err != nil {
		return err
	}
	decl.descPath = path
	err = decl.loadDesc(descdata)
	if err != nil {
		return err
	}

	return nil
}

func (decl *Decl) loadDesc(data []byte) error {
	md, err := toml.Decode(string(data), decl)
	if err != nil {
		return err
	}

	// common isvalid & defaults
	if err := decl.descCommonIsValid(md); err != nil {
		return err
	}
	if err := decl.descCommonSetDefaults(md); err != nil {
		return err
	}

	// match preset
	preset, exists := presets[decl.Preset]
	if !exists {
		return fmt.Errorf("preset not found for preset name: %q", decl.Preset)
	}
	if preset.PreFunc == nil {
		return fmt.Errorf("PreFunc not defined for preset %q", decl.Preset)
	}
	err = preset.PreFunc(decl, md)
	if err != nil {
		return err
	}

	return nil
}

func (decl Decl) descCommonIsValid(md toml.MetaData) error {
	// only checks for common default fields shared among presets
	// check default fields
	if !md.IsDefined("name") {
		return fmt.Errorf("must specify name for dot desc")
	}
	if !md.IsDefined("preset") {
		return fmt.Errorf("must specify preset for dot desc")
	}
	return nil
}

func (decl *Decl) descCommonSetDefaults(md toml.MetaData) error {
	if !md.IsDefined("priority") {
		decl.Priority = consts.DefaultDeclPriority
	}
	if !md.IsDefined("pwd") || decl.Pwd == "" {
		decl.Pwd = decl.descPath
	}
	return nil
}
