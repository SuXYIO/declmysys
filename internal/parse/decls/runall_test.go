package decls

import (
	"io"
	"os"
	"os/user"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/suxyio/declmysys/internal/parse/cmdtype"
)

func TestDeclsRun(t *testing.T) {
	t.Run("echo order test", func(t *testing.T) {
		userinfo, err := user.Current()
		if err != nil {
			t.Fatalf("failed to get user info: %v", err)
		}
		tmpdir := t.TempDir()
		expected := []byte(`Born of God and Void.
You shall seal the blinding light that plagues their dreams.
You are the Vessel.
You are ` + userinfo.Username + ".\n")

		// run
		decls := Decls{
			{
				Name:     "create foo.txt",
				Preset:   "cmds",
				Priority: 16,
				Args: map[string]any{
					"cmds": []cmdtype.Cmd{
						{"touch", "foo.txt"},
					},
				},
			},
			{
				Name:     "echo line1",
				Preset:   "cmds",
				Priority: 8,
				Args: map[string]any{
					"cmds": []cmdtype.Cmd{
						// used redirect here, so must use shell
						{"sh", "-c", "echo 'Born of God and Void.' >> foo.txt"},
					},
				},
			},
			{
				Name:     "echo line2",
				Preset:   "cmds",
				Priority: 7,
				Args: map[string]any{
					"cmds": []cmdtype.Cmd{
						{"sh", "-c", "echo 'You shall seal the blinding light that plagues their dreams.' >> foo.txt"},
					},
				},
			},
			{
				Name:     "echo line3",
				Preset:   "cmds",
				Priority: 6,
				Args: map[string]any{
					"cmds": []cmdtype.Cmd{
						{"sh", "-c", "echo 'You are the Vessel.' >> foo.txt"},
					},
				},
			},
			{
				Name:     "echo line4",
				Preset:   "cmds",
				Priority: 5,
				Args: map[string]any{
					"cmds": []cmdtype.Cmd{
						{"sh", "-c", "echo 'You are {USERNAME}.' >> foo.txt"},
					},
				},
			},
		}

		if err := decls.Run(DeclsRunOpts{
			noPrint:        true,
			WorkingDir:     tmpdir,
			FilterPriority: -1,
		}); err != nil {
			t.Errorf("unexpected error running RunDecls: %v", err)
		}

		if got, err := os.ReadFile(filepath.Join(tmpdir, "foo.txt")); err != nil {
			t.Errorf("unexpected error reading file: %v", err)
		} else if !reflect.DeepEqual(got, expected) {
			t.Errorf("file contents does not match, want %q, got %q", expected, got)
		}
	})
}

// BenchmarkDeclsRunParse benchmarks the speed of parsing decls
func BenchmarkDeclsRunParse(b *testing.B) {
	decls := Decls{}
	for _, s := range []string{"foo", "bar", "baz", "quz"} {
		decls = append(decls, Decl{
			Name:     s,
			Preset:   "cmds",
			Priority: 42,
			Args: map[string]any{
				"cmds": []cmdtype.Cmd{
					{"touch", s},
				},
			},
		})
	}

	for b.Loop() {
		tmpdir := b.TempDir()
		if err := decls.Run(DeclsRunOpts{
			RedirectStdout: io.Discard,
			RedirectStderr: io.Discard,
			WorkingDir:     tmpdir,
			noPrint:        true,
		}); err != nil {
			b.Errorf("unexpected error running RunDecls: %v", err)
		}
	}
}
