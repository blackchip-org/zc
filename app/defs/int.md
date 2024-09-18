An integer value is a number that can either be a:

- `Int`; or
- `Int/s`, `Int/s64`, `Int/s32`, `Int/s16`, `Int/s8`; or
- `Int/u`, `Int/u64`, `Int/u32`, `Int/u16`, `Int/u8`

A `Int` is an integer of an arbitrary size and support is provided by the
math/big package in the go standard library. The calculator prefers working
with `Int` values whenever an operation can use a function in this library.

The `Int/s` and `Int/u` series of types are signed and unsigned integers of a
specific size and are used when an underlying implementation of an operation
needs that type.

Integer values first remove formatting characters when parsing. Those
characters are:

- Thousand separators (`,`, `_`, ` `)
- Currency symbols (`$`, `€`, `¥`)

For example, the following all parse to the same value:

<!-- test: IntParse -->

| Input               | Stack
|---------------------|-------------
| `c 12,345 dec`      | `12345`
| `c 12_345 dec`      | `12345`
| `c '12_345' dec`    | `12345`
| `c $12,345 dec`     | `12345`
| `c 12,345$ dec`     | `12345`

Integers may have a radix prefix of:

- `0b`: binary number, base 2
- `0o`: octal number, base 8
- `0x`: hexadecimal number, base 16

For example, all of the following are the same value:

<!-- test: IntRadix -->

| Input               | Stack
|---------------------|-------------
| `c 0b11111111 dec`  | `255`
| `c 0o377 dec`       | `255`
| `c 0xff dec`        | `255`
