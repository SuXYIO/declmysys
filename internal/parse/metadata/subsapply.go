package metadata

import (
	"fmt"
	"strings"
)

type SubsRules map[string]string

// SubsToReplacer turns SubsRules to Replacer
func (r SubsRules) SubsToReplacer() strings.Replacer {
	var pairs []string
	for k, v := range r {
		pairs = append(pairs, k)
		pairs = append(pairs, v)
	}
	return *strings.NewReplacer(pairs...)
}

// applySubsReplacer applies a set of subs rules to string
func applySubsReplacer(str string, repler *strings.Replacer) string {
	// NOTE: Passes the replacer in seems awkward, but this design allows for high performance,
	// no need to convert to replacer every time
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
	repl := GlobalMetaData.SubsDef.CustomRules.SubsToReplacer()
	s = applySubsReplacer(s, &repl)

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
