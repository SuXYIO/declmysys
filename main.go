package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/globconf"
	"github.com/suxyio/declmysys/internal/parse/subs"
	"github.com/suxyio/declmysys/internal/subcmds"
	"github.com/suxyio/declmysys/internal/utils"
)

func main() {
	// parser setup
	argMain := &subcmds.MainOpts{}
	// HACK: Inconsistent help with no autopaging (via HelpFlag)
	parser := flags.NewParser(argMain, flags.PrintErrors|flags.HelpFlag)
	parser.Name = consts.Name
	parser.ShortDescription = consts.Desc
	parser.SubcommandsOptional = true

	argRun := &subcmds.RunOpts{}
	argInit := &subcmds.InitOpts{}
	argList := &subcmds.ListOpts{}
	argHelp := &subcmds.HelpOpts{}
	argVersion := &subcmds.VersionOpts{}

	if _, err := parser.AddCommand("run", "Execute defined stuff", "Execute defined stuff", argRun); err != nil {
		utils.Panic("add subcommand 'run' fail", err, exitcode.SetupError)
	}
	if _, err := parser.AddCommand("init", "Initialize new decldir", "Initialize new decldir", argInit); err != nil {
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

	if _, err := parser.Parse(); err != nil {
		// PrintErrors flag is set so no need to print manually
		os.Exit(exitcode.InvalidArgs)
	}

	// version (-V)
	if argMain.Version {
		subcmds.Version(*argVersion)
		os.Exit(exitcode.Success)
	}

	// if not using -v or -h, and no subcommand given
	if parser.Command.Active == nil {
		subcmds.Help(*parser, *argHelp)
		utils.Panic("no subcommand specified", nil, exitcode.InvalidArgs)
	}

	// parse argmain.globconf (default case)
	if option := parser.FindOptionByLongName("config"); option != nil && option.IsSetDefault() {
		argMain.GlobConf = consts.DefaultGlobconfPath
	}
	// get globconf
	gc, err := globconf.GetGlobconf(argMain.GlobConf)
	if err != nil {
		utils.Panic("read/create global config fail", err, exitcode.ConfigError)
	}
	// replace argmain ddir with default in globconf if empty
	if option := parser.FindOptionByLongName("decldir"); option != nil && option.IsSetDefault() {
		if ddir, err := subs.ApplyDefaultPC(gc.DDir); err != nil {
			utils.Panic("apply default paths&cmds subs to main arg ddir fail", err, exitcode.Unknown)
		} else {
			argMain.DDir = ddir
		}
	}

	switch parser.Active.Name {
	// not gonna design error capture for subcommands, just panic in the subcmd function if anything goes wrong
	case "help":
		subcmds.Help(*parser, *argHelp)
	case "version":
		subcmds.Version(*argVersion)
	case "run":
		subcmds.Run(gc, *argMain, *argRun)
	case "init":
		subcmds.Init(gc, *argMain, *argInit)
	case "list":
		subcmds.List(gc, *argMain, *argList)
	default:
		utils.Panic(fmt.Sprintf("unknown subcommand: %v", parser.Active.Name), nil, exitcode.InvalidArgs)
	}

	os.Exit(exitcode.Success)
}
