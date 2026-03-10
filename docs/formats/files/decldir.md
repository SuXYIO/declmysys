# DeclDir

It is the directory where the dotfiles and declarations are stored.

## Default

The default decldir is `~/Decl`, and the cli should support for modifying this directory.

## Structure

```text
Decl/
├── .git/
├── packages.toml
├── subs.toml
└── decls/
    ├── decl1/
    ├── decl2/
    └── ...
```

- `.git/`: optional regular git directory, for version control. Created via `git init`
- `subs.toml`: subs table file. See [subs](../represents/subs.md)
- `packages.toml`: defines package manager to use, and also packages to be installed. See [packages](packages.toml.md)
- `decls/`: stores declarations. See [decls](decls.md)
