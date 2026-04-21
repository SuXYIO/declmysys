package subcmds

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"

	"github.com/otiai10/copy"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/parse/metadata"
	"github.com/suxyio/declmysys/internal/templates"
	"github.com/suxyio/declmysys/internal/utils"
)

type InitOpts struct {
	NoGit   bool `long:"no-git" description:"Won't create the .git directory (via git init)"`
	NoTaplo bool `long:"no-taplo" description:"Won't create the .taplo.toml file"`
}

func Init(gc globconf.Globconf, mopts MainOpts, opts InitOpts) {
	// ddir doesn't exist
	if DDirExist(mopts.DDir) {
		if dir, err := os.ReadDir(mopts.DDir); err != nil {
			utils.Panic(fmt.Sprintf("failed to read ddir %s", mopts.DDir), nil, exitcode.FileError)
		} else {
			if len(dir) > 0 {
				utils.Panic(fmt.Sprintf("ddir already %s exists and is not empty", mopts.DDir), nil, exitcode.FileError)
			}
		}
	}

	if !opts.NoGit {
		// init git
		if !gitAvail() {
			utils.Panic("Git is not available, please ensure dependencies are installed correctly", nil, exitcode.SetupError)
		}
		if err := gitInit(mopts.DDir); err != nil {
			utils.Panic("Git initialize failed", err, exitcode.ExecError)
		}
	}

	// copy template
	if err := copy.Copy("Decl", mopts.DDir, copy.Options{
		FS:                templates.DDir,
		PermissionControl: copy.AddPermission(0664), // damn permissions almost ruined my filesystem
		Skip: func(srcinfo os.FileInfo, src, dest string) (bool, error) {
			fname := srcinfo.Name()
			// exclude .gitkeep
			if opts.NoTaplo && (fname == ".taplo.toml" ||
				(fname == ".schemas" && srcinfo.IsDir())) {
				return true, nil
			}
			return false, nil
		},
	}); err != nil {
		utils.Panic("copy template to path fail", err, exitcode.FileError)
	}

	// correct version in metadata.toml
	// NOTE: This is acutally pretty bad design
	mdpath := filepath.Join(mopts.DDir, "metadata.toml")
	if dat, err := os.ReadFile(mdpath); err != nil {
		utils.Panic("error reading metadata.toml", err, exitcode.FileError)
	} else {
		// replace version
		info, ok := debug.ReadBuildInfo()
		if !ok {
			utils.Panic("error reading build info", nil, exitcode.Unknown)
		}
		repl := metadata.SubsRules{"{VERSION}": info.Main.Version}.ToReplacer()

		res := metadata.ApplySubsReplacer(string(dat), &repl)
		if err := os.WriteFile(mdpath, []byte(res), 0644); err != nil {
			utils.Panic("error overwriting metadata.toml", err, exitcode.FileError)
		}
	}

	fmt.Println("Copied template files to", mopts.DDir)
	fmt.Println("Successfully initialized decldir in", mopts.DDir)
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
