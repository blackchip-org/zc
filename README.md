# zc

A fun RPN calculator.

- [zlib](doc/zlib.md): Standard library
- [index](doc/index.md): Operation index

When I'm at a terminal prompt and I need to use a calculator, `bc` has always
been my tool of choice. I thought it would be fun to write a calculator myself
but with some items from my wish list built in. Those items are:

- A Stack based calculator. Typing in a value places it on the stack. An
operation consumes values on the stack and places its results back on the
stack.
- To minimize the use of the shift key. Instead of using `+` for addition,
use `add` or `a` which is easier to type.
- Use arbitrary sized integers and fixed point math by default. `1.1 2.2 add`
should be `3.3` and not `3.3000000000000003`.
- Be more than a simple calculator. Need an external tool to lookup, compute,
or calculate? Put in the calculator as a module instead.
- Backed with a scripting language. When an operation is not directly
supported by native libraries, write the implementation in the scripting
language.

This is the third iteration of this calculator and something fun to work on
when time is available. It is a bit rough at this stage but should be useful
nonetheless.

## Installation

Install [go](https://go.dev/dl/).

Install the calculator with:

    go get github.com/blackchip-org/zc

Run the calculator with:

    zc

## Overview

Each line entered at the calculator prompt is divided into *words*. Each
word is separated by whitespace. A word can be either a:

- *value*: Starts with a numeric character, a decimal point,
or a `+` or `-` sign. Values are placed onto the stack.
- *operation*: Invokes a function with the given name. Parameters are consumed
from the stack and results are placed on the stack.

If `2 3 a` is entered at the prompt, the values of `2` and `3` are placed on
the stack, the `a` operation (for addition) is executed, the values are
consumed and the result `5` is placed on the stack.

Examples of calculator use will be presented in a table such as:

<!-- test: simple_addition -->

| Input   | Stack
|---------|-------------
| `2 3 a` | `5`

The *Input* column shows the text entered at the prompt and the *Stack* column
shows the contents of the stack after the line is evaluated. Each word could
have been placed on a separate line:

<!-- test: simple_addition_2 -->

| Input   | Stack
|---------|-------------
| `2`     | `2`
| `3`     | `2 \| 3`
| `a`     | `5`

If there a multiple items on the stack, they are notated by using a pipe `|`
character to separate each item. The item on the right is the top of the stack.

The basic math operations are:

| Operation       | Description |
|-----------------|-------------
| `add`, `a`, `+` | Addition
| `sub`, `s`, `-` | Subtraction
| `mul`, `m`, `*` | Multiplication
| `div`, `d`, `/` | Division

For each of these operations there are three separate names. For addition there
is:

- `a`: Easy to type without having to use the shift key
- `add`: Easy to read in scripts or documentation
- `+`: Easy to type if you have a keyboard with a number pad.

Additional basic math operations can be found in the documentation for the
[math](doc/zlib/math.md) module.

## Example

Let's compute the distance between two points: `(2, 3)` and `(5, 7)`. The
formula for this uses the Pythagorean theorem:

    dist = sqrt((x2 - x1)^2 + (y2 - y1)^2)

The steps are:

- Compute `x2 - x1`
- Square the result
- Compute `y2 - y1`
- Square the result
- Add them together
- Take the square root

The entry into the calculator looks like the following:

<!-- test: distance -->

| Input     | Stack
|-----------|-------------------
| `5 2 sub` | `3`
| `2 pow`   | `9`
| `7 3 sub` | `9 \| 4`
| `2 pow`   | `9 \| 16`
| `add`     | `25`
| `sqrt`    | `5`

## Text

To use text as a value, surround it with single quotes. If the text value is
the only item on the line, an ending quote is not required. The following
computes the length, in characters, of the given text:

<!-- test: text -->

| Input           | Stack
|-----------------|---------------
| `'one thousand` | `one thousand`
| `len`           | `12`

## Numbers

Thousands separators are ignored when parsing numbers and added by default
when formatting numbers:

<!-- test: thousands_separator -->

| Input         | Stack
|---------------|-------------------
| `65,536 sqrt` | `256`
| `2 pow`       | `65,536`

Use `conf.int-format` and `conf.point` to change the thousands separator
and decimal point:

<!-- test: european_numbers -->

| Input                     | Stack
|---------------------------|-------------------
| `'.000' conf.int-format`  |
| `','    conf.point`       |
| `12345 10 div`            | `1.234,5`

Disable thousands separators with:

<!-- test: disable_thousands_separator -->

| Input                     | Stack
|---------------------------|-------------------
| `'0' conf.int-format`     |
| `256 2 pow`               | `65536`



## External Dependencies

- Fixed point math provided by https://github.com/shopspring/decimal
- CLI auto completion and history provided by https://github.com/peterh/liner

## License

[MIT](LICENSE)

## Feedback

Contact me at zc@blackchip.org
