# Main Command

## Args

```bash
declmysys [-D DOTDECLDIR] [-l|--loglevel LEVEL] [-L|--logfile LOGFILE] [-C|--config FILE] [-h] [-V]
```

- `-D DOTDECLDIR`: The dotdecldir path to operate. Default: the one specified in global config, or default.
- `-l LEVEL` or `--loglevel LEVEL`: Specify log level. One of `DEBUG` `INFO` `WARN` `ERROR`, case-insensitive. Default: `WARN`
- `-L LOGFILE` or `--logfile LOGFILE`: Specify log file. Default: write to stderr
- `-C FILE` or `--config FILE`: Specify global config. Default: `{CONF}/config.toml`
- `-h` or `--help`: Command help, same as subcommand `help`
- `-V` or `--version`: Program version, same as subcommand `version`
