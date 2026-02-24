package globconf

type Globconf struct {
	DDDir string `toml:"dotdecldir"`
}

var DefaultGlobconf Globconf = Globconf{
	DDDir: "~/Dotdecl",
}
