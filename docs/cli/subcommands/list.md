# Subcommands-List

List the procedure defined by order (descending priority).

## Args

```bash
declmysys list [PRIORITY]
```

- `PRIORITY`: Show the specific procedures for a certain priority, shows them more verbosely by default

## Example

```console
user@host:~$ declmysys list
Procedure list for /home/user/Decl:
    (250)
        cmds[apt update]
        cmds[add flathub source to flatpak]
    (200)
        packages[apt]: zsh  git  tmux  neofetch  neovim  flatpak  thunderbird  librewolf  kitty
    (150)
        packages[flatpak-system-flathub]: com.valvesoftware.Steam com.visualstudio.code
    (100)
        stow[dotfiles]: zshrc  git  tmux  neofetch  neovim  kitty  apt-sources
    (50)
        cmds[create ~/Workspace directory]
        cmds[add user to dialout group]
```

```console
user@host:~$ declmysys list 250
Procedure list for /home/user/Decl (priority 250):
    cmds[apt update]
        run: ["sudo", "apt", "update"]
    cmds[add flathub source to flatpak]
        run: ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
```
