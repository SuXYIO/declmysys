package decls

// Decls is a slice of Decl, just for []Decl to be receiver type
type Decls []Decl

// Decl is the info for a single declaration
type Decl struct {
	Name     string         `toml:"name"`
	Preset   string         `toml:"preset"`
	Priority Priority       `toml:"priority"`
	Pwd      string         `toml:"pwd"`
	RunDat   map[string]any `toml:"rundat"`
	descPath string         // will not be loaded by toml, only for setting default pwd
}
