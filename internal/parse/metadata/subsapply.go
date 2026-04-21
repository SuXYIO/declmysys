package metadata

import (
	"fmt"
	"strings"
)

type SubsRules map[string]string

// ToReplacer turns SubsRules to Replacer
func (r SubsRules) ToReplacer() strings.Replacer {
	var pairs []string
	for k, v := range r {
		pairs = append(pairs, k)
		pairs = append(pairs, v)
	}
	return *strings.NewReplacer(pairs...)
}

// ApplySubsReplacer applies a set of subs rules to string
func ApplySubsReplacer(str string, repler *strings.Replacer) string {
	// NOTE: Passes the replacer in seems awkward, but this design allows for high performance,
	// no need to convert to replacer every time
	// (actually used replacer only for non-recursive subs)
	return repler.Replace(str)
}

// SubsApply applies subs to string
func SubsApply(s string) (string, error) {
	// PERF: converting replacers every time is not optimal, (and performance is the reason for creating replacers in the first place),
	// maybe implement a function that takes a batch of strings and replaces them with only one replacer conversion

	if !GlobalMetaData.initialized {
		return "", fmt.Errorf("global subsdef var not initialized")
	}

	// Custom
	repl := GlobalMetaData.SubsDef.CustomRules.ToReplacer()
	s = ApplySubsReplacer(s, &repl)

	// Defaults
	if tmp, err := applyDefaultsSubsNoHD(s); err != nil {
		return "", err
	} else {
		s = tmp
	}

	// special hd
	if !GlobalMetaData.SubsDef.SpecialHDDisable {
		if tmp, err := applyHDSubs(s); err != nil {
			return "", err
		} else {
			s = tmp
		}
	}

	return s, nil
}
