package subs

// NOTE: The idea is to pass this thing around, to let functions know about user custom subs.
// Implies that the first file parsed in dddir should be subs.toml, since parsing others need it
type SubsDef struct {
	SpecialHDDisable bool
	CustomG          SubsRules
	CustomPC         SubsRules
}

// ApplyG applies global subs
func (sd SubsDef) ApplyG(s string) (string, error) {
	// PERF: converting replacers every time is not optimal, (and performance is the reason for creating replacers in the first place),
	// maybe implement a function that takes a batch of strings and replaces them with only one replacer conversion

	// Apply custom before default
	// Custom
	// global
	grepl := sd.CustomG.ToReplacer()
	s = ApplySubs(s, &grepl)

	// Defaults
	tmp, err := applyDefaultGSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp
	// special hd
	if !sd.SpecialHDDisable {
		tmp, err := applySpecialHDSubs(s)
		if err != nil {
			return "", err
		}
		s = tmp
	}

	return s, nil
}

// ApplyPC applies paths&cmds (including globals)
func (sd SubsDef) ApplyPC(s string) (string, error) {
	// can't just call ApplyG here, cuz must follow "apply custom before default" order
	// Custom
	grepl := sd.CustomG.ToReplacer()
	s = ApplySubs(s, &grepl)
	pcrepl := sd.CustomPC.ToReplacer()
	s = ApplySubs(s, &pcrepl)

	// Defaults
	tmp, err := applyDefaultGSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp
	tmp, err = applyDefaultPCSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp
	tmp, err = applySpecialHDSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp

	return s, nil
}
