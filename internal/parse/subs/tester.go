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
func (tests subsRulesTests) run(t *testing.T, rules SubsRules) {
	repl := rules.ToReplacer()
	for in, out := range tests {
		res := ApplySubs(in, &repl)
		if res != out {
			t.Errorf(`expected "%v" for case "%v", got "%v"`, out, in, res)
		}
	}
}

// subsFuncTester helps write tests for subs functions
func (tests subsFuncTests) run(t *testing.T, subsfunc func(string) (string, error)) {
	for in, out := range tests {
		res, err := subsfunc(in)

		if out.ExpectErr && err == nil {
			t.Errorf(`expected error for case "%v", got nil`, in)
		} else if !out.ExpectErr && err != nil {
			t.Errorf(`unexpected error for case "%v", got error "%v"`, in, err)
		} else if res != out.Result {
			t.Errorf(`expected "%v" for case "%v", got "%v"`, out.Result, in, res)
		}
	}
}
