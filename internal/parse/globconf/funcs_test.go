package globconf

import (
	"os"
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
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Errorf("failed to get homedir: %v", err)
	}

	tests := parse.TomlLoadTests{
		// common wrong
		"":                {Result: nil, ExpectErr: true},
		`dotdecldir = "`:  {Result: nil, ExpectErr: true},
		`dotdecldir = 0`:  {Result: nil, ExpectErr: true},
		`dotdecldir = []`: {Result: nil, ExpectErr: true},

		// empty case
		`dotdecldir = ""`: {Result: &Globconf{DDDir: ""}, ExpectErr: false},

		// subs
		`dotdecldir = "~/dot_{USERNAME}"`: {Result: &Globconf{DDDir: homedir + "/dot_" + username}},
	}

	parse.RunTomlLoadTest(t, tests, &Globconf{})
}
