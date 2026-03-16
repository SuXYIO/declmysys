package cmdtype

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCmdRun(t *testing.T) {
	tmpdir := t.TempDir()
	files := []string{"foo", "bar", "baz"}
	cmd := Cmd{"touch"}

	err := cmd.Run(CmdRunOptions{WorkingDir: tmpdir, AppendedArgs: files})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check if created
	for _, f := range files {
		fpath := filepath.Join(tmpdir, f)
		if _, err := os.Stat(fpath); errors.Is(err, os.ErrNotExist) {
			t.Errorf("expected file %s to be created", fpath)
		} else if err != nil {
			t.Errorf("error reading file status: %v", err)
		}
	}
}
