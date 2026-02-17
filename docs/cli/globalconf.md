# GlobalConf

The cli should read from a global (user specific) config file for certain values.

## Default

The default file is `~/.config/declmysys/config.toml`, and is in _toml_ format.

Will add a search order later. If no matches found, create one automatically.

## Values

- `dotdecldir`, the directory path of _dotdecldir_
- `glyphset`, the set of characters the program uses for better visual text. Currently only supports `ascii`, will add NerdFonts and Emojis option later.
- `subs`, the subs table, overrides default and dotdecldir subs. See [substitutions](../formats/represents/substitutions.md)

## Example

```toml
dotdecldir = "~/Dotdecl"
glyphset = "ascii"

[subs]
global = {}
files_cmds = {}
```
