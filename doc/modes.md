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
| [geo](#geo)     | Geographer's calculator
| [prog](#prog)   | Programmer's calculator
| [sci](#sci)     | Scientific calculator
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
Enables auto formatting and auto currency.

Definition: [fin.zc](../internal/modes/fin.zc)

```
2 precision
2 min-digits
'half-even' rounding-mode
',000' int-layout
true auto-currency
true auto-format
```

## geo

Geographer's calculator.

Imports and uses modules for geo-spatial calculations.

Definition: [geo.zc](../internal/modes/geo.zc)

```
use geo
import epsg
```


## prog

Programmer's calculator

Imports modules for bitwise and byte level math.

Definition: [prog.zc](../internal/modes/prog.zc)

```
'0' int-layout
'false' auto-format
'false' auto-currency

use [
    prog
    unicode
]
```

## sci

Scientific calculator.

Imports advanced math functions and units of measure conversions.

Definition: [sci.zc](../internal/modes/sci.zc)

```
'0' int-layout
'false' auto-currency
'false' auto-format

use sci
use si
use unit
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
