package main

import (
	"context"
	"os"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/subcmds"
	"github.com/suxyio/declmysys/internal/utils"
	"github.com/urfave/cli/v3"
)

func main() {
	var cmd = &cli.Command{
		// metadata
		Name:    consts.Name,
		Usage:   consts.Desc,
		Version: consts.BuildInfo,
		Authors: []any{consts.Author},

		// options
		UseShortOptionHandling: true,
		EnableShellCompletion:  false, // won't want the user to use this accidentally

		// flags
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "config",
				Aliases:   []string{"C"},
				Value:     consts.DefaultGlobconfPath,
				TakesFile: true,
				Usage:     "specify the global config path",
			},
			&cli.StringFlag{
				Name:      "decldir",
				Aliases:   []string{"D"},
				Value:     consts.DefaultDDirPath,
				TakesFile: true,
				Usage:     "specify the decldir path, overrides decldir path from global config",
			},
		},

		// subcommands
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "inits new decldir",
				Action: subcmds.Init,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "no-git",
						Value: false,
						Usage: "won't initialize git repository",
					},
					&cli.BoolFlag{
						Name:  "no-taplo",
						Value: false,
						Usage: "won't initialize with taplo config and schemas",
					},
				},
			},
			{
				Name:   "list",
				Usage:  "lists decls",
				Action: subcmds.List,
				Arguments: []cli.Argument{
					&cli.IntArg{
						Name:      "priority",
						Value:     -1,
						UsageText: "show the procedures with this priority, show all if is negative",
					},
				},
			},
			{
				Name:   "run",
				Usage:  "execute decls",
				Action: subcmds.Run,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "dry",
						Aliases: []string{"d"},
						Value:   false,
						Usage:   "dry run",
					},
				},
				Arguments: []cli.Argument{
					&cli.IntArg{
						Name:      "priority",
						Value:     -1,
						UsageText: "runs the procedures with this priority, run all if is negative",
					},
				},
			},
			{
				Name:    "version",
				Usage:   "print the version",
				Aliases: []string{"v"},
				Action: func(_ context.Context, cmd *cli.Command) error {
					cli.VersionPrinter(cmd.Root())
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		utils.ErrorPrintf("%v\n", err)
		os.Exit(exitcode.Unknown)
	}

	os.Exit(exitcode.Success)
}
