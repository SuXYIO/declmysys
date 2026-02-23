# GlobalConf

The cli should read from a global (user specific) config file for certain values.

## Default

The default file is `~/.config/declmysys/config.toml`(`{CONF}/.config/declmysys/config.toml`), and is in _toml_ format.

Will add a search order later. If no matches found, create one automatically.

## Values

- `dotdecldir`, the directory path of _dotdecldir_. Will be parsed through the default paths&cmds subs rules

## Example

```toml
dotdecldir = "~/Dotdecl"
```
