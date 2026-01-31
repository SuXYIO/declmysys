# OperTable

The representation for operations.

## Format

`cmd`: Command, or list of commands. See docs/formats/represents/cmd
`affected`: List affected files, optional

## Example

Copy foo.txt to bar.txt:

```toml
opertable = {
    cmd = ["cp", "foo.txt", "bar.txt"],
    affected = ["bar.txt"],
}
```
