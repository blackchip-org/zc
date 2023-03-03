# unicode

<!-- eval: use unicode -->

Unicode encoding and decoding

<!-- index -->

| Operation                     | Alias | Description
|-------------------------------|-------|--------------
| [decode](#decode)             | `de`  | Decodes Unicode code points to a string
| [encode](#encode)             | `en`  | Encodes a string into Unicode code points
| [utf-8-decode](#utf-8-decode) | `de8` | Decode UTF-8 bytes into a string
| [utf-8-encode](#utf-8-encode) | `em8` | Encode a string into UTF-8 bytes


## decode

Decodes the Unicode code points on the stack into string `s`

    ( points...:Int -- s:Str )

Alias: `de`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0x35 0x34 0xb0`  | `0x35 \| 0x34 \| 0xb0`
| `decode`          | `54°`


## encode

Encodes the string `s` into Unicode code points

    ( s:Str -- points...:Int )

Alias: `en`

Example:

<!-- test: encode -->

| Input             | Stack
|-------------------|------------------
| `use prog`        | *using prog*
| `54°`             | `54°`
| `encode`          | `53 \| 52 \| 176`
| `'hex' map`       | `0x35 \| 0x34 \| 0xb0`


## utf-8-decode

Decode the UTF-8 bytes in `b` to a string.

    ( b:BigInt -- s:Str )

Alias: `de8`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0xc2b0`          | `0xc2b0`
| `utf-8-decode`    | `°`


## utf-8-encode

Encode the string `s` into UTF-8 bytes.

    ( s:Str -- encode:BigInt )

Alias: `en8`

Example:

<!-- test: encode -->

| Input             | Stack
|-------------------|------------------
| `use prog`         | *using prog*
| `'°`               | `°`
| `utf-8-encode hex` | `0xc2b0`


