package globconf

import (
	"os"
	"os/user"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
)

func TestGlobconfLoad(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		t.Fatalf("failed to get username: %v", err)
	}
	username := usr.Username
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("failed to get homedir: %v", err)
	}
	defaultDDDir, err := consts.DefaultDDDirPath()
	if err != nil {
		t.Fatalf("failed to get default dddir: %v", err)
	}

	tests := substoml.TomlLoadTests{
		// common wrong
		`dotdecldir = "`:  {Result: nil, ExpectErr: true},
		`dotdecldir = 0`:  {Result: nil, ExpectErr: true},
		`dotdecldir = []`: {Result: nil, ExpectErr: true},

		// empty case
		``:                {Result: &Globconf{DDDir: defaultDDDir}, ExpectErr: false},
		`dotdecldir = ""`: {Result: &Globconf{DDDir: ""}, ExpectErr: false},

		// subs
		`dotdecldir = "~/dot_{USERNAME}"`: {Result: &Globconf{DDDir: homedir + "/dot_" + username}},
	}

	substoml.RunTomlLoadTest(t, tests, &Globconf{})
}
