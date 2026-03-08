package packagestoml

import (
	"os"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
)

func TestPkgsLoad(t *testing.T) {
	userhomedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get user home dir: %v", err)
	}

	tests := substoml.TomlLoadTests{
		``:              {Result: &Pkgs{Packages: []PacksSpec{}, Priority: consts.DefaultPackagesPriority}, ExpectErr: false},
		`packages = [`:  {Result: nil, ExpectErr: true},
		`packages = []`: {Result: &Pkgs{Packages: []PacksSpec{}, Priority: consts.DefaultPackagesPriority}, ExpectErr: false},
		`packages = []
priority = 0`: {Result: &Pkgs{Packages: []PacksSpec{}, Priority: 0}, ExpectErr: false},

		// missing manager / packs
		`packages = [
  { packs = [
    "foo",
    "bar",
    "baz",
  ] },
  { manager = ["sudo", "apt", "install"] },
]
priority = 42`: {Result: nil, ExpectErr: true},

		// normal (preset manager & cmdlist manager spec)
		`packages = [
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
priority = 42`: {Result: &Pkgs{Packages: []PacksSpec{
			{Manager: manSpec{"bar", nil}, Packs: []string{"{USERNAME}", "bar", "baz"}}, // packs shall not be subed
			{Manager: manSpec{"", []string{"sudo", "apt", userhomedir}}, Packs: []string{"abc", "def", "ghi"}},
		}, Priority: 42}, ExpectErr: false},
	}

	substoml.RunTomlLoadTest(t, tests, &Pkgs{})
}
