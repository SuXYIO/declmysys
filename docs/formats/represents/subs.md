# Subs

Dotdecldir substitutions.

## Defaults

### Global

For all strings.

- `{USER}` is replaced with username
- `{HOST}` is replaced with hostname

### Files & Cmds

Applies for filepaths, cmds, and even some command line args.

Special homedir subs:

- First character `~` or any ` ~` (notice the space) is replaced with user home dir

> [!NOTE]
> You can disable this special subs in config. I personally recommend disabling this and use the safe `{HOME}` subs.

Other:

- `{HOME}` is replaced with user home dir
- `{CONF}` or `{CONFIG}` will be replaced with user config dir
- `{TMP}` will be replaced with system provided temporary dir

## Subs table

### Values

- `subs`: the subs table, overrides default and dotdecldir subs. See [substitutions](../formats/represents/substitutions.md)
  - `disable_homedir_subs`: whether to disable the special homedir substitution feature
  - `global`: set of subs rules for any string parsed
  - `files_cmds`: set of subs rules specific for filepaths and commands

Brackets are recommended since they are lessly used in naming, but not forced.

Subs Rules is just a dict/map (key value pairs), specifying replacement from string to string.
Example:

```toml
[example_subs_rules]
"{GREET}" = "Hello"
"{NAME}" = "Dave"
```

it turns string `{GREET}, {NAME}.` to `Hello, Dave.`

### Example

> [!NOTE]
> Cannot represent complex substitutions. Sorry im not a regex fan.

> [!WARN]
> Substitutions are not garanteed to be safe, use with causion.

```toml
[subs]
disable_homedir_subs = false
[[global]]
[[files_cmds]]
```
