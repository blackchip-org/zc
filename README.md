# zc

A fun stack-based calculator.

Use the browser version here:

TODO

Use tab to auto-complete. First tab completes an operation name as much as
possible. Next tab shows matching candidates. When using on a mobile device, use
the "?" button to emulate pressing the tab button twice.

Example use:

TODO

## Documentation

- [Operation Reference](doc/ops.md)
- [Operations Index](doc/index.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/blackchip-org/zc/v6.svg)](https://pkg.go.dev/github.com/blackchip-org/zc/v6)

![ZC Demo](demo.gif)

## About

When I'm at a terminal prompt and I need to use a calculator, `bc` has always
been my tool of choice. I thought it would be fun to write a calculator myself
but with some items from my wish list built in. Those items are:

- A stack-based calculator. Typing in a value places it on the stack. An
operation consumes values on the stack and places its results back on the
stack.
- To minimize the use of the shift key. Instead of using `+` for addition,
use `add` or `a` which is easier to type.
- Use arbitrary sized integers and fixed point math by default. `1.1 2.2 add`
should be `3.3` and not `3.3000000000000003`.
- Be more than a simple calculator. Need an external tool to lookup, compute,
or calculate? Put it in the calculator as a module instead. Make this
calculator like a Swiss army knife.
- Auto-complete!

This is the third iteration of this calculator and something fun to work on
when time is available. It is a bit rough at this stage but should be useful
nonetheless. Also, it will always be a bit rough--full of bugs and
inconsistencies. Features get added as I need or think of them. Bugs get fixed
or ignored as I see them. There is no grand plan beyond tinkering around for
entertainment. Things may change in backwards incompatible ways with no notice.

## Installation

For the command-line version, install [go](https://go.dev/dl/) and then install
the calculator with:

    go install github.com/blackchip-org/zc/v6/cmd/zc@latest

Run the calculator with:

    zc

## Overview

Each line entered at the calculator prompt is divided into *words*. Each
word is separated by whitespace. A word can be either a:

- *value*: Starts with a numeric character, a decimal point,
a numeric sign, or quotes. Values are placed onto the stack.
- *operation*: Invokes a operation with the given name. Parameters are consumed
from the stack and results are placed on the stack.

If `2 3 a` is entered at the prompt, the values of `2` and `3` are placed on
the stack, the `a` operation (for addition) is executed, the values are
consumed and the result `5` is placed on the stack.

Examples of calculator use will be presented in a table such as:

<!-- test: SimpleAddition -->

| Input   | Stack
|---------|-------------
| `2 3 a` | `5`

The *Input* column shows the text entered at the prompt and the *Stack* column
shows the contents of the stack after the line is evaluated. Each word could
have been placed on a separate line:

<!-- test: SimpleAddition2 -->

| Input   | Stack
|---------|-------------
| `2`     | `2`
| `3`     | `2 \| 3`
| `a`     | `5`

If there are multiple items on the stack, they are notated by using a pipe `|`
character to separate each item. The item on the right is the top of the stack.

If an operation does not change the stack a notification may be printed
right above the prompt. This is indicated in the table by using italics.

<!-- test: Notice -->

| Input         | Stack
|---------------|-------------
| `0 rand.seed` | *seed set to 0*

The basic math functions are:

| Function        | Description |
|-----------------|-------------
| `add`, `a`, `+` | Addition
| `sub`, `s`, `-` | Subtraction
| `mul`, `m`, `*` | Multiplication
| `div`, `d`, `/` | Division

For each of these operations there are three separate names. For addition there
is:

- `a`: Easy to type without having to use the shift key
- `add`: Easy to read in documentation
- `+`: Easy to type if you have a keyboard with a number pad

Additional basic math functions can be found in the [basic](doc/ops/basic.md)
reference.

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

<!-- test: Distance -->

| Input     | Stack
|-----------|-------------------
| `5 2 sub` | `3`
| `2 pow`   | `9`
| `7 3 sub` | `9 \| 4`
| `2 pow`   | `9 \| 16`
| `add`     | `25`
| `sqrt`    | `5`

## Numbers

Thousand separators are ignored when parsing numbers:

<!-- test: ThousandsIgnored -->

| Input         | Stack
|---------------|-------------------
| `65,536 sqrt` | `256`

Currency symbols are also ignored when parsing:

<!-- test: Currency -->

| Input          | Stack
|----------------|-------------------
| `$1234 2 mul`  | `2468`

Integer math uses arbitrary sized values when possible:

<!-- test: BigInt -->

| Input          | Stack
|----------------|-------------------
| `2 128 pow`    | `340282366920938463463374607431768211456`

Real number math uses fixed point math when possible:

<!-- test: Decimal -->

| Input          | Stack
|----------------|-------------------
| `1.1 2.2 a`    | `3.3`

To perform a 128-bit floating-point operation, use `add/f` instead:

<!-- test: Float -->

| Input              | Stack
|--------------------|-------------------
| `1.1 2.2 add/f`    | `3.3000000000000000000000000000000002`

Use either `round` or `r` to round to a certain number of digits after the
decimal point:

<!-- test: Round -->

| Input              | Stack
|--------------------|-------------------
| `1.1 2.2 add/f`    | `3.3000000000000000000000000000000002`
| `2 round`          | `3.3`

Enter fractions in `a/b` notation:

<!-- test: Fraction -->

| Input          | Stack
|----------------|-------------------
| `1/2 1/4`      | `1/2 \| 1/4`
| `add`          | `3/4`

Prefix a whole number to a fraction with either a space, an underscore, or
a hyphen:

<!-- test: WholeFraction -->

| Input          | Stack
|----------------|-------------------
| `2-1/2 3-1/4`  | `2-1/2 \| 3-1/4`
| `add`          | `5 3/4`

Enter complex numbers in `r+i` notation:

<!-- test: Complex -->

| Input          | Stack
|----------------|-------------------
| `1+2i 2+3i`    | `1+2i \| 2+3i`
| `add`          | `3+5i`

## Text

Text can be used as a value by either one of two ways. If the text does not
contain any whitespace, the text can be prefixed with a slash, `/`. For
example:

<!-- test: TextSlash -->

| Input          | Stack
|----------------|-------------------
| `1 2`          | `1 \| 2`
| `/add`         | `1 \| 2 \| add`

Otherwise, surround text with quotes. Single quotes, `' '`, double quotes,
`" "` or square brackets, `[ ]` can be used. If the text value is the only item
on the line, an ending quote is not required. The following computes the
length, in characters, of the given text:

<!-- test: TextQuote -->

| Input           | Stack
|-----------------|---------------
| `'one thousand` | `one thousand`
| `len`           | `12`

Using brackets is convenient when nesting quoted values:

<!-- test: TextBracket -->

| Input           | Stack
|-----------------|---------------
| `1 2 [[[add]]]` | `1 \| 2 \| [[add]]`
| `eval`          | `1 \| 2 \| [add]`
| `eval`          | `1 \| 2 \| add`
| `eval`          | `3`

To use multiple lines as values (for example, when pasting the contents of the
clipboard), use `quote` with a delimiter that marks the end of the values.
Each line is considered a separate value when using `quote`. For example:

<!-- test: Quote -->

| Input           | Stack
|-----------------|---------------
| `1 2 add`       | `3`
| `quote EOF`     | `3`
| `2 3 add`       | `3 \| 2 3 add`
| `EOF`           | `3 \| 2 3 add`


## TODO

## License

[MIT](LICENSE)

## Feedback

Contact me at zc@blackchip.org
