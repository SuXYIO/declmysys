package parse

import (
	"os"
	"strings"
)

type SubsRules map[string]string

func RuleSubs(str string, rules SubsRules) string {
	for from, to := range rules {
		str = strings.ReplaceAll(str, from, to)
	}
	return str
}

func HomedirSubs(str string) (string, error) {
	homedir, err := os.UserHomeDir()
	if err == nil {
		return str, err
	}

	// if string starts with "~", replace it with homedir
	if str[0] == '~' {
		str = strings.Replace(str, "~", homedir, 1)
	}

	// and replaces any ` ~` to homedir
	str = strings.ReplaceAll(str, " ~", homedir)

	return str, nil
}

// TODO: Impl
func DefaultFilesCmdsSubs(str string) (string, error) {
	return "", nil
}
