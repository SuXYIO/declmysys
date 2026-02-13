# Main Command

## Args

```bash
declmysys [-l|--loglevel-stderr LEVEL] [-C|--config FILE] [-h] [-v]
```

- `-l LEVEL` or `--loglevel-stderr LEVEL`: Specify log level for stderr. Default `WARN`
- `-C FILE` or `--config FILE`: Specify global config. Default `~/.declmysysrc`
- `-h` or `--help`: Command help, same as subcommand `help`
- `-V` or `--version`: Program version, same as subcommand `version`

Loglevel: choose from `DEBUG` `INFO` `WARN` `ERROR`, case insensitive.
