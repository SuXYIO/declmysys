# OperTable

The representation for operations.

## Format

`cmd`: Command, or list of commands. See [cmd](cmd.md). Will be processed by paths&cmds subs. See [subs](subs.md)
`affected`: List affected files, optional

> [!NOTE]
> Currently the `affected` has no use other than a note for maintainabiliy, the program won't process on it.

## Example

Copy foo.txt to bar.txt:

```toml
opertable = {
    cmd = ["cp", "foo.txt", "{HOME}/bar.txt"],
    affected = ["{HOME}/bar.txt"],
}
```
