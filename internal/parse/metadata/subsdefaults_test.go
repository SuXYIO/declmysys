package metadata

import (
	"os"
	"os/user"
	"testing"
)

func TestApplySpecialHDSubs(t *testing.T) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("unable to get user home dir:", err)
	}

	tests := subsFuncTests{
		"foo":                               {"foo", false},
		"~/foo ~foo":                        {homedir + "/foo ~foo", false},
		"foo~ ~foo ~foo~ ~/foo~":            {"foo~ ~foo ~foo~ " + homedir + "/foo~", false},
		"~/foo ~ ~~baz~~ ~/foobar ~/barbaz": {homedir + "/foo " + homedir + " ~~baz~~ " + homedir + "/foobar " + homedir + "/barbaz", false},
	}

	tests.run(t, applyHDSubs)
}

func TestApplyDefaultSubs(t *testing.T) {
	userinfo, err := user.Current()
	if err != nil {
		t.Fatal("unable to get user info:", err)
	}
	username := userinfo.Username
	name := userinfo.Name
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatal("unable to get hostname info:", err)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		t.Fatal("unable to get homedir:", err)
	}
	confdir, err := os.UserConfigDir()
	if err != nil {
		t.Fatal("unable to get confdir:", err)
	}
	cachedir, err := os.UserCacheDir()
	if err != nil {
		t.Fatal("unable to get cachedir:", err)
	}
	tmpdir := os.TempDir()

	substr := map[string]string{
		"foo":        "foo",
		"{NAME}":     name,
		"{USERNAME}": username,
		"{HOSTNAME}": hostname,
		"{HOME}":     homedir,
		"{CONF}":     confdir,
		"{CONFIG}":   confdir,
		"{CACHE}":    cachedir,
		"{TMP}":      tmpdir,
	}

	tests := subsFuncTests{}
	for from, to := range substr {
		tests["foo"+from+"/bar"] = subsFuncRet{"foo" + to + "/bar", false}
	}

	tests.run(t, ApplyDefaultsSubs)
}
