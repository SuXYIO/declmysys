package subs

import (
	"strings"
)

type SubsRules map[string]string

// RulesToReplacer turns SubsRules to Replacer
func (r SubsRules) ToReplacer() strings.Replacer {
	var pairs []string
	for k, v := range r {
		pairs = append(pairs, k)
		pairs = append(pairs, v)
	}
	return *strings.NewReplacer(pairs...)
}

// ApplySubs applies a set of subs rules to string
func ApplySubs(str string, repler *strings.Replacer) string {
	// NOTE: Passes the replacer in seems awkward, but this design allows for high performance,
	// no need to convert to replacer every time
	return repler.Replace(str)
}

/*
Example Usage:
	toBeSubed := []string{"foobar", "barbaz"}
	rules := SubsRules{
		"foo": "bar",
		"bar": "foo"
	}
	var results []string

	rules.ToReplacer()
	for _, s := range toBeSubed {
		// this saves a lot of performance if there's a lot to be subed
		results.append(results, ApplySubs(s, repler))
	}
*/
