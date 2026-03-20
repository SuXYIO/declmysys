package subcmds

type MainOpts struct {
	DDir     string `short:"D" long:"decldir" description:"the decldir path to operate"`
	GlobConf string `short:"C" long:"config" description:"Specify global config"`
	Version  bool   `short:"V" long:"version" description:"Gives you version info, same as subcommand version"`
	Help     bool   `short:"h" long:"help" description:"Shows help, same as subcommand help"`
}
