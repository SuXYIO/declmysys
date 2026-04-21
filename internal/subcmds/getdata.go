package subcmds

import (
	"fmt"
	"os"

	"github.com/suxyio/declmysys/internal/exitcode"
	"github.com/suxyio/declmysys/internal/parse/decls"
	"github.com/suxyio/declmysys/internal/parse/metadata"
	"github.com/suxyio/declmysys/internal/utils"
)

// getDeclsData gets data `run` and `list` needs.
// dunno how to name it, just packing stuff that was wrote two times.
// it panics directly btw
func getDeclsData(ddir string) decls.Decls {
	// ddir doesn't exist
	if !DDirExist(ddir) {
		utils.Panic(fmt.Sprintf("ddir %s does not exist, you can create via \"init\" subcommand", ddir), nil, exitcode.FileError)
	}

	// metadata.toml
	if err := metadata.GetGlobalMetadata(ddir); err != nil {
		utils.Panic("error getting metadata.toml", err, exitcode.ConfigError)
	}

	// check version
	version, err := getVer()
	if err != nil {
		utils.Panic("error getting version", err, exitcode.Unknown)
	}
	checkVer(version, metadata.GlobalMetaData.Version)

	// decls
	// can't use "decls" as var name since decls package took it damn it
	declss, err := decls.GetDecls(ddir)
	if err != nil {
		utils.Panic("error getting decls", err, exitcode.ConfigError)
	}
	return declss
}

// checkVer checks if current ver and ddir ver matches, and asks the user whether to proceed if not matching,
// 'll just exit if user chooses not to, and returns if do proceed
func checkVer(ver string, ddver string) {
	if ver == ddver {
		// matching
		return
	}

	// not matching
	utils.WarnPrintf("warning: program version %q and decldir metadata version %q does not match,\ndirectly operating might be unsafe, using another version of the program or manual upgrade for the decldir is recommended\n", ver, ddver)
	if utils.AskYN("proceed with operation?") {
		return
	} else {
		os.Exit(exitcode.Success)
	}
}
