# Command

A command will be described as either a _list of strings_ or a _single string_.

## Format

### List

It follows the default format for go's `os/exec.Command` function, which is a list.

Why? It's much safer than a single string, see golang docs for reason.

It will be executed via `exec.Command(l[0], l[1:]...)` where `l` is the list of commands you passed in.

### Single

Don't use this for god's sake, it's unsafe, it's slower and stuff.

It will be executed via `exec.Command("bash", "-c", s)` where `s` is the single string you passed in.

### Macros

A macro expands to a certain content. List of macros:

- `{HOME}`: Expands to your home directory, e.g. `/home/foobar` for user `foobar`
- `{USERNAME}`: Expands to your user name, e.g. `foobar` for user `foobar`

> [!NOTE]
> This is an awkward design, but due to safety concerns and sticking to lists, the `~` cannot be interpreted (it is a shell feature), so I have to do it this way.
> Theoretically the single string representation or sudo can process this since it is interpreted via bash or the shell for root user.

## Example

### List

- Command `sudo rm -rf /*` should be written as `["sudo", "rm", "-rf", "/*"]`
- Command `echo 'foo bar'` should be written as `["echo", "'foo bar'"]`
- Command `cp ~/Downloads/nvim.bak ~/.config/nvim` should be written as `["cp", "{HOME}/Downloads/nvim.bak", "~/.config/nvim"]`

### Single

- Command `sudo rm -rf /*` should be written as `"sudo rm -rf /*`
