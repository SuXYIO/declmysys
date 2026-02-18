package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"

	"github.com/suxyio/declmysys/cmd/subcmd"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/funcs"
	"github.com/suxyio/declmysys/internal/parse"
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

//go:embed configs/config.toml
var defaultGlobconfFS embed.FS

func main() {
	// TODO: use funcs.Panic
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

	// if not using -v or -h, and no subcommand given
	if parser.Command.Active == nil {
		fmt.Fprint(os.Stderr, "no subcommand specified\n")
		os.Exit(exitcode.InvalidArgs)
	}

	// read globconf
	gc, err := GetGlobconf(argMain.DDDir)
	if err != nil {
		funcs.Panic("read/create global config fail", err, exitcode.ConfigError)
	}

	switch parser.Active.Name {
	case "help":
		subcmd.Help(parser)
	case "version":
		subcmd.Version()
	case "do":
		subcmd.Do(gc, argDo)
	case "init":
		subcmd.Init(gc, argInit)
	case "list":
		subcmd.List(gc, argList)
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %v\n", parser.Active.Name)
		os.Exit(exitcode.InvalidArgs)
	}

	os.Exit(exitcode.Success)
}

func GetGlobconf(gcpath string) (parse.Globconf, error) {
	// Get global config, create one if not exist
	var gc parse.Globconf
	if _, err := os.Stat(gcpath); err == nil {
		// exists, no action
	} else if !os.IsNotExist(err) {
		// other error
		return gc, err
	} else {
		// not exist, create one
		// TODO: add asking here (y/n)
		fmt.Printf("Global config file not found at %s, creating default\n", gcpath)

		// create dir
		dir := filepath.Dir(gcpath)
		if dir != "." && dir != "" {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				return gc, err
			}
		}

		// copy file

		defaultGlobconfDat, err := defaultGlobconfFS.ReadFile("configs/config.toml")
		if err != nil {
			return gc, err
		}
		err = os.WriteFile(gcpath, defaultGlobconfDat, 0644)
		if err != nil {
			return gc, err
		}
	}
	// read it
	dat, err := os.ReadFile(gcpath)
	if err != nil {
		return gc, err
	}
	gc, err = parse.ParseGlobconf(dat)
	if err != nil {
		return gc, err
	}

	return gc, nil
}
