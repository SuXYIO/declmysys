# Priority

Priority is described via an `uint` (unsigned int / positive int).
The higher the priority value, the earlier it gets executed, vice versa.
Sharing priority value defines that the operations run with random order.

The default priority value for decls is 100.

> [!NOTE]
> Designed an `order` model at first, which can spec order via dir name and file name, and even a combination of both. But eventually, I thought it was too hard to implement. Thanks to the _Optimization Principle_ in Unix Philosophy, I abandoned that. Maybe later.

> [!NOTE]
> At first I designed same priorities as running concurrently, but it turns out to be vulnerable, especially if commands ask for input, so not gonna do that until some solution is out.
