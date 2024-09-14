# Types

Each value on the calculator stack is a treated as a text value. When an
operation needs to pop a value off the stack, it must first parse the value into
the desired type. Results are then formatted to a string before being pushed
back onto the stack.

Each type has a function for parsing and formatting. A value is considered to be
of a type if it can be successfully parsed by that type's parse function. The
parse function for a floating point number can parse values such as 6, 6.4, 6e4
but not '6 2/5' which requires a conversion function.

Table of contents:

- [`Int`](ops/int.md)
  - [`Int/s`](ops/int.md)
  - [`Int/s8`](ops/int.md)
  - [`Int/s16`](ops/int.md)
  - [`Int/s32`](ops/int.md)
  - [`Int/s64`](ops/int.md)
  - [`Int/u`](ops/int.md)
  - [`Int/u8`](ops/int.md)
  - [`Int/u16`](ops/int.md)
  - [`Int/u32`](ops/int.md)
  - [`Int/u64`](ops/int.md)
