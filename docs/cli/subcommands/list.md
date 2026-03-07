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
Procedure list for /home/user/Dotdecl:
    - Actions (250)
        apt update
        add flathub source to flatpak
    - Packages (200)
        [apt]: zsh  git  tmux  neofetch  neovim  flatpak  thunderbird  librewolf  kitty
    - Packages (150)
        [flatpak-system-flathub]: com.valvesoftware.Steam com.visualstudio.code
    - Dots (100)
        zshrc  git  tmux  neofetch  neovim  kitty  apt-sources
    - Actions (50)
        create ~/Workspace directory
        add user to dialout group
```

```console
user@host:~$ declmysys list 250
Procedure list for /home/user/Dotdecl, priority 250:
    - Actions
        apt update
            run: ["sudo", "apt", "update"]
        add flathub source to flatpak
            run: ["flatpak", "remote-add", "--if-not-exists", "flathub", "https://dl.flathub.org/repo/flathub.flatpakrepo"]
```
