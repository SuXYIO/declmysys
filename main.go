package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/suxyio/declmysys/cmd"
	"github.com/suxyio/declmysys/cmd/subcmd"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/subs"
	"github.com/suxyio/declmysys/internal/utils"
)

func main() {
	// parser setup
	argMain := &cmd.MainOpts{}
	parser := flags.NewParser(argMain, flags.PrintErrors|flags.HelpFlag)
	parser.Name = consts.Name
	parser.ShortDescription = consts.Desc
	parser.SubcommandsOptional = true

	argDo := &subcmd.DoOpts{}
	argInit := &subcmd.InitOpts{}
	argList := &subcmd.ListOpts{}
	argHelp := &subcmd.HelpOpts{}
	argVersion := &subcmd.VersionOpts{}

	if _, err := parser.AddCommand("do", "Execute defined stuff", "Execute defined stuff", argDo); err != nil {
		utils.Panic("add subcommand 'do' fail", err, exitcode.SetupError)
	}
	if _, err := parser.AddCommand("init", "Initialize new dotdecldir", "Initialize new dotdecldir", argInit); err != nil {
		utils.Panic("add subcommand 'init' fail", err, exitcode.SetupError)
	}
	if _, err := parser.AddCommand("list", "List the procedure defined by order", "List the procedure defined by order (descending priority)", argList); err != nil {
		utils.Panic("add subcommand 'list' fail", err, exitcode.SetupError)
	}
	if _, err := parser.AddCommand("help", "Gives you the commandline help", "Gives you the commandline help, use -h or --help for advanced", argHelp); err != nil {
		utils.Panic("add subcommand 'help' fail", err, exitcode.SetupError)
	}
	if _, err := parser.AddCommand("version", "Gives you the version info", "Gives you the version info, same as -V or --version", argVersion); err != nil {
		utils.Panic("add subcommand 'version' fail", err, exitcode.SetupError)
	}

	_, err := parser.Parse()
	if err != nil {
		// PrintErrors flag is set so no need to print manually
		os.Exit(exitcode.InvalidArgs)
	}

	// version
	if argMain.Version {
		subcmd.Version(argVersion)
		os.Exit(exitcode.Success)
	}

	//TODO: Setup logger here

	// if not using -v or -h, and no subcommand given
	if parser.Command.Active == nil {
		parser.WriteHelp(os.Stdout)
		fmt.Println()
		utils.Panic("no subcommand specified", nil, exitcode.InvalidArgs)
	}

	// parse argmain.globconf (default case)
	if argMain.GlobConf == "" {
		argMain.GlobConf, err = consts.DefaultGlobconfPath()
		if err != nil {
			utils.Panic("get default globconf path fail", err, exitcode.Unknown)
		}
	}
	// get globconf
	gc, err := utils.GetGlobconf(argMain.GlobConf)
	if err != nil {
		utils.Panic("read/create global config fail", err, exitcode.ConfigError)
	}
	// replace argmain dddir with default in globconf if empty
	if argMain.DDDir == "" {
		argMain.DDDir, err = subs.ApplyDefaultPC(gc.DDDir)
		if err != nil {
			utils.Panic("apply default paths&cmds subs to main arg dddir fail", err, exitcode.Unknown)
		}
	}

	switch parser.Active.Name {
	// not gonna design error capture for subcommands, just panic in the subcmd function if anything goes wrong
	case "help":
		subcmd.Help(argHelp, parser)
	case "version":
		subcmd.Version(argVersion)
	case "do":
		subcmd.Do(gc, argMain, argDo)
	case "init":
		subcmd.Init(gc, argMain, argInit)
	case "list":
		subcmd.List(gc, argMain, argList)
	default:
		utils.Panic(fmt.Sprintf("unknown subcommand: %v", parser.Active.Name), nil, exitcode.InvalidArgs)
	}

	os.Exit(exitcode.Success)
}
