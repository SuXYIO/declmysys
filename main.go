package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/suxyio/declmysys/cmd/subcmd"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/utils"
)

type MainOpts struct {
	// since cant embed the const string in the tag, pls ensure DDDir default == consts.DefaultDDDirPath
	DDDir    string `short:"D" default:"~/Dotdecl" description:"the dotdecldir path to operate"`
	Loglevel string `short:"l" long:"loglevel" default:"WARN" description:"Specify log level. One of DEBUG INFO WARN ERROR, case-insensitive"`
	Logfile  string `short:"L" long:"logfile" default:"" description:"Specify log file"`
	// also pls ensure GlobConf default == consts.DefaultGlobconfPath
	GlobConf string `short:"C" long:"config" default:"{CONF}/declmysys/config.toml" description:"Specify global config"`
	Version  bool   `short:"V" long:"version" description:"Gives you version info, same as subcommand version"`
	// -h and --help is automatically set, since HelpFlag is set
}

func main() {
	// parser setup
	argMain := MainOpts{}
	parser := flags.NewParser(&argMain, flags.PrintErrors|flags.HelpFlag)
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
		utils.Panic("no subcommand specified", nil, exitcode.InvalidArgs)
	}

	// get globconf
	gc, err := utils.GetGlobconf(argMain.GlobConf)
	if err != nil {
		utils.Panic("read/create global config fail", err, exitcode.ConfigError)
	}

	switch parser.Active.Name {
	case "help":
		subcmd.Help(argHelp, parser)
	case "version":
		subcmd.Version(argVersion)
	case "do":
		subcmd.Do(gc, argDo)
	case "init":
		subcmd.Init(gc, argInit)
	case "list":
		subcmd.List(gc, argList)
	default:
		utils.Panic(fmt.Sprintf("unknown subcommand: %v", parser.Active.Name), nil, exitcode.InvalidArgs)
	}

	os.Exit(exitcode.Success)
}
