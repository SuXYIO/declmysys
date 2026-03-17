package packagestoml

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestPkgsRun(t *testing.T) {
	t.Run("create file", func(t *testing.T) {
		tmpdir := t.TempDir()
		files := []string{
			"foo",
			"bar",
			"baz_{USERNAME}", // no subs
		}
		for i := range files {
			files[i] = filepath.Join(tmpdir, files[i])
		}

		pkgs := Pkgs{
			Packages: []PacksSpec{
				{
					Manager: manSpec{
						Preset:    "",
						CustomCmd: []string{"touch"},
					},
					Packs: files,
				},
			},
		}

		err := pkgs.Run()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		for _, f := range files {
			if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
				t.Errorf("expected %s to be created, not found", f)
			} else if err != nil {
				t.Errorf("error reading file status for %s", f)
			}
		}
	})
}
