# Subcommands-Init

Inits new dotdecldir in directory.

## Args

```bash
declmysys init [DOTDECLDIR]
```

- `DOTDECLDIR`, the _dotdecldir_ path to be initialized. default: `~/Dotdecl`

## Example

Initialize new dotdecldir in `~/Dotdecl`:

```bash
declmysys init
```

initialize new dotdecldir in `~/Mydotdecl`:

```bash
declmysys init ~/Mydotdecl
```

## Behavior

Asks for overwrite if the directory exists and is not empty. If do overwrite, removes everything in the directory and continue, otherwise quits.

Creates the directory if not present.

Use `git init` for creating `.git/`.

Create empty file for necessary files.

Create empty directory for the user defined sections.

See fotmats/files/dotdecldir for structure.
