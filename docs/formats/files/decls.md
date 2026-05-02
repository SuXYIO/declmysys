# Decls

Where the declarations are stored.

All directories under the decldir (except `.git/`) will be interpreted as a decl.

## Structure

Files:

1. `desc.toml`: The descriptor, describing what the declarations in the pack are for, and also how to operate it

```text
Decl/
├── .git/
├── metadata.toml
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
- `pwd`: Optional present working directory, default is the dir the desc file is under (leave empty for default)
- `args`: Optional args for execution, preset specific fields

#### Presets:

- `cmds`: Runs custom commands
  Required in `args`:
  - `cmds`: list of _cmds_ (list of list of strings) to run
- `stow`: Stows a directory. Evaluates to `stow -t={dest} {src}`
  Optional in `args`
  - `src`: string of the directory being stowed, default `stow`
  - `dest`: string of the target directory for stow, default `{HOME}`
- `gitclone`: Clones a repository to certain location. Evaluates to `git clone {src} {dest}`, see example below.
  Required in `args`:
  - `src`: string of the origin path / url
  - `dest`: string of the destination path
- `packages`: Runs a single command, but appends arguments, also includes presets, useful for installing packages.
  Required in `args`:
  - `manager`: string for a preset name, or a _cmd_ for a custom manager command
  - `packs`: _cmd_ for package names, will be appended after the `manager` command
- `manual`: Print a description, and wait for the user. Useful for operations that cannot be automated.
  Required in `args`:
  - `desc`: string for the description

Presets for `packages` (alphabetical order):
| Name | Manager Command |
| ---- | --------------- |
| `apk` | `["doas", "apk", "add"]` |
| `apt` | `["sudo", "apt", "install"]` |
| `dnf` | `["sudo", "dnf", "install"]` |

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
[args]
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
[args]
cmds = [["bash", "-c", "sudo mv sources/* /etc/apt/sources.list.d"]]
```

Add user to dialout group:

```toml
name = "add user to dialout group"
preset = "cmds"
priority = 50
[args]
cmds = [["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]]
```

Create Workspace dir:

```toml
name = "create ~/Workspace dir"
preset = "cmds"
[args]
cmds = [["mkdir", "{HOME}/Workspace"]]
```

Install packages via apt:

```toml
name = "apt"
preset = "packages"
[args]
manager = "apt"
packs = ["neovim", "tmux", "alacritty"]
```

Install packages with your own manager:

```toml
name = "foopm"
preset = "packages"
[args]
manager = ["foopm", "-i"]
packs = ["foo", "bar", "baz"]
```
