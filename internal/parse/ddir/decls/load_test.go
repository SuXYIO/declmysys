package decls

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/ddir/substoml"
)

func TestDeclsLoad(t *testing.T) {
	// test with empty global subsdef var
	substoml.LoadGlobalSD([]byte(""))

	tests := []struct {
		name     string
		datapath string // name of decl dir under testdata directory, e.g. "foo" for "testdata/foo"
		want     Decl
		wantErr  bool
	}{
		{"empty", "empty", Decl{}, true},
		{"invalid syntax", "invalid-syntax", Decl{}, true},
		{"wrong type", "wrong-type", Decl{}, true},

		{"simple",
			"simple",
			Decl{Name: "foo", Preset: "stow", Priority: consts.DefaultDeclsPriority, RunDat: map[string]any{"datadir": "data"}},
			false,
		},

		{"common",
			"common",
			Decl{
				Name:     "bar",
				Preset:   "cmds",
				Priority: 99,
				RunDat:   map[string]any{"cmds": []cmdtype.Cmd{{"sudo", "foo", "{HOME}/Foobar"}, {"bar", "baz"}}},
			},
			false,
		},

		{"subs", // changed subs to be done with run, so no subs for load
			"subs",
			Decl{
				Name:     "bar",
				Preset:   "gitclone",
				Priority: 99,
				RunDat:   map[string]any{"url": "github.com/foobar/baz", "dest": "~/Foobar"},
			},
			false,
		},
		{"touch",
			"touch",
			Decl{
				Name:     "foobar",
				Preset:   "cmds",
				Priority: 0,
				RunDat: map[string]any{
					"cmds": []cmdtype.Cmd{
						{"touch", "foo"},
						{"touch", "bar", "baz_{USERNAME}"},
					},
				},
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var decl Decl
			err := decl.Load(filepath.Join("testdata", tt.datapath))
			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			}
			if err != nil {
				// don't check value if has error
				return
			}
			if !reflect.DeepEqual(decl, tt.want) {
				t.Errorf("want = %#v, got = %#v", tt.want, decl)
			}
		})
	}
}
