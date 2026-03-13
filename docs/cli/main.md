# Main Command

## Args

```bash
declmysys [-D|--decldir DECLDIR] [-l|--loglevel LEVEL] [-L|--logfile LOGFILE] [-C|--config FILE] [-h|--help] [-V|--version] [SUBCOMMAND]
```

- `SUBCOMMAND`: The subcommand to run, must be specified unless `-h` or `-V` is set
- `-D DECLDIR`: The decldir path to operate. Default: the one specified in global config, or default. Won't be parsed by any subs
- `-l LEVEL` or `--loglevel LEVEL`: Specify log level. One of `DEBUG` `INFO` `WARN` `ERROR`, case-insensitive. Default: `WARN`
- `-L LOGFILE` or `--logfile LOGFILE`: Specify log file. Default: write to stderr
- `-C FILE` or `--config FILE`: Specify global config. Default: `{CONF}/config.toml`. Will be parsed through the default paths&cmds subs rules
- `-h` or `--help`: Command help, similar to subcommand `help`
- `-V` or `--version`: Program version, same as subcommand `version`
