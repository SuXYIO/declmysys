package subs

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

	subsFuncTester(t, ApplySpecialHDSubs, tests)
}

func TestApplyDefaultGSubs(t *testing.T) {
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

	tests := subsFuncTests{
		"foo":              {"foo", false},
		"foo{NAME}bar":     {"foo" + name + "bar", false},
		"foo{USERNAME}bar": {"foo" + username + "bar", false},
		"foo{HOSTNAME}bar": {"foo" + hostname + "bar", false},
	}

	subsFuncTester(t, ApplyDefaultGSubs, tests)
}

func TestApplyDefaultPCSubs(t *testing.T) {
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

	tests := subsFuncTests{
		"foo": {"foo", false},
	}
	substr := map[string]string{"{HOME}": homedir, "{CONF}": confdir, "{CONFIG}": confdir, "{CACHE}": cachedir, "{TMP}": tmpdir}
	for from, to := range substr {
		tests["foo"+from+"/bar"] = subsFuncRet{"foo" + to + "/bar", false}
	}

	subsFuncTester(t, ApplyDefaultPCSubs, tests)
}
