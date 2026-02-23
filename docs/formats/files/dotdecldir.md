# DotDeclDir

It is the directory where the dotfiles and declarations are stored.

## Default

The default dotdecldir is `~/Dotdecl/`, and the cli should support for modifying this directory.

## Structure

```
Dotdecl
├── .git/
├── packages.toml
├── subs.toml
├── dots/
│   ├── dot1/
│   ├── dot2/
│   └── ...
└── actions/
    ├── action1/
    ├── action2/
    └── ...
```

- `.git/`: optional regular git directory, for version control. Created via `git init`
- `subs.toml`: subs table file. See [subs](../represents/subs.md)
- `packages.toml`: defines package manager to use, and also packages to be installed. See [packages](packages.toml.md)
- `dots/`: stores dotfile definitions. See [dots](dots.md)
- `actions/`: stores action definitions. See [actions](actions.md)
