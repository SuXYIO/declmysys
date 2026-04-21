package templates

import "embed"

//go:embed config.toml
var Globconf []byte

//go:embed all:Decl
var DDir embed.FS // all: includes hidden files
