A `Dec` value is a number using fixed-point math and support is
provided by the [cockroach/apd](https://github.com/cockroachdb/apd)
library. The calculator prefers working with `Dec` values whenever an
operation can use a function in this library.

Decimal values first remove [formatting characters](#formatting-characters).
Those characters are:

- Thousand separators (`,`, `_`, ` `)
- Currency symbols (`$`, `€`, `¥`)

For example, the following all parse to the same value:

<!-- test: DecParse -->

| Input                  | Stack
|------------------------|-------------
| `c 12,345.67 dec`      | `12345.67`
| `c 12_345.67 dec`      | `12345.67`
| `c '12_345.67' dec`    | `12345.67`
| `c $12,345.67 dec`     | `12345.67`
| `c 12,345$.67 dec`     | `12345.67`

