# Subcommands-Init

Initialize new decldir.

## Args

```bash
declmysys init [--no-git]
```

- `--no-git`: Won't create the `.git/` directory (via `git init`)
- `--no-taplo`: Won't create the `.taplo.toml` file and the `.schemas` directory, enable if you don't use completion and validation

## Example

Initialize new decldir in default decldir or specified:

```bash
declmysys init
```

initialize new decldir in `~/Mydecl` (using global `-D` option):

```bash
declmysys init -D ~/Mydecl
```

## Behavior

Asks for overwrite if the directory exists and is not empty. If do overwrite, removes everything in the directory and continue, otherwise quits.

Creates the directory if not present.

Use `git init` for creating `.git/`.

Create empty file for necessary files.

Create empty directory for the user-defined sections.

See [decldir](../../formats/files/decldir.md) for structure.
