# actions/

Where the actions are stored.

## Structure

- `cmd`: List of commands. See docs/formats/represents/cmd
- `affected`: Optional affected files, just for convenience
- `priority`: Default `50` for actions. See docs/files/represents/priority

## Example

```toml
cmd = [
    ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
]
affected = []
priority = 0
```
