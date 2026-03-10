# Procedure-Spec

A procedure is one of `decls` `packages`, or a subfile / subdirectory name (without `.toml` for files, `.` separated), e.g. `decls.foobar`.

Using the subfile / subdirectory name does not require you to be in the decldir, the `.` is just a separation character.

(And yes, you guessed it, names must not contain `.`)

> [!NOTE]
> Not using `/` for separation, because that might be misinterpreted by some shells for files.

## Example

`decls` for all declarations under `decls/`

`decls.do-stuff` for action `decls/do-stuff.toml`
