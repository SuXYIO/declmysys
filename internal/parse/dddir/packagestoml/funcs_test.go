package packagestoml

import (
	"os/user"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
)

func TestPkgsLoad(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Errorf("failed to get username: %v", err)
	}
	username := usr.Username

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
  { manager = ["sudo", "apt", "install"], packs = [
    "abc",
    "def",
    "ghi",
  ] },
]
priority = 42`: {Result: &Pkgs{Packages: []PacksSpec{
			{Manager: manSpec{"bar", nil}, Packs: []string{username, "bar", "baz"}},
			{Manager: manSpec{"", []string{"sudo", "apt", "install"}}, Packs: []string{"abc", "def", "ghi"}},
		}, Priority: 42}, ExpectErr: false},
	}

	substoml.RunTomlLoadTest(t, tests, &Pkgs{})
}
