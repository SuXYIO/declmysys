package decls

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/subs"
)

const (
	ToStringModeList int8 = iota
	ToStringModeRun
)

// Preset defines behaviors of a preset
type preset struct {
	PreFunc  func(*Decl, toml.MetaData) error        // does the validation and parsing work
	ToString func(Decl, int8, string) string         // turns into readable string, ASSUMES that prefunc is ran before
	RunFunc  func(Decl, cmdtype.CmdRunOptions) error // self-explanatory, ASSUMES that prefunc is ran before, also does subs
}

//TODO: Change slice-printing work to "%v", or define cmdtype.Cmd to string

// Presets maps preset name to preset definition, shall not be modified
var presets = map[string]preset{
	"packages": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			// manager
			if !md.IsDefined("rundat", "manager") {
				return fmt.Errorf("must specify rundat.manager for preset \"packages\"")
			}
			var man manSpec
			if err := man.toManspec(d.RunDat["manager"]); err != nil {
				return err
			}
			d.RunDat["manager"] = man

			// packs
			if !md.IsDefined("rundat", "packs") {
				return fmt.Errorf("must specify rundat.packs for preset \"packages\"")
			}

			return nil
		},
		ToString: func(d Decl, mode int8, prestr string) string {
			// indent is the level of indents
			switch mode {
			case ToStringModeRun:
				return fmt.Sprintf("%spackages[%s]", prestr, d.Name)
			default:
				// ToStringModeList
				return fmt.Sprintf("%spackages[%s]:\t%v", prestr, d.Name, d.RunDat["packs"])
			}
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			// manager
			man, ok := d.RunDat["manager"].(manSpec)
			if !ok {
				return fmt.Errorf("rundat.manager must be of type manSpec for preset \"packages\" (in run, maybe forgot to run PreFunc before RunFunc? %s)", consts.NotYourFault)
			}
			var rawcmd []string
			if man.Preset != "" {
				// preset
				if cmd, exists := manPresets[man.Preset]; !exists {
					return fmt.Errorf("manager preset not found for preset name %q", man.Preset)
				} else {
					rawcmd = cmd
				}
			} else {
				rawcmd = man.CustomCmd
			}
			var mancmd []string
			for _, elem := range rawcmd {
				tmp, err := subs.ApplyPC(elem)
				if err != nil {
					return err
				} else {
					mancmd = append(mancmd, tmp)
				}
			}
			// no subs for packs
			packany, ok := d.RunDat["packs"].([]any)
			if !ok {
				return fmt.Errorf("packs must be of type []any for preset \"packages\", got %v of type %T", d.RunDat["packs"], d.RunDat["packs"])
			}

			var packs []string
			for i := range packany {
				pack, ok := packany[i].(string)
				if !ok {
					return fmt.Errorf("pack in packs must be of type string, got %v of type %T at index %d", packany[i], packany[i], i)
				}
				packs = append(packs, pack)
			}

			// override opts
			opts.AppendedArgs = packs

			// run
			cmd := cmdtype.Cmd(mancmd)
			if err := cmd.Run(opts); err != nil {
				return err
			}

			return nil
		},
	},

	"gitclone": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			// src
			if !md.IsDefined("rundat", "src") {
				return fmt.Errorf("must specify rundat.src for preset \"gitclone\"")
			}
			if _, ok := d.RunDat["src"].(string); !ok {
				return fmt.Errorf("rundat.src must be of type string for preset \"gitclone\", got %v of type %T", d.RunDat["src"], d.RunDat["src"])
			}
			// dest
			if !md.IsDefined("rundat", "dest") {
				return fmt.Errorf("must specify rundat.dest for preset \"gitclone\"")
			}
			if _, ok := d.RunDat["dest"].(string); !ok {
				return fmt.Errorf("rundat.dest must be of type string for preset \"gitclone\", got %v of type %T", d.RunDat["dest"], d.RunDat["dest"])
			}
			return nil
		},
		ToString: func(d Decl, mode int8, prestr string) string {
			switch mode {
			case ToStringModeRun:
				return fmt.Sprintf("%sgitclone[%s]", prestr, d.Name)
			default:
				// ToStringModeList
				return fmt.Sprintf("%sgitclone[%s]:\tsrc %q, dest %q", prestr, d.Name, d.RunDat["src"], d.RunDat["dest"])
			}
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			src, err := subs.ApplyG(d.RunDat["src"].(string))
			if err != nil {
				return err
			}
			dest, err := subs.ApplyPC(d.RunDat["dest"].(string))
			if err != nil {
				return err
			}

			// run
			cmd := cmdtype.Cmd{"git", "clone", src, dest}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"stow": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			// src
			if !md.IsDefined("rundat", "src") {
				d.RunDat["src"] = "stow"
			} else {
				if _, ok := d.RunDat["src"].(string); !ok {
					return fmt.Errorf("rundat.src must be of type string for preset \"stow\", got %v of type %T", d.RunDat["src"], d.RunDat["src"])
				}
			}
			// dest
			if !md.IsDefined("rundat", "dest") {
				d.RunDat["dest"] = "{HOME}"
			} else {
				if _, ok := d.RunDat["dest"].(string); !ok {
					return fmt.Errorf("rundat.dest must be of type string for preset \"stow\", got %v of type %T", d.RunDat["dest"], d.RunDat["dest"])
				}
			}

			return nil
		},
		ToString: func(d Decl, mode int8, prestr string) string {
			switch mode {
			case ToStringModeRun:
				return fmt.Sprintf("%sstow[%s]", prestr, d.Name)
			default:
				// ToStringModeList
				return fmt.Sprintf("%sstow[%s]:\tsrc %q, dest %q", prestr, d.Name, d.RunDat["src"], d.RunDat["dest"])
			}
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// run
			cmd := cmdtype.Cmd{"stow", "-t", d.RunDat["dest"].(string), d.RunDat["src"].(string)}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"cmds": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			if !md.IsDefined("rundat", "cmds") {
				return fmt.Errorf("must specify rundat.cmds for preset \"cmds\"")
			}

			if cmds, err := cmdtype.ToCmds(d.RunDat["cmds"]); err != nil {
				return err
			} else {
				d.RunDat["cmds"] = cmds
			}
			return nil
		},
		ToString: func(d Decl, mode int8, prestr string) string {
			switch mode {
			case ToStringModeRun:
				return fmt.Sprintf("%scmds[%s]", prestr, d.Name)
			default:
				// ToStringModeList
				return fmt.Sprintf("%scmds[%s]", prestr, d.Name)
			}
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// turn to cmd
			cmds, err := cmdtype.ToCmds(d.RunDat["cmds"])
			if err != nil {
				return err
			}

			// run
			for _, cmd := range cmds {
				if err := cmd.Run(opts); err != nil {
					return err
				}
			}
			return nil
		},
	},
}
