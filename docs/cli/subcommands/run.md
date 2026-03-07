# Subcommands-Run

Execute defined stuff.

## Args

```bash
declmysys run [-d|--dry] [-v|--verbose] [PROCEDURE]
```

- `-d`: Dry run, only print the procedures out. The difference from `list` subcommand is that dry run just prints the named structure with command, without redundant information
- `-v`: Verbose, print verbose information, including procedure outputs and stats
- `PROCEDURE`: Specify the procedure to run, e.g. `dots.foobar`. See [procedure spec](../../formats/represents/procedure-spec.md)

## Example

Run all:

```console
user@host:~$ declmysys run
Run /home/user/Dotdecl:
    - Actions (250)
        apt update
    [sudo] password for user:
    Done!
        add flathub source to flatpak
    Done!
    - Packages (200)
        [apt]: zsh  git  tmux  fastfetch  neovim  flatpak  thunderbird  librewolf  kitty
    Done!
    - Packages (150)
        [flatpak-system-flathub]: com.valvesoftware.Steam com.visualstudio.code
    Done!
    - Dots (100)
        zshrc  git  tmux  fastfetch neovim  kitty  apt-sources
    Done!
    - Actions (50)
        create ~/Workspace directory
    Done!
        add user to dialout group
    Done!
```

> [!Note]
> Managers might ask for stuff, have to test that.
> ~~interactive design is paradise for users, but they really make automation f\*\*ked up~~

Dry run:

```console
user@host:~$ declmysys run -d
Run /home/user/Dotdecl (dry run):
    - Actions (250)
        apt update
            ["sudo", "apt", "update"]
        add flathub source to flatpak
            ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
    - Packages (200)
        [apt]: zsh  git  tmux  fastfetch  neovim  flatpak  thunderbird  librewolf  kitty
            ["sudo", "apt", "install", "-y", "zsh", "git", "tmux", "fastfetch", "neovim", "flatpak", "thunderbird", "librewolf", "kitty"]
    - Packages (150)
        [flatpak-system-flathub]: com.valvesoftware.Steam com.visualstudio.code
            ["flatpak", "install", "flathub", "--noninteractive", "-y", "--user", "com.valvesoftware.Steam", "com.visualstudio.code"]
    - Dots (100)
        zshrc  git  tmux  fastfetch neovim  kitty  apt-sources
            ["stow", "zshrc"]
            ["stow", "git"]
            ["stow", "tmux"]
            ["stow", "fastfetch"]
            ["stow", "neovim"]
            ["stow", "kitty"]
            ["stow", "apt-sources"]
    - Actions (50)
        create ~/Workspace directory
            ["mkdir", "{HOME}/Workspace"]
        add user to dialout group
            ["sudo", "gpasswd", "-a", "{USERNAME}", "dialout"]
```

Running only one procedure:

```console
user@host:~$ declmysys run actions.add-dialout
Run /home/user/Dotdecl (actions.add-dialout):
    - Actions (50)
        add user to dialout group
    Done!
```
