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

Use a published binary here:

https://github.com/blackchip-org/zc/releases

or, install [go](https://go.dev/dl/) and then install the calculator with:

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

If a function does not change the stack an informational line may be printed
right above the prompt. This is notated in the table by showing the
information line in italics.

<!-- test: info -->

| Input         | Stack
|---------------|-------------
| `import rand` | *imported rand*
| `0 rand.seed` | *seed set to 0*

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
| *blank line* | Remove the first item from stack
| `redo`, `r`  | Redo the last undo
| `quit`       | Print the final stack and return to shell
| `undo`, `u`  | Undo the last line entered

## Modules and Modes

The functions of the calculator are grouped by modules. The complete list
of modules can be found in the [standard library](doc/zlib.md) reference.
Some modules are included by default when the starting the calculator
interactively and this list of modules is known as the user's prelude.

Other modules can be loaded by either the `import` or `use` statement. The
`import` statement prefixes each function with the name of the module while
the `use` statement does not. For example:

<!-- test: modules -->

| Input          | Stack
|----------------|-------------------
| `import prog`  | *imported prog*
| `255 prog.hex` | `0xff`
| `use prog`     | *using prog*
| `oct`          | `0o377`

A mode can load a set of predefined modules and execute statements to
configure the calculator for a specific use. Start the calculator with one
of these modes using the `-m` option. For example, to start configured as a
programmer's calculator:

```
zc -m prog
```

See the full list of [modes](doc/modes.md) that are available.

## Numbers

Thousand separators are ignored when parsing numbers:

<!-- test: thousands_ignored -->

| Input         | Stack
|---------------|-------------------
| `65,536 sqrt` | `256`

Use `format` or `f` to print out numbers with thousand separators:

<!-- test: format -->

| Input         | Stack
|---------------|-------------------
| `256 2 pow`   | `65536`
| `format`      | `65,536`

Set `auto-format` to `true` to apply this automatically when numbers are
placed on the stack:

<!-- test: auto-format -->

| Input              | Stack
|--------------------|-------------------
| `true auto-format` | *auto-format set to true*
| `256 2 pow`        | `65,536`

Use `int-layout` and `point` to change the thousands separator
and decimal point:

<!-- test: european_numbers -->

| Input                     | Stack
|---------------------------|-------------------
| `'.000' int-layout`       | *int-layout set to '.000'*
| `','    point`            | *point set to ','*
| `12345 10 div f`          | `1.234,5`

Currency symbols are ignored when parsing:

<!-- test: currency -->

| Input          | Stack
|----------------|-------------------
| `$1234`        | `1234`

Set `auto-currency` to true to preserve currency symbols that are a prefix
or a suffix to a number:

<!-- test: auto-currency -->

| Input                | Stack
|----------------------|-------------------
| `true auto-currency` | *auto-currency set to true*
| `$1234`              | `$1234`
| `2 mul`              | `$2468`


## Text

To use text as a value, surround it with single quotes. If the text value is
the only item on the line, an ending quote is not required. The following
computes the length, in characters, of the given text:

<!-- test: text -->

| Input           | Stack
|-----------------|---------------
| `'one thousand` | `one thousand`
| `len`           | `12`

## Macros

Let's say that you commonly have to compute a sales tax that is 5%. To
compute the sales tax on something that costs $123:

<!-- test: tax -->

| Input               | Stack
|---------------------|-------------------
| `true auto-currency`| *auto-currency set to true*
| `$123`              | `$123`
| `dup`               | `$123 \| $123`
| `0.05`              | `$123 \| $123 \| 0.05`
| `mul`               | `$123 \| $6.15`
| `add`               | `$129.15`

Repeated use of this pattern can be used with a macro:

<!-- test: tax-macro -->

| Input                    | Stack
|--------------------------|-------------------
| `true auto-currency`     | *auto-currency set to true*
| `macro tax dup 0.05 mul` | *macro 'tax' defined*
| `$123`                   | `$123`
| `tax add`                | `$129.15`

The name of `=` is reserved for your macro use in bulk operations:

<!-- test: bulk -->

| Input                     | Stack
|---------------------------|-------------------
| `use unit`                | *using unit*
| `macro = top f-c 2 round` | *macro '=' defined*
| `32 =`                    | `0`
| `68 =`                    | `20`
| `100 =`                   | `37.78`

Play a game of rock, paper, scissors:

<!-- test: rps -->

| Input                                           | Stack
|-------------------------------------------------|-------------------
| `import rand`                                   | *imported rand*
| `0 rand.seed`                                   | *seed set to 0*
| `macro = 'rock' 'paper' 'scissors' rand.choice` | *macro '=' defined*
| `=`                                             | `rock`
| `=`                                             | `paper`


## Higher order functions

The `map` function can be used to apply a function to each item on the stack.
To use this function, the top element of the stack should be the lambda
function to evaluate. Place this lambda on the stack using quotes to prevent
immediate evaluation. For example, to double all numbers on the stack:

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'2 mul`            | `1 \| 2 \| 3 \| 4 \| 5 \| 2 mul`
| `map`               | `2 \| 4 \| 6 \| 8 \| 10`

The `fold` function can be used to reduce all items in the stack to a single
value. For example, to sum all the numbers on the stack:

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'add`              | `1 \| 2 \| 3 \| 4 \| 5 \| add`
| `fold`              | `15`

See the [fn](doc/zlib/fn.md) module for more information.

## To Be Continued...

## Credits

- Fixed point math provided by https://github.com/shopspring/decimal
- CLI auto completion and history provided by https://github.com/peterh/liner
- Terminal demo created with https://github.com/faressoft/terminalizer

## License

[MIT](LICENSE)

## Feedback

Contact me at zc@blackchip.org
