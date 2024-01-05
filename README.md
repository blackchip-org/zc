# zc

A fun stack based calculator.

Documentation:

- [Operation Reference](doc/ops.md)
- [Operations Index](doc/index.md)
- [Type Documentation](doc/types.md)
- [Stack Notation](doc/stack.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/blackchip-org/zc/v5.svg)](https://pkg.go.dev/github.com/blackchip-org/zc/v5)

![ZC Demo](demo.gif)

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

Install [go](https://go.dev/dl/) and then install the calculator with:

    go install github.com/blackchip-org/zc/v5/cmd/zc@latest

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

If there are multiple items on the stack, they are notated by using a pipe `|`
character to separate each item. The item on the right is the top of the stack.

If an operation does not change the stack an informational line may be printed
right above the prompt. This is notated in the table by showing the
information line in italics.

<!-- test: info -->

| Input         | Stack
|---------------|-------------
| `0 rand-seed` | *seed set to 0*

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

<!-- test: distance -->

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

<!-- test: thousands_ignored -->

| Input         | Stack
|---------------|-------------------
| `65,536 sqrt` | `256`

Currency symbols are also ignored when parsing:

<!-- test: currency -->

| Input          | Stack
|----------------|-------------------
| `$1234 2 mul`  | `2468`

Integer math uses arbitrary sized values when possible:

<!-- test: bigint -->

| Input          | Stack
|----------------|-------------------
| `2 128 pow`    | `340282366920938463463374607431768211456`

Real number math uses fixed point math when possible:

<!-- test: decimal -->

| Input          | Stack
|----------------|-------------------
| `1.1 2.2 a`    | `3.3`

Numbers may have a suffix of `f` to use a floating point operation instead:

<!-- test: decimal -->

| Input          | Stack
|----------------|-------------------
| `1.1f 2.2 a`   | `3.3000000000000003`

Enter fractions in `a/b` notation:

<!-- test: fraction -->

| Input          | Stack
|----------------|-------------------
| `1/2 1/4`      | `1/2 \| 1/4`
| `add`          | `3/4`

Prefix a whole number to a fraction with either a space, an underscore, or
a hyphen:

<!-- test: whole-fraction -->

| Input          | Stack
|----------------|-------------------
| `2-1/2 3-1/4`  | `2-1/2 \| 3-1/4`
| `add`          | `5 3/4`

Enter complex numbers in `r+i` notation:

<!-- test: complex -->

| Input          | Stack
|----------------|-------------------
| `1+2i 2+3i`    | `1+2i \| 2+3i`
| `add`          | `3+5i`

## Text

Text can be used as a value by either one of two ways. If the text does not
contain any whitespace, the text can be prefixed with a slash, `/`. For
example:

<!-- test: text-slash -->

| Input          | Stack
|----------------|-------------------
| `1 2`          | `1 \| 2`
| `/add`         | `1 \| 2 \| add`

Otherwise, surround text with quotes. Single quotes, `' '`, double quotes,
`" "` or square brackets, `[ ]` can be used. If the text value is the only item
on the line, an ending quote is not required. The following computes the
length, in characters, of the given text:

<!-- test: text-quote -->

| Input           | Stack
|-----------------|---------------
| `'one thousand` | `one thousand`
| `len`           | `12`

Using brackets is convenient when nesting quoted values:

<!-- test: text-bracket -->

| Input           | Stack
|-----------------|---------------
| `1 2 [[[add]]]` | `1 \| 2 \| [[add]]`
| `eval`          | `1 \| 2 \| [add]`
| `eval`          | `1 \| 2 \| add`
| `eval`          | `3`

To use multiple lines as values (for example, when pasting the contents of the
clipboard), use `quote` with a delimiter that marks the end of the values.
Each line is considered a separate value when using `quote`. For example:

<!-- test: quote -->

| Input           | Stack
|-----------------|---------------
| `1 2 add`       | `3`
| `quote EOF`     | `3`
| `2 3 add`       | `3 \| 2 3 add`
| `EOF`           | `3 \| 2 3 add`

## Macros

Let's say that you commonly have to compute a sales tax that is 5%. To
compute the sales tax on something that costs $123:

<!-- test: tax -->

| Input               | Stack
|---------------------|-------------------
| `$123`              | `$123`
| `dup`               | `$123 \| $123`
| `0.05`              | `$123 \| $123 \| 0.05`
| `mul`               | `$123 \| 6.15`
| `add`               | `129.15`

Repeated use of this pattern can be used with a macro:

<!-- test: tax-macro -->

| Input                    | Stack
|--------------------------|-------------------
| `def tax dup 0.05 mul  ` | *macro 'tax' defined*
| `$123`                   | `$123`
| `tax add`                | `129.15`

The name of `=` is reserved for your macro use in bulk operations:

<!-- test: bulk -->

| Input                     | Stack
|---------------------------|-------------------
| `def = top f-c 2 round`   | *macro '=' defined*
| `32 =`                    | `0`
| `68 =`                    | `20`
| `100 =`                   | `37.78`

No operations start with a `.` character and can be used for macro names. Play
a game of rock, paper, scissors:

<!-- test: rps -->

| Input                                            | Stack
|--------------------------------------------------|-------------------
| `0 rand-seed`                                    | *seed set to 0*
| `def .rps 'rock' 'paper' 'scissors' rand-choice` | *macro '.rps' defined*
| `.rps`                                           | `rock`
| `.rps`                                           | `paper`

Macros can also be used to override calculator operations. Undefine the
macro by using `def` without an expression.

<!-- test: override -->

| Input          | Stack
|----------------|----------------------------
| `def pi 'Yum`  | *macro 'pi' overrides*
| `pi`           | `Yum`
| `def pi`       | *macro 'pi' undefined*
| `pi`           | `Yum \| 3.14159265358979323`

## Higher order functions

The `map` operation can be used to apply a function to each item on the stack.
To use this operation, the top element of the stack should be an expression
to evaluate. Place this expression on the stack using quotes to prevent
immediate evaluation. For example, to double all numbers on the stack:

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `[2 mul`            | `1 \| 2 \| 3 \| 4 \| 5 \| 2 mul`
| `map`               | `2 \| 4 \| 6 \| 8 \| 10`

The `fold` function can be used to reduce all items in the stack to a single
value. For example, to sum all the numbers on the stack:

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `[add`              | `1 \| 2 \| 3 \| 4 \| 5 \| add`
| `fold`              | `15`

Additional higher-order functions can be found in the [hof](doc/ops/hof.md)
reference.

## Memory

Sometimes it is convenient to store the contents of the stack to memory
which can be recalled later. Let's manually compute the population standard
deviation found in the example on the [Wikipedia](https://en.wikipedia.org/wiki/Standard_deviation) page.

The average of the all the data points must first be computed and then
deviations from that average are calculated. The data is first entered
into the calculator and then the stack is saved with the name of `data`.
The average is then computed and stored with the name of `av`. The
data points are then recalled from memory and the deviations from the
average are calculated for each value. The standard deviation is then
simply the square root of the average deviation.

<!-- test: stddev -->

| Input                     | Stack
|---------------------------|---------------------|
| `2 4 4 4 5 5 7 9`         | `2 \| 4 \| 4 \| 4 \| 5 \| 5 \| 7 \| 9`
| `/data set`               | *set*
| `average`                 | `5`
| `/av set`                 | *set*
| `clear /data get`         | `2 \| 4 \| 4 \| 4 \| 5 \| 5 \| 7 \| 9`
| `[/av get sub 2 pow] map` | `9 \| 1 \| 1 \| 1 \| 0 \| 0 \| 4 \| 16`
| `average sqrt`            | `2`

Stacks can also be placed in a general memory location with no name by
using `store` and `recall`. This location is also a stack itself with `store`
pushing the current stack to the memory stack and `recall` popping off the
top stack.

<!-- test: stddev-general -->

| Input                     | Stack
|---------------------------|---------------------|
| `2 4 4 4 5 5 7 9`         | `2 \| 4 \| 4 \| 4 \| 5 \| 5 \| 7 \| 9`
| `store`                   | *stored*
| `average`                 | `5`
| `/av set`                 | *set*
| `clear recall`            | `2 \| 4 \| 4 \| 4 \| 5 \| 5 \| 7 \| 9`
| `[/av get sub 2 pow] map` | `9 \| 1 \| 1 \| 1 \| 0 \| 0 \| 4 \| 16`
| `average sqrt`            | `2`

## Annotations

Values on the stack may have annotations to provide some additional
metadata. An annotated value ends with a hash mark `#` and the annotation
text. For example, values are given an 'inexact' annotation when a
`Rational` value cannot be exactly converted to a `Decimal`. Example:

<!-- test: anno -->

| Input               | Stack
|---------------------|---------------------|
| `2/3 dec`           | `0.6666666666666666 # inexact`

Annotations are also used to indicate the units for a result.

<!-- test: anno-haversine -->

| Input                 | Stack
|-----------------------|---------------------|
| `1 1 2 2`             | `1 \| 1 \| 2 \| 2`
| `haversine`           | `157225.4320380729 # meters`
| `si.kilo div 2 round` | `157.23`

Note that annotations are lost when values are popped off the stack.

Use the `anno` operation to annotate a value and the `no-anno` or `noa`
operation to remove an annotation:

<!-- test: add-rm-anno -->

| Input               | Stack
|---------------------|---------------------|
| `42`                | `42`
| `[the answer] anno` | `42 # the answer`
| `noa`               | `42`

To disable annotations, set the `ZC_NO_ANNO` environment variable to any value.

## Commands

These commands are available when running the calculator interactively:

| Command      | Description
|--------------|------------------------------------
| *blank line* | Remove the first item from stack
| `def`        | Define a macro
| `redo`       | Redo the last undo
| `quit`       | Print the final stack and return to shell
| `quote`      | Add each line to the stack until delimiter is found
| `undo`, `u`  | Undo the last line entered

## Command line

Any arguments found on the command line are passed to the calculator for a
one-time evaluation:

    $ zc 2 3 add
    5

Using quotes from the command line can be tricky since they are interpreted by
the shell:

    $ zc 'foo' len
    (!) unknown operation: foo

In this case, wrap the expression with quotes and then use square brackets
for foo:

    $ zc '[foo] len'
    3

Use a single argument of `-` to read from standard input:

    $ echo "2 3 add" | zc -
    5

## External Libraries

Some features of the calculator use external C libraries. The current
release binaries do not include these external libraries and using one
of these operations will raise a "feature not supported" error.

To use these features, install the necessary dependencies and build
locally using `make`. See the credits below for more information.

## Credits

- Fixed point math provided by https://github.com/shopspring/decimal
- CLI auto completion and history provided by https://github.com/peterh/liner
- Geospatial transformations provided by https://github.com/twpayne/go-proj/
- Emoji JSON file provided by https://github.com/muan/unicode-emoji-json
- Terminal demo created with https://github.com/faressoft/terminalizer

## License

[MIT](LICENSE)

## Feedback

Contact me at zc@blackchip.org

