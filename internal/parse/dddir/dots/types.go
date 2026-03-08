package dots

import "github.com/suxyio/declmysys/internal/parse/priority"

type Dots []Dot

type Dot struct {
	Description Desc
	Data        string // path of the data
}

type Desc struct {
	Name     string            `toml:"name"`
	Preset   string            `toml:"preset"`
	Priority priority.Priority `toml:"priority"`
	RunDat   map[string]any    `toml:"rundat"` // HACK: It works, but requires addtional safety checks manually
}
