package packagestoml

import (
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/ddir/substoml"
)

func TestPkgsLoad(t *testing.T) {
	// test with empty global subsdef var
	err := substoml.LoadGlobalSD([]byte(""))
	if err != nil {
		t.Fatalf("failed to load global subsdef: %v", err)
	}

	tests := []struct {
		name    string
		data    string
		want    Pkgs
		wantErr bool
	}{
		{"empty", ``, Pkgs{Packages: []PacksSpec{}, Priority: consts.DefaultPackagesPriority}, false},
		{"invalid syntax", `packages = [`, Pkgs{}, true},
		{"empty packages", `packages = []`, Pkgs{Packages: []PacksSpec{}, Priority: consts.DefaultPackagesPriority}, false},
		{"empty packages and priority", `packages = []
priority = 0`, Pkgs{Packages: []PacksSpec{}, Priority: 0}, false},

		// missing manager or packs
		{"missing manager or packs", `packages = [
  { packs = [
    "foo",
    "bar",
    "baz",
  ] },
  { manager = ["sudo", "apt", "install"] },
]
priority = 42`, Pkgs{}, true},

		// normal (preset manager & cmdlist manager spec)
		{"common with subs", `packages = [
  { manager = "bar", packs = [
    "{USERNAME}",
    "bar",
    "baz",
  ] },
  { manager = ["sudo", "apt", "{HOME}"], packs = [
    "abc",
    "def",
    "ghi",
  ] },
]
priority = 42`,
			Pkgs{Packages: []PacksSpec{
				{Manager: manSpec{"bar", nil}, Packs: []string{"{USERNAME}", "bar", "baz"}},                     // packs shall not be subed
				{Manager: manSpec{"", []string{"sudo", "apt", "{HOME}"}}, Packs: []string{"abc", "def", "ghi"}}, // cmds are not subed til run
			}, Priority: 42},
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pkgs Pkgs
			err := pkgs.Load([]byte(tt.data))
			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			}
			if err != nil {
				// don't check value if has error
				return
			}
			if !reflect.DeepEqual(pkgs, tt.want) {
				t.Errorf("want = %#v, got = %#v", tt.want, pkgs)
			}
		})
	}
}
