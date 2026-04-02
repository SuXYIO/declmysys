package metadata

type Metadata struct {
	Exclude     []string `toml:"exclude"`
	SubsDef     `toml:"subs"`
	initialized bool
}

type SubsDef struct {
	SpecialHDDisable bool      `toml:"disable_homedir_subs"`
	CustomRules      SubsRules `toml:"rules"`
}

// GlobalMetaData stores global metadata, load before use
var GlobalMetaData = Metadata{
	initialized: false,
}
