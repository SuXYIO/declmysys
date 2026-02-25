package subs

import (
	"os"
	"os/user"
	"strings"
)

// these are really shitty naming, but can't think of better ones

// ApplySpecialHDSubs applies the special homedir subs
func ApplySpecialHDSubs(str string) (string, error) {
	// prevent index out of range for empty string
	if len(str) <= 0 {
		return "", nil
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		return str, err
	}

	// if string starts with "~", replace it with homedir
	if str[0] == '~' {
		str = homedir + str[1:]
	}

	// and replaces any " ~ " to " "+homedir+" "
	str = strings.ReplaceAll(str, " ~ ", " "+homedir+" ")

	// and replaces any " ~/" to " "+homedir
	str = strings.ReplaceAll(str, " ~/", " "+homedir+"/")

	return str, nil
}

// ApplyDefaultGSubs applies the default global subs rules to string
// see no point in using this though, before reading subs.toml, ApplyDefaultPC integrates this, and it won't be used after reading
func ApplyDefaultGSubs(str string) (string, error) {
	userinfo, err := user.Current()
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, "{USERNAME}", userinfo.Username)
	str = strings.ReplaceAll(str, "{NAME}", userinfo.Name)

	hostname, err := os.Hostname()
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, "{HOSTNAME}", hostname)

	return str, nil
}

// ApplyDefaultPCSubs applies the default paths & cmds subs rules to string
func ApplyDefaultPCSubs(str string) (string, error) {
	// first special homedir
	str, err := ApplySpecialHDSubs(str)
	if err != nil {
		return str, err
	}

	// then default fc subs
	homedir, err := os.UserHomeDir()
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, "{HOME}", homedir)

	confdir, err := os.UserConfigDir()
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, "{CONF}", confdir)
	str = strings.ReplaceAll(str, "{CONFIG}", confdir)

	cachedir, err := os.UserCacheDir()
	if err != nil {
		return str, err
	}
	str = strings.ReplaceAll(str, "{CACHE}", cachedir)

	tmpdir := os.TempDir() // and who knows why the fuck getting temp dir won't return errors, odd apis way to go
	str = strings.ReplaceAll(str, "{TMP}", tmpdir)

	return str, nil
}

// ApplyDefaultPC applies the default paths & cmds subs & global subs to string
func ApplyDefaultPC(s string) (string, error) {
	tmp, err := ApplyDefaultGSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp
	tmp, err = ApplyDefaultPCSubs(s)
	if err != nil {
		return "", err
	}
	s = tmp
	return s, nil
}
