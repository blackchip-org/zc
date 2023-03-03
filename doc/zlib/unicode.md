# unicode

<!-- eval: use unicode -->

Unicode encoding and decoding

<!-- index -->

| Operation                    | Description
|------------------------------|-----------------------
| [code-point](#code-point)    | Converts a string into Unicode code points
| [utf8-encode]](#utf8-encode) | Encode a UTF-8 string into bytes
| [utf8-decode](#utf8-decode)  | Decode bytes into a UTF-8 string


## code-point

Converts the string `s` into Unicode code points

    ( s:Str -- points...:Int )

Example:

<!-- test: code-point -->

| Input             | Stack
|-------------------|------------------
| `use prog`        | *using prog*
| `54°`             | `54°`
| `code-point`      | `53 \| 52 \| 176`

## utf-encode

Encode the string `s` into UTF-8 bytes.

    ( s:Str -- encode:BigInt )

Example:

<!-- test: encode -->

| Input             | Stack
|-------------------|------------------
| `use prog`        | *using prog*
| `'°`              | `°`
| `utf8-encode hex` | `0xc2b0`


## utf8-decode

Decode the UTF-8 bytes in `b` to a string.

    ( b:BigInt -- s:Str )

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0xc2b0`          | `0xc2b0`
| `utf8-decode`     | `°`

