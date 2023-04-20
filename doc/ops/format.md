# format

Numeric formatting

<!-- index -->

| Operation                         | Alias | Description
|-----------------------------------|----------------
| [round](#round)                   | `r`   | Round to a given precision
| [rounding-mode](#rounding-mode)   |       | Method to use in rounding, set
| [rounding-mode=](#rounding-mode=) |       | Method to use in rounding, get

## round

Rounds the number `n` to `d` digits using the current rounding mode.

    ( n:Num d:Int -- round:Num )

Example:

<!-- test: round -->

| Input               | Stack
|---------------------|---------------------|
| `2 3 div`           | `0.6666666666666667`
| `2 round`           | `0.67`


## rounding-mode

Sets the mode to be used when rounding. Valid modes are:

- `half-up`
- `ceil`
- `down`
- `floor`
- `half-even`
- `up`

    ( m:Str -- )

Example:

<!-- test: rounding-mode -->

| Input                | Stack
|----------------------|---------------------
| `1.01 0.05 mul`      | `0.0505`
| `2 round`            | `0.05`
| `'up' rounding-mode` | *rounding-mode set to 'up'*
| `clear`              |
| `1.01 0.05 mul`      | `0.0505`
| `2 round`            | `0.06`


## rounding-mode=

Gets the current rounding mode

    ( -- m:Str )

Example:

<!-- test rounding-mode= -->

| Input                | Stack
|----------------------|---------------------
| `rounding-mode=`     | `half-up`
