<!-- import: conf -->

# conf

Configurations and settings

- Prelude: user, dev
- Use: import

<!-- index -->

| Operation                   | Description
|-----------------------------|----------------
| [precision](#precision)     | Sets the precision for fixed-point math
| [precision=](#precision=)   | Gets the prevision for fixed-point math

## precision

Sets the precision to `a` for fixed-point math which is the number of digits
used after the decimal point. Extra digits are rounded using the current
rounding mode.

    ( a:Int32 -- )

Example:

<!-- test: places -->

| Input               | Stack
|---------------------|---------------------|
| `2 3 div`           | `0.6666666666666667`
| `clear`             |
| `2 conf.precision`  |
| `2 3 div`           | `0.67`


## precision=

Gets the precision for fixed-point math.

    ( -- places:Int32 )

Example:

<!-- test: places= -->

| Input               | Stack
|---------------------|---------------------|
| `conf.precision=`   | `16`

