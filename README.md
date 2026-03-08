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

## TODO

Basic:

- [x] Design config formats
- [ ] Basic implementation
  - [x] Cmdline args
  - [ ] Parsing stuff
  - [ ] Executing stuff

Design:

- [ ] Update docs (i mean, at least make it _readable_)
- [ ] Better UI (Use `github.com/charmbracelet/bubbletea`?)
  - [ ] Progress bar
  - [ ] Colorscheme
- [ ] QoL
  - [ ] Toml autocompletion integration (via taplo)

Implementation:

- [ ] Concurrency support
- [ ] Cross-platform support

## Note

I'm just getting started to programming, **Go** is pretty new to me, so it's also a learning project.

btw, **Rust** will NOT be used as far as I'm concerned. Yes, I tried to learn it, and I failed :sweat_smile:.
