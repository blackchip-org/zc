# modes

Calculator modes.

To start a calculator in a given mode, use the `-m` flag at the command line.
For example:

```bash
zc -m fin
```

| Name            | Description
|-----------------|------------------
| [dev](#dev)     | Developer's mode
| [fin](#fin)     | Financial calculator
| [plain](#plain) | Plain numbers
| [prog](#prog)   | Programmer's calculator
| [time](#time)   | Time calculations

## dev

Developer's mode.

Imports and uses a group of modules that are useful when developing
calculator scripts.

Definition: [dev.zc](../internal/modes/dev.zc)

```
use dev
use stack
use bool
use conf
use math
use str
use dict
```

## fin

Financial calculator.

Formats numbers to two decimal places and sets the rounding mode to half-even.

Definition: [fin.zc](../internal/modes/fin.zc)

```
2 conf.precision
2 conf.min-digits
'half-even' conf.rounding-mode
```

## plain

Plain numbers.

Disables default formatting.

Definition: [plain.zc](../internal/modes/plain.zc)

```
'0' conf.int-format
'false' conf.auto-currency
```

## prog

Programmer's calculator

Imports modules for bitwise and byte level math.

Definition: [prog.zc](../internal/modes/prog.zc)

```
'0' conf.int-format
'false' conf.auto-currency

use prog
```

## time

Time calculation.

Imports modules for date, time, and duration operations and a time zone
database.

Definition: [time.zc](../internal/modes/time.zc)

```
use time
use tz
```
