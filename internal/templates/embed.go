package templates

import "embed"

//go:embed config.toml
var Globconf []byte

//go:embed all:Decl
var DDir embed.FS

// NOTE: Used `all:` to include hidden files, in this case `.gitkeep` in empty directories,
// pls remember to exclude `.gitkeep` when copying files
