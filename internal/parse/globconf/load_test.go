package globconf

import (
	"os"
	"os/user"
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/consts"
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

	tests := []struct {
		name    string
		data    string // actually []byte in the argument, but I'll use string and put conversion in loop
		want    Globconf
		wantErr bool
	}{
		// common wrong
		{"invalid syntax", `decldir = "`, Globconf{}, true},
		{"wrong type", `decldir = 0`, Globconf{}, true},
		{"wrong type", `decldir = []`, Globconf{}, true},

		// empty case
		{"empty", ``, Globconf{DDir: consts.DefaultDDirPath}, false},
		{"empty value", `decldir = ""`, Globconf{DDir: ""}, false},

		// subs
		{"subs", `decldir = "~/dot_{USERNAME}"`, Globconf{DDir: homedir + "/dot_" + username}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gc Globconf
			err := gc.Load([]byte(tt.data))
			if tt.wantErr != (err != nil) {
				t.Errorf("wantErr = %v, err = %v", tt.wantErr, err)
				return
			}
			if err != nil {
				// don't check value if has error
				return
			}
			if !reflect.DeepEqual(gc, tt.want) {
				t.Errorf("want = %#v, got = %#v", tt.want, gc)
			}
		})
	}
}
