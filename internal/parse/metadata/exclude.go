package metadata

import (
	"fmt"
	"regexp"
)

func IsExcluded(fname []byte) (bool, error) {
	if !GlobalMetaData.initialized {
		return false, fmt.Errorf("global metadata not initialized")
	}

	for _, r := range GlobalMetaData.Exclude {
		if is, err := regexp.Match(r, fname); err != nil {
			return false, err
		} else if is {
			return true, nil
		}
	}

	return false, nil
}
