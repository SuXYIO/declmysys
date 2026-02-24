# packages.toml

Defines the package manager and packages to install.

> Note: This method is not flexible and is manager-dependent. If there is any library (in go) that provides a common interface for operating with package managers, pull it up in Issues!

## Structure

See example, pretty self-explainatory.

## Values

I'll use `packages//` for any subtable under `packages` list.

- `packages`: List of tables of managers, the managers will be executed in this order
- `packages//manager`: A preset manager name, or put your manager's `install` command here. Note that the manager must support passing multiple packages to command at the same time (e.g. `["manager", "install", "pkg1", "pkg2", "pkg3]` works), managers that doesn't support this is currently not supported
- `packages//packs`: The list of package specs
- `priority`: Default 200 for packages. See [priority](../represents/priority.md)

> [!NOTE]
> Doesn't see the reason for setting manager specific priority, the install operation is executed in the order of the `packages` list

### Package Spec

It's manager-dependent so good luck.

Example:

```toml
packs = [
    "foo",
    "bar=1.0.0",    # Use your own manager's version spec format
]
```

## Preset Managers

Not gonna add much, since I usually only deal with these. Adding these is pretty dangerous, let alone adding something I've never used.

Welcome to add more via Pull Request.

`apt`: `["sudo", "apt", "install", "-y"]`
`flatpak-user-flathub`: `["flatpak", "install", "flathub", "--noninteractive", "-y", "--user"]`
`flatpak-system-flathub`: `["flatpak", "install", "flathub", "--noninteractive", "-y", "--system"]`

## Example

```toml
packages = [
    { manager = "apt", packs = ["git", "neovim", "python=3.14"] },
    # or spec it yourself
    { manager = ["flatpak", "install", "mysource", "--noninteractive", "-y", "--user"], packs = ["com.valvesoftware.Steam", "com.visualstudio.code"] }
]
priority = 500
```

> [!NOTE]
> Notice it's flatpak-flathub instead of flatpak. It looks better visually when listing.
> It prints `[flatpak-flathub]: com.valvesoftware.Steam  com.visualstudio.code` instead of `[flatpak-flathub]: flathub com.valvesoftware.Steam  flathub com.visualstudio.code`

## Behavior

When executing packages, the manager command and element in `packs` will be concated into a single command.

For example, when using list representation, the Example will be translated to command `{"sudo", "apt", "install", "git", "neovim", "python=3.14"}` and `flatpak install io.gitlab.librewolf-community"`.
