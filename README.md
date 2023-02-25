# zc

A fun RPN calculator.

- [modes](doc/modes.md): Calculator modes
- [zlib](doc/zlib.md): Standard library
- [index](doc/index.md): Function index

![ZC Demo](demo.gif)

When I'm at a terminal prompt and I need to use a calculator, `bc` has always
been my tool of choice. I thought it would be fun to write a calculator myself
but with some items from my wish list built in. Those items are:

- A Stack based calculator. Typing in a value places it on the stack. A
function consumes values on the stack and places its results back on the
stack.
- To minimize the use of the shift key. Instead of using `+` for addition,
use `add` or `a` which is easier to type.
- Use arbitrary sized integers and fixed point math by default. `1.1 2.2 add`
should be `3.3` and not `3.3000000000000003`.
- Be more than a simple calculator. Need an external tool to lookup, compute,
or calculate? Put in the calculator as a module instead.
- Backed with a scripting language. When a function is not directly
supported by native libraries, write the implementation in the scripting
language.

This is the third iteration of this calculator and something fun to work on
when time is available. It is a bit rough at this stage but should be useful
nonetheless. Also, it will always be a bit rough--full of bugs and
inconsistencies. Features get added as I need or think of them. Bugs get fixed
or ignored as I see them. There is no grand plan beyond tinkering around
for entertainment.

## Installation

Install [go](https://go.dev/dl/).

Install the calculator with:

    go install github.com/blackchip-org/zc/cmd/zc@latest

Run the calculator with:

    zc

## Overview

Each line entered at the calculator prompt is divided into *words*. Each
word is separated by whitespace. A word can be either a:

- *value*: Starts with a numeric character, a decimal point,
or a `+` or `-` sign. Values are placed onto the stack.
- *function*: Invokes a function with the given name. Parameters are consumed
from the stack and results are placed on the stack.

If `2 3 a` is entered at the prompt, the values of `2` and `3` are placed on
the stack, the `a` function (for addition) is executed, the values are
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

The basic math functions are:

| Function        | Description |
|-----------------|-------------
| `add`, `a`, `+` | Addition
| `sub`, `s`, `-` | Subtraction
| `mul`, `m`, `*` | Multiplication
| `div`, `d`, `/` | Division

For each of these functions there are three separate names. For addition there
is:

- `a`: Easy to type without having to use the shift key
- `add`: Easy to read in scripts or documentation
- `+`: Easy to type if you have a keyboard with a number pad.

Additional basic math functions can be found in the documentation for the
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

## Commands

These commands are available when running the calculator interactively:

| Command      | Description
|--------------|------------------------------------
| *blank line* | Pop first item from stack
| `redo`, `r`  | Redo the last undo
| `quit`       | Print the final stack and return to shell
| `undo`, `u`  | Undo the last line entered

## Modes

The calculator can be started at the command line with a pre-defined mode
using the `-m` option. For example, to start configured as a programmer's calculator:

```
zc -m prog
```

See the full list of [modes](doc/modes.md) that are available.

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

or use the [plain](doc/modes.md#plain) mode.

Currency symbols are ignored when parsing but are preserved when formatting
if found at the beginning or end of a number.

<!-- test: currency -->

| Input          | Stack
|----------------|-------------------
| `$1234`        | `$1,234`
| `2 mul`        | `$2,468`

## Macros

Let's say that you commonly have to compute a sales tax that is 5%. To
compute the sales tax on something that costs $123:

<!-- test: tax -->

| Input          | Stack
|----------------|-------------------
| `$123`         | `$123`
| `dup`          | `$123 \| $123`
| `0.05`         | `$123 \| $123 \| 0.05`
| `mul`          | `$123 \| $6.15`
| `add`          | `$129.15`

Repeated use of this pattern can be used with a macro:

<!-- test: tax-macro -->

| Input                    | Stack
|--------------------------|-------------------
| `macro tax dup 0.05 mul` |
| `$123`                   | `$123`
| `tax add`                | `$129.15`

The name of `=` is reserved for your macro use in bulk operations:

<!-- test: bulk -->

| Input                     | Stack
|---------------------------|-------------------
| `use unit`                |
| `macro = top f-c 2 round` |
| `32 =`                    | `0`
| `68 =`                    | `20`
| `100 =`                   | `37.78`

Play a game of rock, paper, scissors:

<!-- test: rps -->

| Input                                           | Stack
|-------------------------------------------------|-------------------
| `import rand`                                   |
| `0 rand.seed`                                   |
| `macro = 'rock' 'paper' 'scissors' rand.choice` |
| `=`                                             | `rock`
| `=`                                             | `paper`

## To Be Continued...

## Credits

- Fixed point math provided by https://github.com/shopspring/decimal
- CLI auto completion and history provided by https://github.com/peterh/liner
- Terminal demo created with https://github.com/faressoft/terminalizer

## License

[MIT](LICENSE)

## Feedback

Contact me at zc@blackchip.org
