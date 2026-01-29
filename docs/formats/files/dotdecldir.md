# DotDeclDir

It is the directory where the dotfiles and declarations are stored.

## Default

The default dotdecldir is `~/Dotdecl/`, and the cli should support for modifying this directory.

## Structure

```
DotDecl
├── .git/
├── packages.toml
├── dots/
│   ├── dot1/
│   ├── dot2/
│   └── ...
└── actions/
    ├── action1/
    ├── action2/
    └── ...
```

- `.git/`: regular git directory, for version control. Created via `git init`
- `packages.toml`: defines package manager to use, and also packages to be installed. See cli/packages.toml
- `dots/`: stores dotfile definitions. See cli/dots
- `actions/`: stores action definitions. See cli/actions
