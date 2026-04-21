# DeclDir

It is the directory where the dotfiles and declarations are stored.

## Default

The default decldir is `~/Decl`, you can use another by specifying in global config file, or passing argument in the cli.

## Structure

```text
Decl/
├── .git/
├── .taplo.toml
├── .schemas/
├── metadata.toml
├── decl1/
├── decl2/
└── ...
```

- `.git/`: optional regular git directory, for version control
- `.taplo.toml` and `.schemas/`: config file for taplo, providing toml completion and validation rules
- `metadata.toml`: defines metadata. See [metadata](metadata.md)
- decls: stores declarations. See [decls](decls.md)

## Updating

The `version` value in `metadata.toml` offers a check for the cli, refusing to operate when version does not match.

Currently, I haven't come up with any good solutions (except just installing the correct version), so if you want to migrate your decls to a newer version, here's a path:

- Use `init` subcommand to create a new decldir (this step creates right default files for the version, and new `.schemas` which'll help you validate your decls with Taplo linter)
- Move your files to the new one, after changing them to the new decls syntax (or just merge the new and old one manually)
