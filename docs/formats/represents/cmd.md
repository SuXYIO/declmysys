# Command

A command will be described as a _list of strings_.

Will be parsed through paths&cmds subs rules. See [subs](../represents/subs.md), all elements of the list will be parsed separately.

## Format

It follows the default format for go's `os/exec.Command` function, which is a list.

Why? It's much safer than a single string, see golang docs for reason.

It will be executed via `exec.Command(l[0], l[1:]...)` where `l` is the list of commands you passed in.

> [!NOTE]
> Why not use a single string? Well, using a list of strings has many benefits:
>
> 1. Safety. The command line arguments are passed clearly, non-ambiguously, and prevents command injection.
> 2. Performance. The command can be executed without interpreting by shell (theoretically, I'm not an expert on this).
> 3. Easy for implementation. Well, as you might know, Go is not good at handling generic types, heck it doesn't have a union. So if using multiple types, it's hard for me to implement it.
>    ~~Actually I intended to support single string command representations at first, but threw it away when having a hard time painstakingly implementing generic command types~~

## Example

- Command `sudo rm -rf /*` should be written as `["sudo", "rm", "-rf", "/*"]`
- Command `echo 'foo bar'` should be written as `["echo", "foo bar"]`
