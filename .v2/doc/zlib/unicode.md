# unicode

<!-- eval: use unicode -->

Unicode encoding, decoding, and character operations

<!-- index -->

| Operation                     | Alias | Description
|-------------------------------|-------|--------------
| [decode](#decode)             | `de`  | Decodes a Unicode code point to a character
| [encode](#encode)             | `en`  | Encodes a character into Unicode code point
| [lower](#lower)               |       | Convert character to lower case
| [lower=](#lower=)             |       | Is character lower case
| [title](#title)               |       | Convert character to title case
| [title=](#title=)             |       | Is character title case
| [upper](#upper)               |       | Convert character to upper case
| [upper=](#upper=)             |       | Is character upper case
| [utf-8-decode](#utf-8-decode) | `de8` | Decode UTF-8 bytes into a string
| [utf-8-encode](#utf-8-encode) | `en8` | Encode a string into UTF-8 bytes


## decode

Decodes the Unicode code point `p` to a character

    ( p:Int -- s:Char )

Alias: `de`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0xb0`            | `0xb0`
| `decode`          | `°`


## encode

Encodes the character `c` into Unicode code point

    ( s:Char -- point:Int )

Alias: `en`

Example:

<!-- test: encode -->

| Input             | Stack
|-------------------|------------------
| `use prog`        | *using prog*
| `'°`              | `°`
| `encode`          | `176`
| `'hex' map`       | `0xb0`


## lower

Convert character `c` to lower case

    ( s:Char -- lower:Char )

Example:

<!-- test: lower -->

| Input             | Stack
|-------------------|------------------
| `'H`              | `H`
| `lower`           | `h`


## lower=

Is character `c` in lower case?

    ( c:Char -- lower:Bool )

Example:

<!-- test: lower -->

| Input             | Stack
|-------------------|------------------
| `'H' lower=`      | `false`
| `clear`           |
| `'h' lower=`      | `true`


## title

Convert character `c` to title case

    ( c:Char -- lower:Char )

Example:

<!-- test: title -->

| Input             | Stack
|-------------------|------------------
| `'`Ǆ              | Ǆ
| `title`           | ǅ


## title=

Is character `c` in title case?

    ( c:Char -- title:Bool )

Example:

<!-- test: lower -->

| Input             | Stack
|-------------------|------------------
| `'`Ǆ`'` `title=`  | `false`
| `clear`           |
| `'`ǅ`'` `title=`  | `true`
| `clear`           |
| `'`ǆ`'` `title=`  | `false`


## upper

Convert character `c` to upper case

    ( c:Char -- lower:Str )

Example:

<!-- test: upper -->

| Input             | Stack
|-------------------|------------------
| `'h`              | `h`
| `upper`           | `H`


## upper=

Is character `c` in upper case?

    ( c:Char -- lower:Bool )

Example:

<!-- test: lower -->

| Input             | Stack
|-------------------|------------------
| `'H' upper=`      | `true`
| `clear`           |
| `'h' upper=`      | `false`


## utf-8-decode

Decode the UTF-8 bytes in `b` to a string.

    ( b:BigInt -- s:Str )

Alias: `de8`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0x3534c2b0`      | `0x3534c2b0`
| `utf-8-decode`    | `54°`


## utf-8-encode

Encode the string `s` into UTF-8 bytes.

    ( s:Str -- encode:BigInt )

Alias: `en8`

Example:

<!-- test: encode -->

| Input              | Stack
|--------------------|------------------
| `use prog`         | *using prog*
| `54°`              | `54°`
| `utf-8-encode hex` | `0x3534c2b0`


