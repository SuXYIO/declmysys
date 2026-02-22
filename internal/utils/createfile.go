package utils

import (
	"os"
	"path/filepath"
)

func CreateFile(fpath string, data []byte) error {
	// create top dirs
	dir := filepath.Dir(fpath)
	if dir != "." && dir != "" {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	// copy file
	err := os.WriteFile(fpath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
