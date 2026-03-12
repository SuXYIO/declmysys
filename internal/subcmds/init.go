package subcmds

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/otiai10/copy"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/templates"
	"github.com/suxyio/declmysys/internal/utils"
)

type InitOpts struct {
	NoGit bool `long:"no-git" description:"Won't create the .git directory (via git init)"`
}

func Init(gc globconf.Globconf, mopts *MainOpts, opts *InitOpts) {
	if !opts.NoGit {
		// init git
		if !gitAvail() {
			utils.Panic("Git is not available, please ensure dependencies are installed correctly", nil, exitcode.SetupError)
		}
		err := gitInit(mopts.DDir)
		if err != nil {
			utils.Panic("Git initialize failed", err, exitcode.ExecError)
		}
	}

	// copy template
	err := copy.Copy("Decl", mopts.DDir, copy.Options{
		FS:                templates.DDir,
		PermissionControl: copy.AddPermission(0664), // damn permissions almost ruined my filesystem
		Skip: func(srcinfo os.FileInfo, src, dest string) (bool, error) {
			fname := srcinfo.Name()
			// exclude .gitkeep
			return fname == ".gitkeep", nil
		},
	})
	if err != nil {
		utils.Panic("copy template to path fail", err, exitcode.FileError)
	}
	fmt.Println("Copied template files to", mopts.DDir)
	fmt.Println("Successfully initialized decldir in", mopts.DDir)
}

func gitInit(path string) error {
	cmd := exec.Command("git", "init", path)
	// let it output to terminal directly
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func gitAvail() bool {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
