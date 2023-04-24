# types

Each value on the calculator stack is a string of bytes. When an operation
needs to pop a value off the stack, it must first parse the value into the
desired type. Results are then formatted to a string before being pushed
back onto the stack.

Each type has a function for parsing and formatting. A value is considered to
be of a type if it can be successfully parsed by that type's parse function.
The parse function for a floating point number can parse values such as
`6`, `6.4`, `6e4` but not `'6 2/5'` which requires a conversion function.

Table of contents:

- [BigInt](#int)
- [Complex](#complex)
- [Date](#datetime)
- [DateTime](#datetime)
- [Decimal](#decimal)
- [Duration](#duration)
- [Int](#integer), [Int64](#integer), [Int32](#integer), [Int16](#integer), [Int8](#integer)
- [Real](#real)
- [Num](#num)
- [Str](#strval)
- [Time](#datetime)
- [Uint](#integer), [Uint64](#integer), [Uint32](#integer), [Uint16](#integer), [Uint8](#integer)
- [Val](#strval)

## Integer

An `Integer` value is a [`Real`](#real) that can either be a:

- BigInt
- Int, Int64, Int32, Int16, Int8
- Uint, Uint64, Unit32, Uint16, Uint8

A `BigInt` is an integer of an arbitrary size and is supported by the
[math/big](https://pkg.go.dev/math/big) library. The calculator prefers working
with `BigInt`s whenever an operation can use a function in this library.

The `Int` and `Uint` series of types are signed and unsigned integers of
a specific size and are used when an underlying implementation of an
operation needs that type.


## Str/Val

The type `Str` is a string of bytes is the native type of values stored on the
stack. The parse function accepts any value and the formatting function uses
the string as-is.

The type `Val` is used when an operation doesn't depend on the type of the
value but is otherwise is the same as `Str`. It is used to notate operations
such as `swap` where the types of the values being swapped is irrelevant.

Sub-types:

- [Num](#num)

## Num

Sub-types:

- [Complex](#complex)
- [Datetime](#datetime)
- [Real](#real)

## Complex

## Real

A `Real` number is a [Num](#num) that can be either a:

- [Integer](#integer)
- [Decimal](#decimal)
- [Float](#float)
- [Rational](#rational)

Formatting characters are first removed when parsing real numbers. Those
characters are:

- Thousand separators (`','`, `'_'`, `' '`)
- Currency symbols ('`$'`, `'€'`, `'¥'`)

The following strings all parse to the same real number:

- 12,345.67
- 12_345.67
- '12 345.67'
- '$12,345.67'
- '12,345.67$'



## Bool

A boolean value that is either true or false.

An item on the stack can be parsed as a boolean if it is equal to `true`
or `false` when all characters are converted to lowercase. Operations
are defined for `true` and `false` that simply return that string.

Example:

<!-- test: types-bool -->

| Input           | Stack
|-----------------|-------------
| `true true and` | `true`
| `'FALSE' and`   | `false`
| `1 and`         | `no operation for: 'false' 1 and`

