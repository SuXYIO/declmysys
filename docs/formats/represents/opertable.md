# OperTable

This document describes the format which operation representations will be in.

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
