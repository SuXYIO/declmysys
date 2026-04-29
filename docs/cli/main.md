# Main Command

## Args

```bash
declmysys [-D|--decldir DECLDIR] [-C|--config FILE] [-h|--help] [-v|--version] [SUBCOMMAND]
```

- `SUBCOMMAND`: The subcommand to run, must be specified unless `-h` or `-v` is set
- `-D DECLDIR`: The decldir path to operate. Default: the one specified in global config, or default
- `-C FILE` or `--config FILE`: Specify global config. Default: `{CONF}/config.toml`
- `-h` or `--help`: Command help, similar to subcommand `help`
- `-v` or `--version`: Program version, same as subcommand `version`
