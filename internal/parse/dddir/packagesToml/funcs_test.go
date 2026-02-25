package packagestoml

import (
	"os/user"
	"testing"

	"github.com/suxyio/declmysys/internal/parse"
)

func TestPkgsLoad(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Errorf("failed to get username: %v", err)
	}
	username := usr.Username

	tests := parse.TomlLoadTests{
		// common wrong cases
		"":              {Result: nil, ExpectErr: true},
		"packages = [":  {Result: nil, ExpectErr: true},
		"packages = []": {Result: nil, ExpectErr: true},
		"priority = 0":  {Result: nil, ExpectErr: true},

		// empty case
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

	parse.RunTomlLoadTest(t, tests, &Pkgs{})
}
