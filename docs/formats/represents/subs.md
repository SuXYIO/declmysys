# Subs

Dotdecldir substitutions.

## Defaults

### Global

For every string in toml files.

- `{USERNAME}` is replaced with username (login name)
- `{NAME}` is replaced with username (display name)
- `{HOSTNAME}` is replaced with hostname

> [!NOTE]
> Check out the [type User: Go os/user Package Docs](https://pkg.go.dev/os/user#User) for differences between login name and display name.
> `userinfo.Username` is the login name, and `userinfo.Name` is the display name (where `userinfo, err := user.Current()`), this program uses this API to fetch these names.
> It's recommended to use login name over display name, since display name is optional (might be empty), while login name is mandatory.

### Paths & Cmds

Applies for filepaths, cmds, and even some command line args (see the specific docs for details).

Special homedir subs:

- First character `~` is replaced with homedir
- `~` with spaces on it's left and right, is replaced with ` ` + homedir + ` `
- `~/` with space on it's left, is replaced with ` ` + homedir + `/`

`homedir` is probably `/home/foo` if your username is `foo`. (Unless you set a homedir manually, the `os` package'll probably correctly get the path anyway)

If you think of it in a command line scenario, these rule'll be more obvious.

> [!NOTE]
> You can disable this special subs in config. I personally recommend disabling this and use the safe `{HOME}` subs.

Other:

- `{HOME}` is replaced with user home dir
- `{CONF}` or `{CONFIG}` will be replaced with user config dir
- `{CACHE}` will be replaced with user cache dir
- `{TMP}` will be replaced with system provided temporary dir

## Subs table

### Values

- `subs`: the subs table, overrides default and dotdecldir subs. See [substitutions](../formats/represents/substitutions.md)
  - `disable_homedir_subs`: whether to disable the special homedir substitution feature (won't disable `{HOME}`, don't worry)
  - `global`: set of subs rules for any string parsed
  - `files_cmds`: set of subs rules specific for filepaths and commands

Brackets are recommended since they are lessly used in naming, but not forced.

Subs Rules is just a dict/map (key value pairs), specifying replacement from string to string.
Example:

```toml
[example_subs_rules]
"{GREET}" = "Hello"
"{NAME}" = "Dave"
```

it turns string `{GREET}, {NAME}.` to `Hello, Dave.`

### Order

The rest of the subs defined in your dotdecldir is applied before the defaults, so that it enables aliases.
For example, if you love git or vim style homedir specification, you can specify

```toml
"%USERPROFILE%" = "{HOME}"
```

where after your custom substitution, it will be substituted again by the defaults, turning `%USERPROFILE%` to the actual home dir.

> [!NOTE]
> Sadly, there are limitations to this rough design, for example you can override defaults, but you cannot disable defaults.
> Create a enhancement/feature request Issue, or pull up in the discussions, if you have a better idea. (and i hope it's not regex)

### Example

> [!NOTE]
> Cannot represent complex substitutions. Sorry im not a regex fan.

> [!WARN]
> Substitutions are not garanteed to be safe, use with causion.

```toml
[subs]
disable_homedir_subs = false
[[global]]
[[files_cmds]]
```
