# Priority

Priority is described via an `int`. The higher the priority value, the earlier it gets executed, vice versa. Sharing priority value defines order doesn't matter.

The priority value is shared among `packages.toml`, `dots/`, `actions/` (i.e. tasks sorted by priority together), which their default priority values are: `200`, `100`, `50`.

> [!NOTE]
> Designed an `order` model at first, which can spec order via dir name and file name, and even a combination of both. But eventually, I thought it was too hard to implement. Thanks to the _Optimization Principle_ in Unix Philosophy, I abandoned that. Maybe later.
