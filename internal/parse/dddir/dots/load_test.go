package dots

import (
	"os"
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
)

func TestDescLoad(t *testing.T) {
	userhomedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get user home dir: %v", err)
	}

	// test with empty global subsdef var
	substoml.LoadGlobalSD([]byte(""))

	tests := []struct {
		name    string
		data    string
		want    Desc
		wantErr bool
	}{
		{"empty", ``, Desc{}, true},
		{"invalid syntax", `name = "`, Desc{}, true},
		{"wrong type", `name = []`, Desc{}, true},

		{"simple",
			`name = "foo"
preset = "stow"`,
			Desc{Name: "foo", Preset: "stow", Priority: consts.DefaultDotsPriority, RunDat: map[string]any{"datadir": "data"}},
			false,
		},

		{"common",
			`name = "bar"
preset = "cmds"
priority = 99
[rundat]
cmds = [
	["sudo", "foo", "{HOME}/Foobar"],
	["bar", "baz"],
]
`,
			Desc{
				Name:     "bar",
				Preset:   "cmds",
				Priority: 99,
				RunDat:   map[string]any{"cmds": []cmdtype.Cmd{{"sudo", "foo", userhomedir + "/Foobar"}, {"bar", "baz"}}},
			},
			false,
		},

		{"subs",
			`name = "bar"
preset = "gitclone"
priority = 99
[rundat]
url = "github.com/foobar/baz"
dest = "~/Foobar"`,
			Desc{
				Name:     "bar",
				Preset:   "gitclone",
				Priority: 99,
				RunDat:   map[string]any{"url": "github.com/foobar/baz", "dest": userhomedir + "/Foobar"},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var desc Desc
			err := desc.Load([]byte(tt.data))
			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			}
			if err != nil {
				// don't check value if has error
				return
			}
			if !reflect.DeepEqual(desc, tt.want) {
				t.Errorf("want = %#v, got = %#v", tt.want, desc)
			}
		})
	}
}
