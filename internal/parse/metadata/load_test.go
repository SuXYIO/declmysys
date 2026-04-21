package metadata

import (
	"reflect"
	"testing"
)

func TestMetadataLoad(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    Metadata
		wantErr bool
	}{
		// common wrong
		{"invalid syntax", `foo = `, Metadata{}, true},
		{"wrong type", `[subs.rules]
foo = []`, Metadata{}, true},

		{"empty", ``, Metadata{
			Exclude: []string{
				"^\\.git$",
			},
			SubsDef: SubsDef{
				SpecialHDDisable: false,
				CustomRules:      map[string]string{},
			},
			initialized: true,
		}, false},

		{"common", `exclude = [
	"^\\.myexclude$",
	"^\\.git$"
]
[subs]
disable_homedir_subs = true
[subs.rules]
"foo" = "bar"
"{~}" = "{HOME}"`, Metadata{
			Exclude: []string{
				"^\\.myexclude$",
				"^\\.git$",
			},
			SubsDef: SubsDef{
				SpecialHDDisable: true,
				CustomRules:      map[string]string{"foo": "bar", "{~}": "{HOME}"},
			},
			initialized: true,
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sd Metadata
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
