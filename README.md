<p align="center">
  <img width="768" height="256" alt="DeclmysysBanner" src="https://github.com/user-attachments/assets/01d030b8-7f00-42fb-8457-727bcec973d4" />
</p>

# DeclMySys

## Intro

### What

DeclMySys (_Declare My System_), the **simple**, **declarative**, system config manager.

The **intuitive** way for us **geeks**.

The config manager, that is more than _stow_, less than _nix_.

### Why

When it comes to managing system configs, you are basically left with a few options:

1. **Writing Scripts By Hand**
   Well if you really enjoy this and are good at this, then I admire you.
2. **Stow**
   Simple, intuitive, the old-school manager, but lacks features we need. (If stow can run scripts, I'll be sticking to it. The simplicity is really a beauty, made with Unix Philosophy and all)
3. **YADM & Other Modern Dotfile Managers**
   While they implemented encryption and stuff, I don't really enjoy them. Not denying their work, but in my opinion, some are overly simplified while others are overly complicated. What's more, they're not really declarative.
4. **Nix**
   Really powerful, yet complicated.
   Don't get me wrong, _Nix_ did a great job, but it complicated things for sure. It's like using a high-precision laser just for baking bread, too much for us who just tinkers around our configs.
   After all, _Nix_ is made for _strict reproducibility_. It's mostly a **production environment** thing.

I wish to build a manager, not only limited to dotfiles, but can also describe any detail of your system (via scripts).

I'm afraid it won't be as intuitive as **Stow**, but I promise using it will be easy.

## Requirements

- OS: Linux
- Architecture: amd64

Cross-platform support is not guaranteed, but works on any Linux distro and any architecture theoretically.

> [!NOTE]
> All manual tests are done in an Alpine Virtual Machine (Alpine 3.23.3 Virtual `Linux 6.18.9-0-virt x86_64` on VirtualBox)

## Install

Dependencies:

- `git`

Recommended:

- `taplo` for toml completion and validation
- `stow` for using the stow preset

> [!NOTE]
> For those of you who don't know, Taplo is an excellent linter and formatter for TOML (no this is not an ad).
> I wrote the schemas files (under `.schemas`) and the `.taplo.toml` files under every decldir, so writing decls will be hopefully easier.

### Go-install

Run:

```bash
go install github.com/suxyio/declmysys
```

And the executable will be put to `GOPATH/bin`.
If you do not know where `GOPATH` is (usually `~/go`), try `go env GOPATH`.

### Manual build

> [!TIP]
> Try `git clone --branch=main --depth=1 https://github.com/suxyio/declmysys.git` if network is slow.

Requires `go` cli tool

```bash
git clone https://github.com/suxyio/declmysys.git
cd declmysys
go build
```

The result is `./declmysys` binary.

Add custom build options for your own need.

### Uninstall

If you installed it via go-install, remove the binary under `GOPATH/bin` (usually `~/go/bin`).

The program also asks to create config file under default config dir (usually `~/.config/declmysys`), you can also remove that directory.

## Usage

See docs under `./docs`.

I'll improve the documentation later, sorry for the inconvenience.

## TODO

Basic:

- [x] Design config formats
- [x] Basic implementation
  - [x] Cmdline args
  - [x] Parsing stuff
  - [x] Executing stuff

Design:

- [ ] Update docs (i mean, at least make it _readable_)
  - [ ] Write quickstart and stuff
  - [ ] Add docs site via github pages (try docsify?)
- [x] Use better cli framework (try `github.com/urfave/cli`?)
- [ ] Better UI (Use `github.com/charmbracelet/bubbletea`?)
  - [ ] Progress bar
  - [ ] Colorscheme
- [ ] Logic
  - [ ] Ablility to specify which decl to run
  - [ ] Better error messages (show which operation went wrong)
  - [x] Less nested file structure
- [ ] Features / QoL
  - [x] Toml autocompletion integration (via taplo)
  - [ ] Shell completion integration
  - [x] Dry run for `run` subcmd
  - [ ] Pre-install command for `packages` decl (e.g. `sudo apt update`)
  - [ ] Locales

Implementation:

- [ ] Concurrency support
- [ ] Cross-platform support

## Note

I'm just getting started to programming, **Go** is pretty new to me, so it's also a learning project.

btw, **Rust** will NOT be used as far as I'm concerned. Yes, I tried to learn it, and I failed :sweat_smile:.
