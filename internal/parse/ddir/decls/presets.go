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
	PreFunc func(Decl, toml.MetaData) error         // does the validation and parsing work
	RunFunc func(Decl, cmdtype.CmdRunOptions) error // self-explanatory, ASSUMES that prefunc is ran before, also does subs
}

// Presets maps preset name to preset definition, shall not be modified
var presets = map[string]preset{
	"packages": {
		PreFunc: func(d Decl, md toml.MetaData) error {
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
			packs, ok := d.RunDat["packs"].([]string)
			if !ok {
				return fmt.Errorf("packs must be of type []string for preset \"packages\", got %v of type %T", d.RunDat["packs"], d.RunDat["packs"])
			}

			// override opts
			opts.AppendedArgs = packs
			opts.DoPCSubs = false

			// run
			cmd := cmdtype.Cmd(mancmd)
			if err := cmd.Run(opts); err != nil {
				return err
			}

			return nil
		},
	},

	"gitclone": {
		PreFunc: func(d Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			// url
			if !md.IsDefined("rundat", "url") {
				return fmt.Errorf("must specify rundat.url for preset \"gitclone\"")
			}
			if _, ok := d.RunDat["url"].(string); !ok {
				return fmt.Errorf("rundat.url must be of type string for preset \"gitclone\", got %v of type %T", d.RunDat["url"], d.RunDat["url"])
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
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			url, err := subs.ApplyG(d.RunDat["url"].(string))
			if err != nil {
				return err
			}
			dest, err := subs.ApplyPC(d.RunDat["dest"].(string))
			if err != nil {
				return err
			}
			opts.DoPCSubs = false

			// run
			cmd := cmdtype.Cmd{"git", "clone", url, dest}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"stow": {
		PreFunc: func(d Decl, md toml.MetaData) error {
			if d.RunDat == nil {
				d.RunDat = make(map[string]any)
			}

			if !md.IsDefined("rundat", "datadir") {
				d.RunDat["datadir"] = consts.DefaultDeclsDataDir
			} else {
				if _, ok := d.RunDat["datadir"].(string); !ok {
					return fmt.Errorf("rundat.datadir must be of type string for preset \"stow\", got %v of type %T", d.RunDat["datadir"], d.RunDat["datadir"])
				}
			}
			return nil
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// subs
			datadir, err := subs.ApplyPC(d.RunDat["datadir"].(string))
			if err != nil {
				return err
			}
			opts.DoPCSubs = false

			// run
			cmd := cmdtype.Cmd{"stow", datadir}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"cmds": {
		PreFunc: func(d Decl, md toml.MetaData) error {
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
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// turn to cmd
			cmds, err := cmdtype.ToCmds(d.RunDat["cmds"])
			if err != nil {
				return err
			}
			opts.DoPCSubs = true

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
