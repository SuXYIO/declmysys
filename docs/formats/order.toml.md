# order.toml

Defines specific order for execution, in case that order matters.

## Default

```toml
order = [
    "dots",
    "packages",
    "actions"
]

[dots]
include = "*"
exclude = ""

[actions]
include = "*"
exclude = ""
```

## Values

- `order`: Spec the execution order for dots and actions. Order string array, (see explaination below), or `"any"` for non-ordered. The filenames not included in the array will be interpreted as non-ordered
- `include`: The files that must be included, `"*"` for all, `""` for none, can also be array of filenames or string regex
- `exclude`: The files that are excluded, `"*"` for all, `""` for none, can also be array of filenames or string regex

### Order String Array

Array containing strings of the following three types:

1. filename, for execution of file, e.g. `"dots/dot1"`, `"actions/action2`
2. `"dots"` for non-ordered execution of `dots/*`, `"actions"` for non-ordered execution of `actions/*`. If mixed with filenames, represents the rest of files in the category.
3. `"packages"` or `"packages.toml"`, for the action of installation/upgrade of system packages
   Types can be mixed.

~~god damn it, made this thing so flexible, it's gonna be hell when writing the interpreter~~
Maybe consider using priority stuff later.

## Example

```toml
# Executes in order:
#   install packages
#   execute dots with any order
#   execute actions with any order
order = [
    "dots",
    "packages",
    "actions"
]
```

```toml
# Mixed, pretty self-explainatory
order = [
    "dots/dot1",
    "action/action2",
    "dots",     # Meaning execution of the rest of the dots: dot3, dot4, dot5, ...
    "packages",
    "dots/dot2",
    "actions"   # execution of rest of the actions: action1, action3, action4, ...
]
```
