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
- `run`: Can be a name (string) of a preset, or a list of cmds. See [cmd](../represents/cmd.md)
- `priority`: Default `100` for dots. See [priority](../represents/priority.md)
- `rundat`: Optional data for run spec, used in presets, no specific fields, depends on preset

Example for preset:

```toml
name = "foobar"
run = "stow"
priority = 1000
```

Example for custom:

```toml
run = [["cp", "data/foo.txt", "/root/bar.txt"],["rm", "/root/baz"]]
```

Presets:

- `stow`: Processes the `data/` directory with `stow data` command
- `gitclone`: Clones a repository to certain location, requires `url`, `dest` to be set in `run_dat`, translates to `git clone {url} {dest}`, see example below

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

name = ".bashrc dots"
run = "stow"
priority = 1000
```

Clone your own repo:

```toml
# Structure:
# nvim/
# └── desc.toml

name = "clone neovim config"
run = "gitclone"
priority = 500  # Maybe this can wait a little?
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
run = [["bash", "-c", "sudo mv data/* /etc/apt/sources.list.d"]]
priority = 1000
```
