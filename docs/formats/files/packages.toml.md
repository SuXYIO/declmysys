# packages.toml

Defines the package manager and packages to install.

> Note: This method is not flexible and is manager-dependent. If there is any library (in go) that provides a common interface for operating with package managers, pull it up in Issues!

## Structure

See example, pretty self-explainatory.

## Values

- `packages`: List of tables of managers, the managers will be executed in this order
- `packages/manager`: Put your manager's `install` command here. See formats/represents/cmd
- `packages/list`: The list of package specs
- `priority`: Default 200 for packages. See docs/files/represents/priority

### Package Spec

Example:

```toml
list = [
    "foo",
    "bar=1.0.0",    # Use your own manager's version spec
]
```

## Example

```toml
packages = [
    {
        manager = ["sudo", "apt", "install"],
        list = [
            "git",
            "neovim",
            "python=3.14",
        ]
    },
    {
        manager = ["flatpak", "install"],
        list = [
            "io.gitlab.librewolf-community",
        ]
    }
]
priority = 500
```

## Behavior

When executing packages, the manager command and element in `list` will be concated into a single command.

For example, when using list representation, the Example will be translated to command `{"sudo", "apt", "install", "git", "neovim", "python=3.14"}` and `flatpak install io.gitlab.librewolf-community".

If you use single string representation, the command will be translated to `{"bash", "-c", yourcommand}`. It's not safe, don't use it. See formats/represents/cmd.
