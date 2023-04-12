# str

<!-- eval: use str -->

String operations

    use str -- dev prelude

<!-- index -->

| Operation                   | Description
|-----------------------------|----------------
| [join](#join)               | Join stack elements into a single string
| [left](#left)               | Substring from the left
| [len](#len)                 | Length
| [right](#right)             | Substring from the right
| [split](#split)             | Splits a string by a separator
| [starts-with](#starts-with) | Starts with


## join

Join all stack elements into a single string separated by `sep`.

    ( items...:Str sep:Str -- join:Str )

Example:

<!-- test: join -->

| Input        | Stack
|--------------|------------------|
| `128 8 74 2` | `128 \| 8 \| 74 \| 2`
| `'.' join`   | `128.8.74.2`
| `clear`      |
| `1 2 3 4`    |  `1 \| 2 \| 3 \| 4`
| `'' join`    | `1234`


## left

Substring of `s` from the left.

If `i` is positive, `i` characters are taken from the left. If `i` is negative,
characters are taken from the left until there are `i` characters remaining. If
`i` is zero, `s` is returned without change.

If the absolute value of `i` is greater then then length of `s`, an error
is raised.

    ( s:String i:Int -- left:String )

Example:

<!-- test: left -->

| Input        | Stack
|--------------|------------------|
| `'abcdef`    | `abcdef`
| `4 left`     | `abcd`
| `-1 left`    | `abc`


## len

Length of string `a`.

    ( a:Str -- len:Int )

Example:

<!-- test: len -->

| Input        | Stack
|--------------|------------------|
| `'abcd`      | `abcd`
| `len`        | `4`


## right

Substring of `s` from the right.

If `i` is positive, `i` characters are taken from the right. If `i` is
negative, characters are taken from the right until there are `i` characters
remaining. If `i` is zero, `s` is returned without change.

If the absolute value of `i` is greater then then length of `s`, an error
is raised.

    ( s:String i:Int -- right:String )

Example:

<!-- test: right -->

| Input        | Stack
|--------------|------------------|
| `'abcdef`    | `abcdef`
| `4 right`    | `cdef`
| `-1 right`   | `def`


## split

Split `s` into multiple strings that are separated by `sep`.

    ( s:Str sep:Str -- split...:Str )

Example:

<!-- test: split -->

| Input        | Stack
|--------------|------------------|
| `128.8.74.2` | `128.8.74.2`
| `'.' split`  | `128 \| 8 \| 74 \| 2`
| `clear`
| `1234`       | `1234`
| `'' split`   | `1 \| 2 \| 3 \| 4`


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
