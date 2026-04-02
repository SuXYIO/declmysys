# Metadata

`metadata.toml` under decldir, stores metadata.

## Values

```toml
exclude = [
    "^.git",
    "^.myfiletoexclude"
]

[subs]
    disable_homedir_subs = true
    [subs.rules]
        "{GREET}" = "Hello"
        "{NAME}" = "Dave"
```

`exclude`: list of regex filenames to exclude when searching for decls, default `["^.git"]`
`subs`: table of subs options. See [subs](../represents/subs.md)
