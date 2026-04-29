package subcmds

import (
	"bytes"
	"context"
	"fmt"

	"github.com/suxyio/declmysys/internal/parse/decls"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/utils"
	"github.com/urfave/cli/v3"
)

func List(ctx context.Context, cmd *cli.Command) error {
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

	declss, err := getDeclsData(ddir)
	if err != nil {
		return fmt.Errorf("failed to get decls data: %v", err)
	}

	// list
	var w bytes.Buffer
	if priority < 0 {
		fmt.Fprintf(&w, "Listing %s:\n", gc.DDir)
	} else {
		fmt.Fprintf(&w, "Listing %s (priority %d):\n", gc.DDir, priority)
	}
	if err := declss.List(&w, decls.DeclsListOpts{
		Indent:         1,
		FilterPriority: priority,
	}); err != nil {
		return fmt.Errorf("error listing decls: %v", err)
	}

	if err := utils.AutoPager(w.Bytes()); err != nil {
		return fmt.Errorf("failed to autopage: %v", err)
	}

	return nil
}
