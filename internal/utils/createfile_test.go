package utils

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCreateFile(t *testing.T) {
	tmpdirpath := t.TempDir()

	// creates file under TMPDIR/foobar/baz.txt
	fpath := filepath.Join(tmpdirpath, "foobar", "baz.txt")
	content := []byte(`the quick brown fox
jumps over the lazy dog.`)
	if err := CreateFile(fpath, content); err != nil {
		t.Errorf("CreateFile failed: %v", err)
	}

	// checks if exists
	if _, err := os.Stat(fpath); errors.Is(err, os.ErrNotExist) {
		// doesn't exist
		t.Errorf("expected to create file %s, but does not exist", fpath)
	} else if err != nil {
		// error getting file status
		t.Errorf("error reading file status: %v", err)
	} else {
		// check if contents match
		data, err := os.ReadFile(fpath)
		if err != nil {
			t.Errorf("failed to read file %s: %v", fpath, err)
		}
		if !reflect.DeepEqual(content, data) {
			t.Errorf("file contents does not match: expected %q, got %q", content, data)
		}
	}
}
