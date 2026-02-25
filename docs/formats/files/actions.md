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
- `do`: List of commands. See [cmd](../represents/cmd.md)
- `undo`: Optional list of commands, that revert the changes
- `affected`: Optional list of affected files. No action made yet, maybe add backup feature later
- `priority`: Default `50` for actions. See [priority](../represents/priority.md)

## Example

```toml
name = "add flathub source to flatpak"
do = [
    ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
]
undo = [
    ["flatpak", "remote-delete", "flathub"]
]
priority = 250
```

```toml
name = "add user to dialout group"
do = [
    ["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]
]
undo = [
    ["sudo", "gpasswd", "-d", "{USERNAME}", "dialout"]
]
priority = 50
```

```toml
name = "create ~/Workspace dir"
do = [
    ["mkdir", "{HOME}/Workspace"]
]
# Not recommended to add
# undo = [
#     ["rm", "-rf", "{HOME}/Workspace"]
# ]
# since it might be destructive
affected = ["{HOME}/Workspace"]
```
