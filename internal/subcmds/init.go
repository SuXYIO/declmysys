package subcmds

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/otiai10/copy"
	"github.com/urfave/cli/v3"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/parse/metadata"
	"github.com/suxyio/declmysys/internal/templates"
)

func Init(ctx context.Context, cmd *cli.Command) error {
	// get ddir and stuff
	gc, err := globconf.GetGlobconf()
	if err != nil {
		return fmt.Errorf("failed to get global config: %v", err)
	}
	ddir, err := getDDir(gc, cmd)
	if err != nil {
		return fmt.Errorf("failed to get decldir: %v", err)
	}

	if !cmd.Bool("no-git") {
		// init git
		if !gitAvail() {
			return fmt.Errorf("git is not available")
		}
		if err := gitInit(ddir); err != nil {
			return fmt.Errorf("failed to init git repository: %v", err)
		}
	}

	// copy template
	if err := copy.Copy("Decl", ddir, copy.Options{
		FS:                templates.DDir,
		PermissionControl: copy.AddPermission(0664), // damn permissions almost ruined my filesystem
		Skip: func(srcinfo os.FileInfo, src, dest string) (bool, error) {
			fname := srcinfo.Name()
			// exclude
			if cmd.Bool("no-taplo") && (fname == ".taplo.toml" ||
				(fname == ".schemas" && srcinfo.IsDir())) {
				return true, nil
			}
			return false, nil
		},
	}); err != nil {
		return fmt.Errorf("failed to copy template to path: %v", err)
	}

	// correct version in metadata.toml
	// NOTE: This is acutally pretty bad design
	mdpath := filepath.Join(ddir, "metadata.toml")
	if dat, err := os.ReadFile(mdpath); err != nil {
		return fmt.Errorf("failed to read metadata.toml: %v", err)
	} else {
		// replace version
		repl := metadata.SubsRules{"{VERSION}": consts.Version}.ToReplacer()

		res := metadata.ApplySubsReplacer(string(dat), &repl)
		if err := os.WriteFile(mdpath, []byte(res), 0644); err != nil {
			return fmt.Errorf("failed to overwrite metadata.toml: %v", err)
		}
	}

	fmt.Println("Copied template files to", ddir)
	fmt.Println("Successfully initialized decldir in", ddir)

	return nil
}

func gitInit(path string) error {
	cmd := exec.Command("git", "init", path)
	// let it output to terminal directly
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func gitAvail() bool {
	cmd := exec.Command("git", "version")
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
