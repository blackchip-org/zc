# unicode

<!-- eval: use unicode -->

Unicode encoding and decoding

<!-- index -->

| Operation                    | Alias | Description
|------------------------------|-------|--------------
| [code-point](#code-point)    | `cp`  | Converts a string into Unicode code points
| [utf8-encode]](#utf8-encode) | `u8e` | Encode a UTF-8 string into bytes
| [utf8-decode](#utf8-decode)  | `u8d` | Decode bytes into a UTF-8 string


## code-points

Converts the string `s` into Unicode code points

    ( s:Str -- points...:Int )

Alias: `cp`

Example:

<!-- test: code-point -->

| Input             | Stack
|-------------------|------------------
| `use prog`        | *using prog*
| `54°`             | `54°`
| `code-points`     | `53 \| 52 \| 176`
| `'hex' map`       | `0x35 \| 0x34 \| 0xb0`

## utf-8-encode

Encode the string `s` into UTF-8 bytes.

    ( s:Str -- encode:BigInt )

Alias: `u8e`

Example:

<!-- test: encode -->

| Input             | Stack
|-------------------|------------------
| `use prog`         | *using prog*
| `'°`               | `°`
| `utf-8-encode hex` | `0xc2b0`


## utf-8-decode

Decode the UTF-8 bytes in `b` to a string.

    ( b:BigInt -- s:Str )

Alias: `u8d`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0xc2b0`          | `0xc2b0`
| `utf-8-decode`    | `°`

