package substoml

import (
	"github.com/BurntSushi/toml"
)

// Load parses the subs.toml data
// sd option will be ignored, pass whatever you want,
// subs.toml is not loaded so passing it is meaningless, the argument is just for interface compatability
func (s *SubsDef) Load(data []byte, _ SubsDef) error {
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
