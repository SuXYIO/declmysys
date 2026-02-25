package packagestoml

type Pkgs struct {
	Packages []PacksSpec `toml:"packages"`
	Priority int         `toml:"priority"`
}

type PacksSpec struct {
	Manager manSpec  `toml:"manager"`
	Packs   []string `toml:"packs"`
}
