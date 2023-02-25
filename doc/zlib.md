# zlib

The zc standard library.

## Default Modules

The following modules are loaded by default:

| Name                           | Description
|--------------------------------|--------------------------------------
| [conf](zlib/conf.md)           | Configuration and settings
| [math](zlib/math.md)           | Basic mathematical operations
| [stack](zlib/stack.md)         | Stack manipulations

## User Modules

These modules can be imported or used to extend the features of the
calculator:

| Name                           | Description
|--------------------------------|----------------------------------------
| [prog](zlib/prog.md)           | Programmer's calculator
| [si](zlib/si.md)               | Prefixes for the International System of Units
| [str](zlib/str.md)             | String operations
| [time](zlib/time.md)           | Date, time, and duration operations
| [tz](zlib/tz.md)               | Time zone database
| [unit](zlib/unit.md)           | Units of measure

## Developer's Modules

These modules are useful when developing calculator scripts:

| Name                           | Description
|--------------------------------|----------------------------------------
| [assert](zlib/assert.md)       | Assertions
| [bool](zlib/bool.md)           | Boolean operations
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
