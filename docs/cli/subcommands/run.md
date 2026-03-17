# Subcommands-Run

Execute defined stuff.

## Args

```bash
declmysys run [-d|--dry] [-v|--verbose]
```

- `-d`: Dry run, only print the procedures out. The difference from `list` subcommand is that dry run just prints the named structure with command, without redundant information
- `-v`: Verbose, print verbose information, including procedure outputs and stats

## Example

Run all:

```console
user@host:~$ declmysys run
Run /home/user/Decl:
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

Dry run:

```console
user@host:~$ declmysys run -d
Run /home/user/Decl (dry):
    (250)
        cmds[apt update]
            ["sudo", "apt", "update"]
        cmds[add flathub source to flatpak]
            ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
    (200)
        packages[apt]
            ["sudo", "apt", "install", "-y", "zsh", "git", "tmux", "fastfetch", "neovim", "flatpak", "thunderbird", "librewolf", "kitty"]
    (150)
        packages[flatpak-system-flathub]
            ["flatpak", "install", "flathub", "--noninteractive", "-y", "--user", "com.valvesoftware.Steam", "com.visualstudio.code"]
    (100)
        stow[dotfiles]
            ["stow", "zshrc"]
            ["stow", "git"]
            ["stow", "tmux"]
            ["stow", "fastfetch"]
            ["stow", "neovim"]
            ["stow", "kitty"]
            ["stow", "apt-sources"]
    (50)
        cmds[create ~/Workspace directory]
            ["mkdir", "{HOME}/Workspace"]
        cmds[add user to dialout group]
            ["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]
```
