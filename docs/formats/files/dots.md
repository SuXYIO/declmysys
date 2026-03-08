# dots/

Where the dotfiles are stored.

## Structure

The top `dots/` directory contains many subdirectories, I'll call them _dotpacks_.

Each _dotpack_ consists of two major parts:

1. `desc.toml`: The descriptor, describing what the dotfiles in the pack are for, and also how to operate it (via `stow` by default)
2. `data/`: The optional content, which is usually a `stow` structure or something the descriptor can operate with

```text
dots/
├── bash/
│   ├── desc.toml
│   └── data/
│       └── .bashrc
├── nvim/
│   ├── desc.toml
│   └── data/
│       └── .config/
│           └── nvim/
│               ├── init.lua
│               └── theme.lua
└── tmux/
    ├── desc.toml
    └── data/
        └── .config/
            └── tmux/
                └── tmux.conf
```

### desc.toml

- `name`: Description name, make it human-readable. See [name](../represents/name.md)
- `preset`: Name (string) of a preset
- `priority`: Default `100` for dots. See [priority](../represents/priority.md)
- `rundat`: Optional data for run spec, used in presets, no specific fields, depends on preset

```toml
name = "foobar"
preset = "stow"
priority = 1000
```

#### Presets:

- `cmds`: Runs custom commands
  Required in `rundat`:
  - `cmds`: list of cmds (`[]Cmd` i.e. `[][]string`) to run, will be parsed through paths & cmds subs
- `stow`: Processes the `data/` directory with `stow data` command
  Optional in `rundat`:
  - `datadir`: default `data`, string of the directory being stowed, relative path is recommended, but still parsed through paths & cmds subs
- `gitclone`: Clones a repository to certain location. Translates to `git clone {url} {dest}`, see example below.
  Required in `rundat`:
  - `url`: string of the origin url, parsed by global subs
  - `dest`: string of the destination path, parsed by paths & cmds subs

### data/

Optional.

Usually a `stow` structure. It is the common way people manage dotfiles so it should be familiar.
(if you're not familiar, check out [GNU Stow Docs](https://www.gnu.org/software/stow/manual/), or search it online, there's plenty of tutorials for this)

You can also use your own way, as long as you declare the operations properly in your `desc.toml`.

## Example

Bashrc via stow:

```toml
# Structure:
# bash/
# ├── desc.toml
# └── data/
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
url = "https://github.com/username/neovim_config"
dest = "{HOME}/.config/nvim"
```

Copy apt source:

```toml
# Structure:
# apt-source/
# ├── desc.toml
# └── data/
#     ├── debian.sources
#     └── extrepo.sources

name = "apt-sources"
preset = "cmds"
priority = 1000
[rundat]
cmds = [["bash", "-c", "sudo mv data/* /etc/apt/sources.list.d"]]
```
