package decls

import "github.com/suxyio/declmysys/internal/parse/priority"

// Decl is the info for a single declaration
type Decl struct {
	Name     string            `toml:"name"`
	Preset   string            `toml:"preset"`
	Priority priority.Priority `toml:"priority"`
	RunDat   map[string]any    `toml:"rundat"` // HACK: It works, but requires addtional safety checks manually
}
