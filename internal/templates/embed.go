package templates

import "embed"

//go:embed config.toml
var Globconf []byte

//go:embed Dotdecl
var DDDir embed.FS
