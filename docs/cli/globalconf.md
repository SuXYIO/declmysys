# GlobalConf

The CLI should read from a global (user-specific) config file for certain values.

## Default

The default file is `~/.config/declmysys/config.toml`(`{CONF}/.config/declmysys/config.toml`), and is in _toml_ format.

Will add a search order later. If no matches found, create one automatically.

## Values

- `decldir`, the directory path of \_decldir. Will be parsed through the default paths&cmds subs rules

## Example

```toml
decldir = "~/Decl"
```
