package subs

import "testing"

func TestApplySubs(t *testing.T) {
	tests := subsRulesTests{
		"foo": "foo",
		// this case tests for recursive subs, which shall not be allowed
		"foobarbaz": "foobazbar",
	}
	rules := SubsRules{
		"bar": "baz",
		"baz": "bar",
	}
	tests.run(t, rules)
}
