# actions

Where the actions are stored.

Since actions are single files, for procedures that involve copying / writing files (especially large files), e.g. copying apt sources, using dots is recommended. But I think you know better, giving the user enough freedom.

## Structure (actions/)

Contains action toml files.

Example:

```text
actions
├── foo.toml
├── bar.toml
└── baz.toml
```

## Structure (action toml file)

- `name`: Description name, make it human-readable. See [name](../represents/name.md)
- `run`: List of commands. See [cmd](../represents/cmd.md)
- `affected`: Optional list of affected files. No action made yet, maybe add backup feature later
- `priority`: Default `50` for actions. See [priority](../represents/priority.md)

## Example

```toml
name = "add flathub source to flatpak"
run = [
    ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
]
priority = 250
```

```toml
name = "add user to dialout group"
run = [
    ["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]
]
priority = 50
```

```toml
name = "create ~/Workspace dir"
run = [
    ["mkdir", "{HOME}/Workspace"]
]
affected = ["{HOME}/Workspace"]
```
