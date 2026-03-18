package decls

// Decl is the info for a single declaration
type Decl struct {
	Name     string         `toml:"name"`
	Preset   string         `toml:"preset"`
	Priority Priority       `toml:"priority"`
	RunDat   map[string]any `toml:"rundat"` // HACK: It works, but requires addtional safety checks manually
}
