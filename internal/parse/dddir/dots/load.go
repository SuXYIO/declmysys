package dots

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/suxyio/declmysys/internal/consts"
	"github.com/suxyio/declmysys/internal/parse/cmdtype"
	"github.com/suxyio/declmysys/internal/parse/dddir/substoml"
)

func (dots *Dots) Load(dddir string) error {
	dotsdir := filepath.Join(dddir, "dots")

	dotsEnt, err := os.ReadDir(dotsdir)
	if err != nil {
		return err
	}

	for _, ent := range dotsEnt {
		if ent.IsDir() {
			var dot Dot
			err := dot.Load(filepath.Join(dotsdir, ent.Name()))
			if err != nil {
				return err
			}
			*dots = append(*dots, dot)
		}
	}

	return nil
}

func (dot *Dot) Load(path string) error {
	// desc.toml
	descdata, err := os.ReadFile(filepath.Join(path, "desc.toml"))
	if err != nil {
		return err
	}
	err = dot.Description.Load(descdata)
	if err != nil {
		return err
	}

	// data/
	dot.Data = filepath.Join(path, "data")

	return nil
}

func (desc *Desc) Load(data []byte) error {
	metadat, err := toml.Decode(string(data), desc)
	if err != nil {
		return err
	}

	// check default fields
	if !metadat.IsDefined("name") {
		return fmt.Errorf("must specify name for dot desc")
	}
	if !metadat.IsDefined("preset") {
		return fmt.Errorf("must specify preset for dot desc")
	}
	if !metadat.IsDefined("priority") {
		desc.Priority = consts.DefaultDotsPriority
	}
	// check custom rundat fields for presets
	if desc.RunDat == nil {
		desc.RunDat = make(map[string]any)
	}
	switch desc.Preset {
	case "stow":
		// stow
		if !metadat.IsDefined("rundat", "datadir") {
			desc.RunDat["datadir"] = consts.DefaultDotsDataDir
		}

	case "gitclone":
		// url
		if !metadat.IsDefined("rundat", "url") {
			return fmt.Errorf("must specify rundat.url for preset gitclone")
		}
		if _, ok := desc.RunDat["url"].(string); !ok {
			return fmt.Errorf("rundat.url must be of type string for preset gitclone")
		}
		// dest
		if !metadat.IsDefined("rundat", "dest") {
			return fmt.Errorf("must specify rundat.dest for preset gitclone")
		}
		if _, ok := desc.RunDat["dest"].(string); !ok {
			return fmt.Errorf("rundat.dest must be of type string for preset gitclone")
		}

	case "cmds":
		// cmds
		if !metadat.IsDefined("rundat", "cmds") {
			return fmt.Errorf("must specify rundat.cmds for preset cmds")
		}
		desc.RunDat["cmds"], err = convertCmds(desc.RunDat["cmds"])
		if err != nil {
			return err
		}
	}

	// subs
	err = desc.subs()
	if err != nil {
		return err
	}

	return nil
}

func (desc *Desc) subs() error {
	// MUST run this after checking rundat contents, not gonna handle type assertion error again
	switch desc.Preset {
	case "stow":
		tmp, err := substoml.ApplyPC(desc.RunDat["datadir"].(string))
		if err != nil {
			return err
		}
		desc.RunDat["datadir"] = tmp
	case "gitclone":
		tmp, err := substoml.ApplyG(desc.RunDat["url"].(string))
		if err != nil {
			return err
		}
		desc.RunDat["url"] = tmp
		tmp, err = substoml.ApplyPC(desc.RunDat["dest"].(string))
		if err != nil {
			return err
		}
		desc.RunDat["dest"] = tmp
	case "cmds":
		for _, v := range desc.RunDat["cmds"].([]cmdtype.Cmd) {
			for j := range v {
				tmp, err := substoml.ApplyPC(v[j])
				if err != nil {
					return err
				}
				v[j] = tmp
			}
		}
	}
	return nil
}

// convert rundat.cmds from any to []cmdtype.Cmd
func convertCmds(incmds any) ([]cmdtype.Cmd, error) {
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
