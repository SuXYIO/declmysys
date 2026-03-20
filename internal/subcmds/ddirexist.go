package subcmds

import "os"

// DDirExist checks if decldir exists
func DDirExist(path string) bool {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return true
	}
	return false
}
