# text

Text operations

<!-- index -->

| Operation        | Alias     | Description
|------------------|-----------|----------------
| [char-codepoint](#char-codepoint) | `char-cp` | Convert character to code point
| [codepoint-char](#codepoint-char) | `cp-char` | Convert code point to character
| [is](#is)        |           | True if strings are equal
| [join](#join)    |           | Join stack elements into a single string
| [left](#left)    |           | Substring from the left
| [len](#len)      |           | Length
| [lower](#lower)  |           | Converts string into lowercase
| [right](#right)  |           | Substring from the right
| [split](#split)  |           | Splits a string by a separator
| [upper](#upper)  |           | Converts string into uppercase
| [utf-8-decode](#utf-8-decode) | `u8de` | Decode UTF-8 bytes into a string
| [utf-8-encode](#utf-8-encode) | `u8en` | Encode a string into UTF-8 bytes


## char-codepoint

Convert character into an integer code point.

    ( Char -- Int32 )

<!-- test: char-codepoint -->

| Input             | Stack
|-------------------|------------------
| `'°`              | `°`
| `char-cp`         | `176`
| `hex`             | `0xb0`

## codepoint-char

Convert code point to character.

    ( Int32 -- Char )

<!-- test: codepoint-char -->

| Input             | Stack
|-------------------|------------------
| `0xb0`            | `0xb0`
| `cp-char`         | `°`

## is

True if *a0* and *a1* are the same, character by character.

    ( a0:String a1:String -- Bool )

Example:

<!-- test: is -->

| Input        | Stack
|--------------|------------------|
| `1.2 1.20`   | `1.2 \| 1.20`
| `is`         | `false`
| `clear`      |
| `1.2 1.2`    |  `1.2 \| 1.2`
| `is`         | `true`

## join

Join all stack elements into a single string separated by *sep*.

    ( Val* sep:Str -- Str )

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

Substring of *s* from the left.

If *n* is positive, *m* characters are taken from the left. If *n* is negative,
characters are taken from the left until there are *n* characters remaining. If
*n* is zero, *s* is returned without change.

If the absolute value of *n* is greater then then length of *s*, an
'illegal arguments' error is raised.

    ( s:String n:Int -- String )

Example:

<!-- test: left -->

| Input        | Stack
|--------------|------------------|
| `'abcdef`    | `abcdef`
| `4 left`     | `abcd`
| `-1 left`    | `abc`

## len

Length of string in characters.

    ( Str -- Int )

Example:

<!-- test: len -->

| Input        | Stack
|--------------|------------------|
| `'abcd`      | `abcd`
| `len`        | `4`

## lower

Converts the string to lowercase.

    ( Str -- Str )

Example:

<!-- test: lower -->

| Input        | Stack
|--------------|------------------|
| `'AbCd`      | `AbCd`
| `lower`      | `abcd`

## right

Substring of *s* from the right.

If *n* is positive, *n* characters are taken from the right. If *n* is
negative, characters are taken from the right until there are *n* characters
remaining. If *n* is zero, *s* is returned without change.

If the absolute value of *n* is greater then then length of *s*, an
'illegal arguments' error is raised.

    ( s:String n:Int -- right:String )

Example:

<!-- test: right -->

| Input        | Stack
|--------------|------------------|
| `'abcdef`    | `abcdef`
| `4 right`    | `cdef`
| `-1 right`   | `def`


## split

Split `p0` into multiple strings that are separated by `sep`.

    ( p0:Str sep:Str -- Str* )

Example:

<!-- test: split -->

| Input        | Stack
|--------------|------------------|
| `128.8.74.2` | `128.8.74.2`
| `'.' split`  | `128 \| 8 \| 74 \| 2`
| `clear`
| `1234`       | `1234`
| `'' split`   | `1 \| 2 \| 3 \| 4`

## upper

Converts a string to uppercase.

    ( Str -- Str )

Example:

<!-- test: upper -->

| Input        | Stack
|--------------|------------------|
| `'AbCd`      | `AbCd`
| `upper`      | `ABCD`

## utf-8-decode

Decode UTF-8 bytes into to a string.

    ( BigInt -- Str )

Alias: `u8de`

Example:

<!-- test: decode -->

| Input             | Stack
|-------------------|------------------
| `0x3534c2b0`      | `0x3534c2b0`
| `utf-8-decode`    | `54°`

## utf-8-encode

Encode a string into UTF-8 bytes.

    ( Str -- BigInt )

Alias: `u8en`

Example:

<!-- test: encode -->

| Input              | Stack
|--------------------|------------------
| `54°`              | `54°`
| `utf-8-encode hex` | `0x3534c2b0`
