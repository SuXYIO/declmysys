# DeclDir

It is the directory where the dotfiles and declarations are stored.

## Default

The default decldir is `~/Decl`, you can use another by specifying in global config file, or passing argument in the cli.

## Structure

```text
Decl/
├── .git/
├── metadata.toml
├── decl1/
├── decl2/
└── ...
```

- `.git/`: optional regular git directory, for version control. Created via `git init`
- `metadata.toml`: defines metadata. See [metadata](metadata.md)
- decls: stores declarations. See [decls](decls.md)
