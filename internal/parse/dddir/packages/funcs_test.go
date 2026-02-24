package packages

import (
	"testing"

	"github.com/suxyio/declmysys/internal/parse"
)

func TestPkgsLoad(t *testing.T) {
	tests := parse.TomlLoadTests{
		// common wrong cases
		"":              {Result: &Pkgs{}, ExpectErr: true},
		"packages = [":  {Result: &Pkgs{}, ExpectErr: true},
		"packages = []": {Result: &Pkgs{}, ExpectErr: true},
		"priority = 0":  {Result: &Pkgs{}, ExpectErr: true},

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
    "foo",
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
			{Manager: manSpec{"bar", nil}, Packs: []string{"foo", "bar", "baz"}},
			{Manager: manSpec{"", []string{"sudo", "apt", "install"}}, Packs: []string{"abc", "def", "ghi"}},
		}, Priority: 42}, ExpectErr: false},
	}

	parse.RunTomlLoadTest(t, tests, &Pkgs{})
}
