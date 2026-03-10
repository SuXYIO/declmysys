package globconf

type Globconf struct {
	DDir string `toml:"decldir"`
}

var DefaultGlobconf Globconf = Globconf{
	DDir: "~/Decl",
}
