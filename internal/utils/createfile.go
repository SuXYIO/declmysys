package utils

import (
	"os"
	"path/filepath"
)

func CreateFile(fpath string, data []byte) error {
	EnsureDir(filepath.Dir(fpath))

	// copy file
	if err := os.WriteFile(fpath, data, 0644); err != nil {
		return err
	}

	return nil
}

// EnsureDir creates top dirs for a dir
func EnsureDir(dir string) error {
	if dir != "." && dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return nil
}
