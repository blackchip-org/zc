# zlib

The zc standard library.

## Default Modules

The following modules are loaded by default by the user's prelude:

| Name                           | Description
|--------------------------------|--------------------------------------
| [bool](zlib/bool.md)           | Boolean operations
| [format](zlib/format.md)       | Numeric formatting
| [fn](zlib/fn.md)               | Higher order functions
| [math](zlib/math.md)           | Basic mathematical operations
| [stack](zlib/stack.md)         | Stack manipulations
| [zc](zlib/zc.md)               | Information about the calculator

## User Modules

These modules can be imported or used to extend the features of the
calculator:

| Name                           | Description
|--------------------------------|----------------------------------------
| [dice](zlib/dice.md)           | Dice roller
| [prog](zlib/prog.md)           | Programmer's calculator
| [rand](zlib/rand.md)           | Random numbers
| [rot](zlib/rot.md)             | Rotation ciphers
| [si](zlib/si.md)               | Prefixes for the International System of Units
| [str](zlib/str.md)             | String operations
| [time](zlib/time.md)           | Date, time, and duration operations
| [tz](zlib/tz.md)               | Time zone database
| [unit](zlib/unit.md)           | Units of measure
| [unicode](zlib/unicode.md)     | Unicode encoding and decoding

## Developer's Modules

These modules are useful when developing calculator scripts:

| Name                           | Description
|--------------------------------|----------------------------------------
| [assert](zlib/assert.md)       | Type assertions
| dev                            | Development functions commonly used
| [io](zlib/io.md)               | Input/output functions
| runtime                        |
| test                           |

### Submodules

These modules are automatically in use through overloaded functions. These
never need to be imported but can be for testing purposes.

| Name                               | Description
|------------------------------------|----------------------------------------
| [bool.bigint](zlib/bool-bigint.md) | Boolean operations using integer math
| [bool.fixed](zlib/bool-fixed.md)   | Boolean operations using fixed-point math
| [math.bigint](zlib/math-bigint.md) | Basic mathematical operations with integers
| [math.fixed](zlib/math-fixed.md)   | Basic mathematical operations with fixed-point numbers
| math.float                         |
| str.bool                           |
