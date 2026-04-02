package cmdtype

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/suxyio/declmysys/internal/parse/metadata"
)

func TestCmdRun(t *testing.T) {
	t.Run("create-file", func(t *testing.T) {
		userinfo, err := user.Current()
		if err != nil {
			t.Fatalf("failed to get user info: %v", err)
		}

		// test with empty global subsdef var
		metadata.GlobalMetaData.Load([]byte(""))

		tmpdir := t.TempDir()
		fnames := []string{"foo", "bar", "baz_{USERNAME}"}
		expected := []string{"foo", "bar", "baz_" + userinfo.Username}
		cmd := Cmd{"touch"}

		if err := cmd.Run(CmdRunOptions{
			WorkingDir:   tmpdir,
			AppendedArgs: fnames,
		}); err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// check if created
		for _, f := range expected {
			fpath := filepath.Join(tmpdir, f)
			if _, err := os.Stat(fpath); errors.Is(err, os.ErrNotExist) {
				// also print what files are created for debug info
				dir, err := os.ReadDir(tmpdir)
				if err != nil {
					t.Errorf("expected file %s to be created, and failed to read dir", fpath)
				}

				var gotfiles []string
				for _, f := range dir {
					gotfiles = append(gotfiles, f.Name())
				}

				t.Errorf("expected file %s to be created, got files: %v", fpath, gotfiles)
			} else if err != nil {
				t.Errorf("error reading file status: %v", err)
			}
		}
	})
}
