package decls

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/ddir/substoml"
)

// Preset behavior definition
type Preset struct {
	DescIsValidFunc func(Desc, toml.MetaData) error // return nil if isvalid
	DescSubsFunc    func(*Desc) error               // do subs, ASSUMES that valid check is ran before
	RunFunc         func(Decl) error                // self-explanatory, ASSUMES that valid check & subs are ran before
}

// Maps preset name to preset definition, shall not be modified
var Presets = map[string]Preset{
	"gitclone": {
		DescIsValidFunc: func(d Desc, md toml.MetaData) error {
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
		DescSubsFunc: func(d *Desc) error {
			tmp, err := substoml.ApplyG(d.RunDat["url"].(string))
			if err != nil {
				return err
			}
			d.RunDat["url"] = tmp
			tmp, err = substoml.ApplyPC(d.RunDat["dest"].(string))
			if err != nil {
				return err
			}
			d.RunDat["dest"] = tmp

			return nil
		},
		RunFunc: func(d Decl) error {
			//TODO: impl
			return nil
		},
	},

	"stow": {
		DescIsValidFunc: func(d Desc, md toml.MetaData) error {
			if !md.IsDefined("rundat", "datadir") {
				d.RunDat["datadir"] = consts.DefaultDeclsDataDir
			} else {
				if _, ok := d.RunDat["datadir"].(string); !ok {
					return fmt.Errorf("rundat.datadir must be of type string for preset stow")
				}
			}
			return nil
		},
		DescSubsFunc: func(d *Desc) error {
			tmp, err := substoml.ApplyPC(d.RunDat["datadir"].(string))
			if err != nil {
				return err
			}
			d.RunDat["datadir"] = tmp
			return nil
		},
		RunFunc: func(d Decl) error {
			//TODO: impl
			return nil
		},
	},

	"cmds": {
		DescIsValidFunc: func(d Desc, md toml.MetaData) error {
			if !md.IsDefined("rundat", "cmds") {
				return fmt.Errorf("must specify rundat.cmds for preset cmds")
			}
			tmp, err := convertCmds(d.RunDat["cmds"])
			if err != nil {
				return err
			}
			d.RunDat["cmds"] = tmp
			return nil
		},
		DescSubsFunc: func(d *Desc) error {
			for _, v := range d.RunDat["cmds"].([]cmdtype.Cmd) {
				for j := range v {
					tmp, err := substoml.ApplyPC(v[j])
					if err != nil {
						return err
					}
					v[j] = tmp
				}
			}
			return nil
		},
		RunFunc: func(d Decl) error {
			//TODO: impl
			return nil
		},
	},
}

// convert rundat.cmds from any to []cmdtype.Cmd
func convertCmds(incmds any) ([]cmdtype.Cmd, error) {
	// and watch these error msgs, if you don't know whats "verbose" lmao
	var res []cmdtype.Cmd

	// first layer: any slice
	slice1, ok := incmds.([]any)
	if !ok {
		return nil, fmt.Errorf("rundat.cmds must be of type []any for preset cmds, got cmds %v of type %T", incmds, incmds)
	}
	for _, v := range slice1 {
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

		res = append(res, cmd)
	}

	return res, nil
}
