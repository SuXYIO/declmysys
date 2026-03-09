package substoml

import (
	"reflect"
	"testing"
)

func TestSubstomlLoad(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    SubsDef
		wantErr bool
	}{
		// common wrong
		{"invalid syntax", `[global]
foo = `, SubsDef{}, true},
		{"wrong type", `[global]
foo = []`, SubsDef{}, true},

		// empty case
		{"empty", ``, SubsDef{SpecialHDDisable: false, CustomG: map[string]string{}, CustomPC: map[string]string{}}, false},
		{"default", `disable_homedir_subs = false
[global]
[paths_cmds]`, SubsDef{SpecialHDDisable: false, CustomG: map[string]string{}, CustomPC: map[string]string{}}, false},

		{"subs", `disable_homedir_subs = true
[global]
"foo" = "bar"
[paths_cmds]
"{~}" = "{HOME}"`, SubsDef{SpecialHDDisable: true, CustomG: map[string]string{"foo": "bar"}, CustomPC: map[string]string{"{~}": "{HOME}"}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sd SubsDef
			err := sd.Load([]byte(tt.data))
			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			}
			if err != nil {
				// don't check value if has error
				return
			}
			if !reflect.DeepEqual(sd, tt.want) {
				t.Errorf("want = %#v, got = %#v", tt.want, sd)
			}
		})
	}
}
