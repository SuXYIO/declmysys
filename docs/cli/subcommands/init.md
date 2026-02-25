# Subcommands-Init

Initialize new dotdecldir.

## Args

```bash
declmysys init [--no-git]
```

- `--no-git`: Won't create the `.git/` directory (via `git init`)

## Example

Initialize new dotdecldir in default dotdecldir or specified:

```bash
declmysys init
```

initialize new dotdecldir in `~/Mydotdecl` (using global `-D` option):

```bash
declmysys init -D ~/Mydotdecl
```

## Behavior

Asks for overwrite if the directory exists and is not empty. If do overwrite, removes everything in the directory and continue, otherwise quits.

Creates the directory if not present.

Use `git init` for creating `.git/`.

Create empty file for necessary files.

Create empty directory for the user-defined sections.

See [dotdecldir](../../formats/files/dotdecldir.md) for structure.
