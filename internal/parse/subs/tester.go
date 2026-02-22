package subs

import (
	"testing"
)

type subsFuncRet struct {
	Result    string
	ExpectErr bool
}

type subsRulesTests map[string]string
type subsFuncTests map[string]subsFuncRet

// subsRulesTester helps write tests for subs rules
func subsRulesTester(t *testing.T, rules SubsRules, tests subsRulesTests) {
	repl := rules.ToReplacer()
	for in, out := range tests {
		res := ApplySubs(in, &repl)
		if res != out {
			t.Errorf(`expected "%v" for case "%v", got "%v"`, out, in, res)
		}
	}
}

// subsFuncTester helps write tests for subs functions
func subsFuncTester(t *testing.T, subsfunc func(string) (string, error), tests subsFuncTests) {
	for in, out := range tests {
		res, err := subsfunc(in)

		if out.ExpectErr && err == nil {
			t.Errorf(`error expected for case "%v", got nil`, in)
		} else if !out.ExpectErr && err != nil {
			t.Errorf(`error not expected for case "%v", got error "%v"`, in, err)
		} else if res != out.Result {
			t.Errorf(`expected "%v" for case "%v", got "%v"`, out.Result, in, res)
		}
	}
}
