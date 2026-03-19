# Subcommands-Run

Execute defined stuff.

## Args

```bash
declmysys run [PRIORITY]
```

- `PRIORITY`: Run the specific procedures for a certain priority

> [!WARNING]
> `go-flags` lib does not support default value for positionals,
> so program assumes `0` is not user input, and `0` means priority not specified (run all priorities)

## Example

Run all:

```console
user@host:~$ declmysys run
Running /home/user/Decl:
    (250)
        cmds[apt update]
            [sudo] password for user:
            Done!
        cmds[add flathub source to flatpak]
            Done!
    (200)
        packages[apt]
            Done!
    (150)
        packages[flatpak-system-flathub]
            Done!
    (100)
        stow[dotfiles]
            Done!
    (50)
        cmds[create ~/Workspace directory]
            Done!
        cmds[add user to dialout group]
            Done!
```

> [!Note]
> Managers might ask for stuff, have to test that.
> ~~interactive design is paradise for users, but they really make automation f\*\*ked up~~
