# Subs

Dotdecldir substitutions.

## Defaults

### Global

For all strings.

- `{USER}` is replaced with username
- `{HOST}` is replaced with hostname

### Files & Cmds

Applies for filepaths, cmds, and even some command line args.

- First character `~` or any ` ~` (notice the space) is replaced with user home dir
- `{HOME}` is replaced with user home dir
- `{CONF}` or `{CONFIG}` will be replaced with user config dir
- `{TMP}` will be replaced with system provided temporary dir

## Subs table

### Values

- `subs`, the subs table, overrides default and dotdecldir subs. See [substitutions](../formats/represents/substitutions.md)

Brackets are recommended since they are lessly used in naming, but not forced.

### Example

> [!NOTE]
> Cannot represent complex substitutions. Sorry im not a regex fan.

> [!WARN]
> Substitutions are not garanteed to be safe, use with causion.

```toml
[subs]
global = {}
files_cmds = {}
```
