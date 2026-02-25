package substoml

import (
	"testing"
)

func TestSubstomlLoad(t *testing.T) {
	tests := TomlLoadTests{
		// common wrong
		`[global]
foo = `: {Result: nil, ExpectErr: true},
		`[global]
foo = []`: {Result: nil, ExpectErr: true},

		// empty case
		``: {Result: &SubsDef{SpecialHDDisable: false, CustomG: map[string]string{}, CustomPC: map[string]string{}}, ExpectErr: false},
		`disable_homedir_subs = false
[global]
[paths_cmds]`: {Result: &SubsDef{SpecialHDDisable: false, CustomG: map[string]string{}, CustomPC: map[string]string{}}, ExpectErr: false},

		`disable_homedir_subs = true
[global]
"foo" = "bar"
[paths_cmds]
"{~}" = "{HOME}"`: {Result: &SubsDef{SpecialHDDisable: true, CustomG: map[string]string{"foo": "bar"}, CustomPC: map[string]string{"{~}": "{HOME}"}}, ExpectErr: false},
	}

	RunTomlLoadTest(t, tests, &SubsDef{})
}
