# dots/

Where the dotfiles are stored.

## Structure

The top `dots/` directory contains many subdirectories, I'll call them _dotpacks_.

Each _dotpack_ consists of two major parts:

1. `desc.toml`: The descriptor, describing what the dotfiles in the pack are for, and also how to operate it (via `stow` by default)
2. `data/`: The optional content, which is usually a `stow` structure or something the descriptor can operate with

```
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

#### Structure

- `name`: Description name, make it human-readable. See docs/formats/represents/name
- `proc`: Can be a string of a built-in processor, or a table describing the processing operation needed. See formats/represents/opertable.toml.
- `priority`: Default `100` for dots. See docs/formats/represents/priority

Example for built-in:

```toml
name = "foobar"
proc = "stow"
priority = 1000
```

Example for opertable:

```toml
proc = [
    cmd = [
        ["cp", "data/foo.txt", "/root/bar.txt"],
        ["rm", "/root/baz.txt"],
    ],
    affected = [
        "bar.txt",
        "baz.txt",
    ],
]
```

Built-in processors:

- `stow`: Processes the `data/` directory with `stow data` command
- `git`: Copies a repository to certain location, requires `url`, `dest` and `affected` to be set in desc, translates to `git clone {url} {dest}`, see example below

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
proc = "stow"
priority = 1000
```

Clone your own repo:

```toml
# Structure:
# nvim/
# └── desc.toml

name = "clone neovim config"
proc = "git"
url = "https://github.com/username/neovim_config"
dest = "{HOME}/.config/nvim"
priority = 500  # Maybe this can wait a little?
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
proc = [
    cmd = [
        ["sudo", "mv", "data/*", "/etc/apt/sources.list.d"]
    ]
    affected = ["/etc/apt/sources.list.d"]
]
priority = 1000
```
