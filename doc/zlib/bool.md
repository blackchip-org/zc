# bool

Boolean operations.

- Prelude: dev
- Use: include

Comparision operations are parsed as numerics when possible. Therefore the following values are considered equal:

- `1234.56`
- `1234.560`
- `1,234.56`
- `$1,234.56`
- `+1,234.56`

| Operation         | Alias | Description
|-------------------|-------|---------------
| [and](#and)       |       | Logical conjunction
| [eq](#eq)         |       | Equals
| [false](#false)   | `f`   | False
| [gt](#gt)         |       | Greater than
| [gte](#gte)       |       | Greater than or equals
| [true](#true)     | `t`   | True
| [lt](#lt)         |       | Less than
| [lte](#lte)       |       | Less than or equals
| [neq](#neq)       |       | Not Equals
| [not](#not)       |       | Negation
| [or](#or)         |       | Logical disjunction

## and

The logical conjunction of `a` and `b` is placed on the stack.

    ( a:Bool b:Bool -- and:Bool )

Example:

| Input   | Stack
|---------|-------------|
| `t`     | `true`
| `t`     | `true \| true`
| `and`   | `true`
| `f`     | `true \| false`
| `and`   | `false`

## eq

Places `true` on the stack if `a` and `b` are equal to each other, otherwise `false`.

    ( a:Val b:Val -- Bool )

Examples:

| Input                  | Stack
|------------------------|-------------|
| `1234.56 1,234.56 eq`  | `true`
| `clear`                |
| `1234.56 1234.560 eq`  | `true`
| `clear`                |
| `1234.56 $1,234.56 eq` | `true`
| `clear`                |
| `1234.56 +1,234.56 eq` | `true`

## false

Places `false` on the stack.

    ( -- 'false' )

Example:

| Input    | Stack
|----------|-------------|
| `false`  | `false`


## gt

If `a` is greater then `b` then `true` is placed on the stack, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 gt`   | `true`
| `clear`    |
| `0 0 gt`   | `false`
| `clear`    |
| `-1 0 gt`  | `false`

