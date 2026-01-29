# Command

This document describes the format which commands will be in.

A command will be described as either a _list of strings_ or a _single string_.

## Format

### List

It follows the default format for go's `os/exec.Command` function, which is a list.

Why? It's much safer than a single string, see golang docs for reason.

It will be executed via `exec.Command(l[0], l[1:]...)` where `l` is the list of commands you passed in.

### Single

Don't use this for god's sake, it's unsafe, it's slower and stuff.

It will be executed via `exec.Command("bash", "-c", s)` where `s` is the single string you passed in.

## Example

### List

- Command `sudo rm -rf /*` should be written as `["sudo", "rm", "-rf", "/*"]`
- Command `echo 'foo bar'` should be written as `["echo", "'foo bar'"]`

### Single

- Command `sudo rm -rf /*` should be written as `"sudo rm -rf /*`
