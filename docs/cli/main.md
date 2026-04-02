# Main Command

## Args

```bash
declmysys [-D|--decldir DECLDIR] [-C|--config FILE] [-h|--help] [-V|--version] [SUBCOMMAND]
```

- `SUBCOMMAND`: The subcommand to run, must be specified unless `-h` or `-V` is set
- `-D DECLDIR`: The decldir path to operate. Default: the one specified in global config, or default
- `-C FILE` or `--config FILE`: Specify global config. Default: `{CONF}/config.toml`
- `-h` or `--help`: Command help, similar to subcommand `help`
- `-V` or `--version`: Program version, same as subcommand `version`
