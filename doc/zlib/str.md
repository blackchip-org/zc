<!-- use: str -->

# str

String operations

- Prelude: dev (include)

<!-- index -->

| Operation                   | Description
|-----------------------------|----------------
| [len](#len)                 | Length
| [starts-with](#starts-with) | Starts with


## len

Length of string `a`.

    ( a:Str -- len:Int )

Example:

<!-- test: len -->

| Input        | Stack
|--------------|------------------|
| `'abcd`      | `abcd`
| `len`        | `4`


## starts-with

`true` if string `a` starts with `b`, otherwise `false`

    ( a:Str b:Str -- starts-with:Bool )

Example:

<!-- test: starts-with -->

| Input               | Stack
|---------------------|------------------|
| `'foobar`           | `foobar`
| `'foo' starts-with` | `true`
| `clear`             |
| `'foobar`           | `foobar`
| `'bar' starts-with` | `false`
