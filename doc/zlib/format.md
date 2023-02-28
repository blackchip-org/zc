# format

<!-- eval: use format -->

Numeric formatting

- Prelude: user, dev (use)

<!-- index -->

| Operation                   | Description
|-----------------------------|----------------
| [precision](#precision)     | Precision for fixed-point math, set
| [precision=](#precision=)   | Prevision for fixed-point math, get

## precision

Sets the precision to `a` for fixed-point math which is the number of digits
used after the decimal point. Extra digits are rounded using the current
rounding mode. If set to zero, no rounding is performed.

    ( a:Int32 -- )

Example:

<!-- test: places -->

| Input               | Stack
|---------------------|---------------------|
| `2 3 div`           | `0.6666666666666667`
| `clear`             |
| `2 precision`       |
| `2 3 div`           | `0.67`


## precision=

Gets the precision for fixed-point math.

    ( -- places:Int32 )

Example:

<!-- test: places= -->

| Input               | Stack
|---------------------|---------------------|
| `precision=`        | `0`

