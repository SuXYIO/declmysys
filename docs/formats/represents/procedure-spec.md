# Procedure-Spec

A procedure is one of `actions` `dots` `packages`, or a subfile / subdirectory name (without `.toml` for files), e.g. `dots.foobar`.

Using the subfile / subdirectory name does not require you to be in the dotdecldir, the `.` is just a seperation character.

(And yes, you guessed it, names must not contain `.`)

> Note: Not using `/` for seperation, cuz that might be misinterpreted by some shells for files.

## Example

`actions` for all actions under `actions/`

`actions.do-stuff` for action `actions/do-stuff.toml`
