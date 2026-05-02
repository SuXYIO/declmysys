package decls

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/utils"
)

const (
	ToStringModeList = iota
	ToStringModeRun
)

// Preset defines behaviors of a preset
type preset struct {
	PreFunc  func(*Decl, toml.MetaData) error        // does the validation and parsing work
	ToString func(Decl) string                       // turns specific attributes into readable string, ASSUMES that prefunc is ran before
	RunFunc  func(Decl, cmdtype.CmdRunOptions) error // self-explanatory, ASSUMES that prefunc is ran before
}

// Presets maps preset name to preset definition, shall not be modified
var presets = map[string]preset{
	"packages": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			// manager
			if !md.IsDefined("args", "manager") {
				return fmt.Errorf("must specify args.manager for preset \"packages\"")
			}
			var man manSpec
			if err := man.toManspec(d.Args["manager"]); err != nil {
				return err
			}
			d.Args["manager"] = man

			// packs
			if !md.IsDefined("args", "packs") {
				return fmt.Errorf("must specify args.packs for preset \"packages\"")
			}

			return nil
		},
		ToString: func(d Decl) string {
			return fmt.Sprintf(": %v", d.Args["packs"])
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// manager
			man, ok := d.Args["manager"].(manSpec)
			if !ok {
				return fmt.Errorf("args.manager must be of type manSpec for preset \"packages\" (in run, maybe forgot to run PreFunc before RunFunc? %s)", consts.NotYourFault)
			}
			var mancmd []string
			if man.Preset != "" {
				// preset
				if cmd, exists := manPresets[man.Preset]; !exists {
					return fmt.Errorf("manager preset not found for preset name %q", man.Preset)
				} else {
					mancmd = cmd
				}
			} else {
				mancmd = man.CustomCmd
			}

			// packs
			packany, ok := d.Args["packs"].([]any)
			if !ok {
				return fmt.Errorf("packs must be of type []any for preset \"packages\", got %v of type %T", d.Args["packs"], d.Args["packs"])
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
			// src
			if !md.IsDefined("args", "src") {
				return fmt.Errorf("must specify args.src for preset \"gitclone\"")
			}
			if _, ok := d.Args["src"].(string); !ok {
				return fmt.Errorf("args.src must be of type string for preset \"gitclone\", got %v of type %T", d.Args["src"], d.Args["src"])
			}
			// dest
			if !md.IsDefined("args", "dest") {
				return fmt.Errorf("must specify args.dest for preset \"gitclone\"")
			}
			if _, ok := d.Args["dest"].(string); !ok {
				return fmt.Errorf("args.dest must be of type string for preset \"gitclone\", got %v of type %T", d.Args["dest"], d.Args["dest"])
			}
			return nil
		},
		ToString: func(d Decl) string {
			return fmt.Sprintf(": %s => %s", d.Args["src"], d.Args["dest"])
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			cmd := cmdtype.Cmd{"git", "clone", d.Args["src"].(string), d.Args["dest"].(string)}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"stow": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			// src
			if !md.IsDefined("args", "src") {
				d.Args["src"] = "stow"
			} else {
				if _, ok := d.Args["src"].(string); !ok {
					return fmt.Errorf("args.src must be of type string for preset \"stow\", got %v of type %T", d.Args["src"], d.Args["src"])
				}
			}
			// dest
			if !md.IsDefined("args", "dest") {
				d.Args["dest"] = "{HOME}"
			} else {
				if _, ok := d.Args["dest"].(string); !ok {
					return fmt.Errorf("args.dest must be of type string for preset \"stow\", got %v of type %T", d.Args["dest"], d.Args["dest"])
				}
			}

			return nil
		},
		ToString: func(d Decl) string {
			return fmt.Sprintf(": %s => %s", d.Args["src"], d.Args["dest"])
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			cmd := cmdtype.Cmd{"stow", "-t", d.Args["dest"].(string), d.Args["src"].(string)}
			if err := cmd.Run(opts); err != nil {
				return err
			}
			return nil
		},
	},

	"cmds": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if !md.IsDefined("args", "cmds") {
				return fmt.Errorf("must specify args.cmds for preset \"cmds\"")
			}

			if cmds, err := cmdtype.ToCmds(d.Args["cmds"]); err != nil {
				return err
			} else {
				d.Args["cmds"] = cmds
			}
			return nil
		},
		ToString: func(d Decl) string {
			return fmt.Sprintf(": %v", d.Args["cmds"])
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			// turn to cmd
			cmds, err := cmdtype.ToCmds(d.Args["cmds"])
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

	"manual": {
		PreFunc: func(d *Decl, md toml.MetaData) error {
			if !md.IsDefined("args", "desc") {
				return fmt.Errorf("must specify args.desc for preset \"manual\"")
			}
			if desc, exists := d.Args["desc"]; !exists {
				return fmt.Errorf("must specify args.desc for preset \"manual\"")
			} else {
				if _, ok := desc.(string); !ok {
					return fmt.Errorf("args.desc must be of string type for preset \"manual\"")
				}
			}
			return nil
		},
		ToString: func(d Decl) string {
			return fmt.Sprintf(": %q", d.Args["desc"])
		},
		RunFunc: func(d Decl, opts cmdtype.CmdRunOptions) error {
			desc, ok := d.Args["desc"].(string)
			if !ok {
				return fmt.Errorf("args.desc must be of string type for preset \"manual\"")
			}

			if opts.DryRun != nil {
				_, err := opts.DryRun.Write([]byte(desc))
				return err
			}

			fmt.Println(desc)

			fmt.Print("press any key to continue...")
			if _, err := utils.GetAnyKey(); err != nil {
				return fmt.Errorf("error getting any key: %v", err)
			}

			return nil
		},
	},
}
