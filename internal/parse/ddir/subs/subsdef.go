package subs

import (
	"fmt"
)

// Implies that the first file parsed in ddir should be subs.toml, since parsing others need it
type SubsDef struct {
	SpecialHDDisable bool      `toml:"disable_homedir_subs"`
	CustomG          SubsRules `toml:"global"`
	CustomPC         SubsRules `toml:"paths_cmds"`
}

// The global var that stores subsdef
var GlobalSubsDef struct {
	SubsDef     SubsDef
	Initialized bool
}

// ApplyG applies global subs
func ApplyG(s string) (string, error) {
	// PERF: converting replacers every time is not optimal, (and performance is the reason for creating replacers in the first place),
	// maybe implement a function that takes a batch of strings and replaces them with only one replacer conversion

	if !GlobalSubsDef.Initialized {
		return "", fmt.Errorf("global subsdef var not initialized")
	}

	// Apply custom before default
	// Custom
	// global
	grepl := GlobalSubsDef.SubsDef.CustomG.ToReplacer()
	s = ApplySubs(s, &grepl)

	// Defaults
	if tmp, err := applyDefaultGSubs(s); err != nil {
		return "", err
	} else {
		s = tmp
	}
	// special hd
	if !GlobalSubsDef.SubsDef.SpecialHDDisable {
		if tmp, err := applySpecialHDSubs(s); err != nil {
			return "", err
		} else {
			s = tmp
		}
	}

	return s, nil
}

// ApplyPC applies paths&cmds (including globals)
func ApplyPC(s string) (string, error) {
	if !GlobalSubsDef.Initialized {
		return "", fmt.Errorf("global subsdef var not initialized")
	}

	// can't just call ApplyG here, cuz must follow "apply custom before default" order
	// Custom
	gRepl := GlobalSubsDef.SubsDef.CustomG.ToReplacer()
	s = ApplySubs(s, &gRepl)
	pcRepl := GlobalSubsDef.SubsDef.CustomPC.ToReplacer()
	s = ApplySubs(s, &pcRepl)

	// Defaults
	if tmp, err := applyDefaultGSubs(s); err != nil {
		return "", err
	} else {
		s = tmp
	}
	if tmp, err := applyDefaultPCSubs(s); err != nil {
		return "", err
	} else {
		s = tmp
	}
	if !GlobalSubsDef.SubsDef.SpecialHDDisable {
		if tmp, err := applySpecialHDSubs(s); err != nil {
			return "", err
		} else {
			s = tmp
		}
	}

	return s, nil
}
