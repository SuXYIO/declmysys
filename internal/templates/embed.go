package templates

import "embed"

//go:embed config.toml
var Globconf []byte

//go:embed all:Dotdecl
var DDDir embed.FS

// NOTE: Used `all:` to include hidden files, in this case `.gitkeep` in empty directories,
// pls remember to exclude `.gitkeep` when copying files
