# decls

Where the declarations are stored.

All directories under the decldir (except `.git/`) will be interpreted as a decl.

## Structure

Files:

1. `desc.toml`: The descriptor, describing what the declarations in the pack are for, and also how to operate it

```text
Decl/
├── .git/
├── subs.toml
├── workdir/
│   └── desc.toml
├── nvim/
│   ├── desc.toml
│   └── stow/
│       └── .config/
│           └── nvim/
│               ├── init.lua
│               └── theme.lua
└── tmux/
    ├── desc.toml
    └── stow/
        └── .config/
            └── tmux/
                └── tmux.conf
```

### desc.toml

- `name`: Description name, make it human-readable. See [name](../represents/name.md)
- `preset`: Name (string) of a preset
- `priority`: Default `100` for decls. See [priority](../represents/priority.md)
- `pwd`: Optional present working directory, default is the dir the desc file is under (leave empty for default), paths & cmds subs will be used
- `rundat`: Optional data for run spec, used in presets, no specific fields, depends on preset

#### Presets:

- `cmds`: Runs custom commands
  Required in `rundat`:
  - `cmds`: list of cmds (list of list of strings) to run, will be parsed through paths & cmds subs
- `stow`: Stows a directory. Evaluates to `stow -t={dest} {src}`
  Optional in `rundat`
  - `src`: string of the directory being stowed, default `stow`
  - `dest`: string of the target dir for stow, default `{HOME}`
- `gitclone`: Clones a repository to certain location. Evaluates to `git clone {src} {dest}`, see example below.
  Required in `rundat`:
  - `src`: string of the origin path / url, parsed by global subs
  - `dest`: string of the destination path, parsed by paths & cmds subs
- `packages`: Runs a single command, but appends arguments, also includes presets, useful for installing packages.
  - `manager`: string for a preset name, or a list of strings for a custom manager command, will be parsed through paths & cmds subs if custom cmd, through global subs otherwise
  - `packs`: list of strings (cmd, see [cmd](represents/cmd.md)) for package names, will be appended after the `manager` command, will not be parsed through any subs, I don't see the need here

Presets for `packages`:
| Name | Manager Command |
| ---- | --------------- |
| `apt` | `["sudo", "apt", "install"]` |
| `apk` | `["doas", "apk", "add"]` |

> [!NOTE]
> I'm not really familiar with other package managers, so I welcome you to contribute your favorate manager's command here!

## Example

Bashrc via stow:

```toml
# Structure:
# bash/
# ├── desc.toml
# └── stow/
#     └── .bashrc

name = "bashrc"
preset = "stow"
priority = 1000
```

Clone your own repo:

```toml
# Structure:
# nvim/
# └── desc.toml

name = "clone neovim config"
preset = "gitclone"
priority = 500
[rundat]
src = "https://github.com/username/neovim_config"
dest = "{HOME}/.config/nvim"
```

Copy apt source:

```toml
# Structure:
# apt-source/
# ├── desc.toml
# └── sources/
#     ├── debian.sources
#     └── mysource.sources

name = "apt-sources"
preset = "cmds"
priority = 1000
[rundat]
cmds = [["bash", "-c", "sudo mv sources/* /etc/apt/sources.list.d"]]
```

Add user to dialout group:

```toml
name = "add user to dialout group"
preset = "cmds"
priority = 50
[rundat]
cmds = [["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]]
```

Create Workspace dir:

```toml
name = "create ~/Workspace dir"
preset = "cmds"
[rundat]
cmds = [["mkdir", "{HOME}/Workspace"]]
```

Install packages via apt:

```toml
name = "apt"
preset = "packages"
[rundat]
manager = "apt"
packs = ["neovim", "tmux", "alacritty"]
```

Install packages with your own manager:

```toml
name = "foopm"
preset = "packages"
[rundat]
manager = ["foopm", "-i"]
packs = ["foo", "bar", "baz"]
```
