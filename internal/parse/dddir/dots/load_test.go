package dots

import (
	"os"
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

	tests := substoml.TomlLoadTests{
		``:          {Result: nil, ExpectErr: true},
		`name = "`:  {Result: nil, ExpectErr: true},
		`name = []`: {Result: nil, ExpectErr: true},

		`name = "foo"
preset = "stow"`: {
			Result:    &Desc{Name: "foo", Preset: "stow", Priority: consts.DefaultDotsPriority, RunDat: map[string]any{"datadir": "data"}},
			ExpectErr: false,
		},

		`name = "bar"
preset = "cmds"
priority = 99
[rundat]
cmds = [
	["sudo", "foo", "{HOME}/Foobar"],
	["bar", "baz"],
]
`: {
			Result: &Desc{
				Name:     "bar",
				Preset:   "cmds",
				Priority: 99,
				RunDat:   map[string]any{"cmds": []cmdtype.Cmd{{"sudo", "foo", userhomedir + "/Foobar"}, {"bar", "baz"}}},
			},
			ExpectErr: false,
		},

		`name = "bar"
preset = "gitclone"
priority = 99
[rundat]
url = "github.com/foobar/baz"
dest = "~/Foobar"`: {
			Result: &Desc{
				Name:     "bar",
				Preset:   "gitclone",
				Priority: 99,
				RunDat:   map[string]any{"url": "github.com/foobar/baz", "dest": userhomedir + "/Foobar"},
			},
			ExpectErr: false,
		},
	}

	substoml.RunTomlLoadTest(t, tests, &Desc{})
}
