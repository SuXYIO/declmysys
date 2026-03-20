package subs

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

func LoadGlobalSD(data []byte) error {
	if GlobalSubsDef.initialized {
		return fmt.Errorf("global subsdef var already initialized")
	}

	if err := GlobalSubsDef.SubsDef.Load(data); err != nil {
		return err
	}

	GlobalSubsDef.initialized = true
	return nil
}

// Load parses the subs.toml data
func (s *SubsDef) Load(data []byte) error {
	// toml decode
	metadata, err := toml.Decode(string(data), s)
	if err != nil {
		return err
	}

	// replace default fields
	if !metadata.IsDefined("global") {
		s.CustomG = map[string]string{}
	}
	if !metadata.IsDefined("paths_cmds") {
		s.CustomPC = map[string]string{}
	}

	// subs
	// nah no need to substitude for subsdef

	// this is the simplest Load function in my program!

	return nil
}
