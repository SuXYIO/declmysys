# Subcommands-List

List the procedure defined by order (descending priority).

## Args

```bash
declmysys list [PRIORITY]
```

- `PRIORITY`: Show the specific procedures for a certain priority

> [!WARNING]
> `go-flags` lib does not support default value for positionals,
> so program assumes `0` is not user input, and `0` means priority not specified (list all priorities)

## Example

```console
user@host:~$ declmysys list
Listing /home/user/Decl:
    (250)
        cmds[apt update]
        cmds[add flathub source to flatpak]
    (200)
        packages[apt]: zsh  git  tmux  neofetch  neovim  flatpak  thunderbird  librewolf  kitty
    (150)
        packages[flatpak-system-flathub]: com.valvesoftware.Steam com.visualstudio.code
    (100)
        stow[dotfiles]: zshrc
        stow[dotfiles]: git
        stow[dotfiles]: tmux
        stow[dotfiles]: neofetch
        stow[dotfiles]: neovim  kitty  apt-sources
        stow[dotfiles]: kitty
        stow[dotfiles]: apt-sources
    (50)
        cmds[create ~/Workspace directory]
        cmds[add user to dialout group]
```

```console
user@host:~$ declmysys list 250
Listing /home/user/Decl (priority 250):
    cmds[apt update]
    cmds[add flathub source to flatpak]
```
