package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"

	"github.com/suxyio/declmysys/cmd/subcmd"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
)

type MainOpts struct {
	DDDir    string `short:"D" description:"the dotdecldir path to operate"`
	Loglevel string `short:"l" long:"loglevel" default:"WARN" description:"Specify log level. One of DEBUG INFO WARN ERROR, case-insensitive"`
	Logfile  string `short:"L" long:"logfile" default:"" description:"Specify log file"`
	GlobConf string `short:"C" long:"config" default:"~/.config/declmysys/config.toml" description:"Specify global config"`
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

	if _, e := parser.AddCommand("do", "Execute defined stuff", "Execute defined stuff", argDo); e != nil {
		fmt.Fprintf(os.Stderr, "add subcommand do fail: %v\n", e)
		os.Exit(exitcode.SetupError)
	}
	if _, e := parser.AddCommand("init", "Initialize new dotdecldir", "Initialize new dotdecldir", argInit); e != nil {
		fmt.Fprintf(os.Stderr, "add subcommand init fail: %v\n", e)
		os.Exit(exitcode.SetupError)
	}
	if _, e := parser.AddCommand("list", "List the procedure defined by order", "List the procedure defined by order (descending priority)", argList); e != nil {
		fmt.Fprintf(os.Stderr, "add subcommand list fail: %v\n", e)
		os.Exit(exitcode.SetupError)
	}
	if _, e := parser.AddCommand("help", "Gives you the commandline help", "Gives you the commandline help, use -h or --help for advanced", argHelp); e != nil {
		fmt.Fprintf(os.Stderr, "add subcommand help fail: %v\n", e)
		os.Exit(exitcode.SetupError)
	}
	if _, e := parser.AddCommand("version", "Gives you the version info", "Gives you the version info, same as -V or --version", argVersion); e != nil {
		fmt.Fprintf(os.Stderr, "add subcommand version fail: %v\n", e)
		os.Exit(exitcode.SetupError)
	}

	_, err := parser.Parse()
	if err != nil {
		// PrintErrors flag is set so no need to print manually
		os.Exit(exitcode.InvalidArgs)
	}

	// version
	if argMain.Version {
		subcmd.Version()
		os.Exit(exitcode.Success)
	}

	//TODO: Setup logger here

	if parser.Command.Active == nil {
		fmt.Fprint(os.Stderr, "no subcommand specified\n")
		os.Exit(exitcode.InvalidArgs)
	}

	switch parser.Active.Name {
	case "help":
		subcmd.Help(parser)
	case "version":
		subcmd.Version()
	case "do":
		subcmd.Do()
	case "init":
		subcmd.Init()
	case "list":
		subcmd.List()
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %v\n", parser.Active.Name)
		os.Exit(exitcode.InvalidArgs)
	}

	os.Exit(exitcode.Success)
}
