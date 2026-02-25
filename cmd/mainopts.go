package cmd

type MainOpts struct {
	DDDir    string `short:"D" long:"dddir" description:"the dotdecldir path to operate"`
	Loglevel string `short:"l" long:"loglevel" default:"WARN" description:"Specify log level. One of DEBUG INFO WARN ERROR, case-insensitive"`
	Logfile  string `short:"L" long:"logfile" description:"Specify log file"`
	GlobConf string `short:"C" long:"config" description:"Specify global config"`
	Version  bool   `short:"V" long:"version" description:"Gives you version info, same as subcommand version"`
	// -h and --help is automatically set, since HelpFlag is set
}
