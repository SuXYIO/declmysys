package decls

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/ddir/subs"
)

// Preset defines behaviors of a preset
type preset struct {
	IsValidFunc func(Decl, toml.MetaData) error         // return nil if isvalid
	RunFunc     func(Decl, cmdtype.CmdRunOptions) error // self-explanatory, ASSUMES that valid checks are ran before, also does subs
}

// Presets maps preset name to preset definition, shall not be modified
var presets = map[string]preset{
	"gitclone": {
		IsValidFunc: func(d Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			// url
			if !md.IsDefined("rundat", "url") {
				return fmt.Errorf("must specify rundat.url for preset gitclone")
			}
			if _, ok := d.RunDat["url"].(string); !ok {
				return fmt.Errorf("rundat.url must be of type string for preset gitclone")
			}
			// dest
			if !md.IsDefined("rundat", "dest") {
				return fmt.Errorf("must specify rundat.dest for preset gitclone")
			}
			if _, ok := d.RunDat["dest"].(string); !ok {
				return fmt.Errorf("rundat.dest must be of type string for preset gitclone")
			}
			return nil
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			if url, err := subs.ApplyG(d.RunDat["url"].(string)); err != nil {
				return err
			} else {
				d.RunDat["url"] = url
			}
			if dest, err := subs.ApplyPC(d.RunDat["dest"].(string)); err != nil {
				return err
			} else {
				d.RunDat["dest"] = dest
			}
			//TODO: impl
			return nil
		},
	},

	"stow": {
		IsValidFunc: func(d Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			if !md.IsDefined("rundat", "datadir") {
				d.RunDat["datadir"] = consts.DefaultDeclsDataDir
			} else {
				if _, ok := d.RunDat["datadir"].(string); !ok {
					return fmt.Errorf("rundat.datadir must be of type string for preset stow")
				}
			}
			return nil
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			if datadir, err := subs.ApplyPC(d.RunDat["datadir"].(string)); err != nil {
				return err
			} else {
				d.RunDat["datadir"] = datadir
			}
			//TODO: impl
			return nil
		},
	},

	"cmds": {
		IsValidFunc: func(d Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			if !md.IsDefined("rundat", "cmds") {
				return fmt.Errorf("must specify rundat.cmds for preset cmds")
			}

			if cmds, err := convertCmds(d.RunDat["cmds"]); err != nil {
				return err
			} else {
				d.RunDat["cmds"] = cmds
			}
			return nil
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// turn to cmd
			cmds, err := convertCmds(d.RunDat["cmds"])
			if err != nil {
				return err
			}

			// run
			for _, cmd := range cmds {
				cmd.Run(opts)
			}
			return nil
		},
	},
}

// convert rundat.cmds from any to []cmdtype.Cmd
func convertCmds(incmds any) ([]cmdtype.Cmd, error) {
	// and watch these error msgs, if you don't know whats "verbose" lmao
	if cmds, ok := incmds.([]cmdtype.Cmd); ok {
		return cmds, nil
	}

	var res []cmdtype.Cmd

	// first layer: any slice
	slice1, ok := incmds.([]any)
	if !ok {
		return nil, fmt.Errorf("rundat.cmds must be of type []any for preset cmds, got cmds %v of type %T", incmds, incmds)
	}
	for _, v := range slice1 {
		cmd, err := convertCmd(v)
		if err != nil {
			return nil, err
		}

		res = append(res, cmd)
	}

	return res, nil
}

func convertCmd(v any) (cmdtype.Cmd, error) {
	if cmd, ok := v.(cmdtype.Cmd); ok {
		return cmd, nil
	}

	// second layer: slice of any slice
	slice2, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("rundat.cmds[i] must be of type []any for preset cmds, got %v of type %T", v, v)
	}

	var cmd cmdtype.Cmd
	for _, t := range slice2 {
		// third layer: slice of string slice
		cmdElem, ok := t.(string)
		if !ok {
			return nil, fmt.Errorf("rundat.cmds[i][j] must be of type string for preset cmds, got %v of type %T", t, t)
		}
		cmd = append(cmd, cmdElem)
	}

	return cmd, nil
}
