package metadata

import (
	"github.com/BurntSushi/toml"
)

func (m *Metadata) Load(data []byte) error {
	// toml decode
	md, err := toml.Decode(string(data), m)
	if err != nil {
		return err
	}

	// replace default fields
	if !md.IsDefined("exclude") {
		m.Exclude = []string{
			"^\\.git$",
		}
	}
	if !md.IsDefined("subs", "rules") {
		m.SubsDef.CustomRules = map[string]string{}
	}
	if !md.IsDefined("subs", "disable_homedir_subs") {
		m.SubsDef.SpecialHDDisable = false
	}

	// this is the simplest Load function in my program!

	m.initialized = true

	return nil
}
