# packages.toml

Defines the package manager and packages to install.

## Default

```toml
[manager]
preset = ""

[packages]
list = []
```

## Values

- `manager/preset`: Choose from a set of strings of common managers built-in the code, or use a custom manager spec table
- `packages/list`: The list of package specs

### Package Spec

Example:

```toml
{
    name = "foo",
    version = "3.14",
    channel = "stable",
    manager = "apt",
}
```

quite self-explainatory, note that the `version` and `channel` and `manager` are optional, and `manager` falls back to the `manager.preset` value.

### Manager Spec Table

Not really sure how to represent this yet, I wish there's a lib to do this without figuring out everything myself.

Example for `apt`:

```toml
{
    command = "apt",
    flags = "",
    require_root = true,
    namever_fmt = "{name}={ver}"
    use = {
        install = "{cmd} install {namever}",
        remove = "{cmd} remove {namever}"
        upgrade = "{cmd} upgrade {namever}",
        update = "{cmd} update"
    }
}
```
