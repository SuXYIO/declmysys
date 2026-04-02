package metadata

import (
	"os"
	"os/user"
	"strings"
)

// these are really shitty naming, but can't think of better ones
// won't expose these outside this module

// applyHDSubs applies the special homedir subs
func applyHDSubs(str string) (string, error) {
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

func applyDefaultsSubsNoHD(str string) (string, error) {
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

// ApplyDefaultSubs applies the default subs rules to string, note that it does not include special hd subs
func ApplyDefaultsSubs(str string) (string, error) {
	if tmp, err := applyHDSubs(str); err != nil {
		return str, err
	} else {
		str = tmp
	}

	if tmp, err := applyDefaultsSubsNoHD(str); err != nil {
		return str, err
	} else {
		str = tmp
	}

	return str, nil
}
