package subcmds

import (
	"context"
	"fmt"

	"github.com/suxyio/declmysys/internal/parse/decls"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, cmd *cli.Command) error {
	// get ddir and stuff
	gc, err := globconf.GetGlobconf(cmd.String("config"))
	if err != nil {
		return fmt.Errorf("failed to get global config: %v", err)
	}
	ddir, err := getDDir(gc, cmd)
	if err != nil {
		return fmt.Errorf("failed to get decldir: %v", err)
	}
	priority := cmd.IntArg("priority")
	dry := cmd.Bool("dry")

	declss, err := getDeclsData(ddir)
	if err != nil {
		return fmt.Errorf("failed to get decls data: %v", err)
	}

	// run
	if priority == 0 {
		if dry {
			fmt.Printf("Running %s (dry):\n", gc.DDir)
		} else {
			fmt.Printf("Running %s:\n", gc.DDir)
		}
	} else {
		if dry {
			fmt.Printf("Running %s (dry, priority %d):\n", gc.DDir, priority)
		} else {
			fmt.Printf("Running %s (priority %d):\n", gc.DDir, priority)
		}
	}

	declopts := decls.DeclsRunOpts{
		Indent:         1,
		FilterPriority: priority,
	}
	if dry {
		declopts.Dry = true
	}

	if err := declss.Run(declopts); err != nil {
		return fmt.Errorf("error running decls: %v", err)
	}

	return nil
}
