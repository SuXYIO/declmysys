# config

The CLI should read from a global (user-specific) config file for certain values.

You can also specify the path via environment variable `DECLMYSYS_CONFIG`, which overrides the default.

## Default

The default dir is `~/.config/declmysys`.

Will ask to create one if not present.

## Files

```text
Decl/
└── config.toml
```

### config.toml

Values:

- `decldir`: the directory path of decldir, default `~/Decl`
