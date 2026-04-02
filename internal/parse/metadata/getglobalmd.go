package metadata

import (
	"os"
	"path/filepath"
)

func GetGlobalMetadata(ddirpath string) error {
	data, err := os.ReadFile(filepath.Join(ddirpath, "metadata.toml"))
	if err != nil {
		return err
	}

	return GlobalMetaData.Load(data)
}
