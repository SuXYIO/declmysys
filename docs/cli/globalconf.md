# GlobalConf

The CLI should read from a global (user-specific) config file for certain values.

## Default

The default dir is `~/.config/declmysys`(`{CONF}/.config/declmysys/config.toml`).

Will ask to create one if not present.

## Files

```text
Decl/
└── config.toml
```

### config.toml

Values:

- `decldir`: the directory path of decldir, default `~/Decl`
