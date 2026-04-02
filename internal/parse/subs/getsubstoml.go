package subs

import (
	"os"
	"path/filepath"
)

// GetSubsToml is a helper function to load the subs.toml into global subs,
// which has no need for return (except error) since it's in the global subs
func GetSubsToml(ddir string) error {
	subsTomlData, err := os.ReadFile(filepath.Join(ddir, "subs.toml"))
	if err != nil {
		return err
	}
	if err := LoadGlobalSD(subsTomlData); err != nil {
		return err
	}

	return nil
}
