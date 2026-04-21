# Metadata

`metadata.toml` under decldir, stores metadata.

## Values

```toml
# version used when creating this decldir, don't change unless you know what you're doing
version = "{VERSION}"

# regex pattern for files to exclude, default is ["^\\.git$", "^\\.schemas$"]
exclude = ["^\\.git$", "^\\.schemas$"]

# decldir specific substitutions
[subs]

# whether to disable the homedir special subs
# set true for safety, and use "{HOME}" for homedir
# default is false
disable_homedir_subs = false

# key-value pairs for subs rules
[subs.rules]
```

`version`: version used when creating this decldir, will be used by the program to determine whether parsing is safe
`exclude`: regex pattern for files to exclude, default `["^\\.git$", "^\\.schemas$"]`, can be overrided (e.g. using `exclude = []` will not exclude `.git` and `.schemas`)
`subs`: table of subs options. See [subs](../represents/subs.md)
