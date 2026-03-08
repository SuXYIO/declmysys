package packagestoml

import "github.com/suxyio/declmysys/internal/parse/priority"

type Pkgs struct {
	Packages []PacksSpec       `toml:"packages"`
	Priority priority.Priority `toml:"priority"`
}

type PacksSpec struct {
	Manager manSpec  `toml:"manager"`
	Packs   []string `toml:"packs"`
}
