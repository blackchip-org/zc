# bool

Boolean operations.

- Prelude: dev
- Use: include

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

Example:

| Input    | Stack
|----------|-------------|
| `1234`   | `1,234`
| `'1234'` | `1,234 \| 1234`
| `eq`     | `true`
| `clear`  |
| `1.1`    | `1.1`
| `1.10`   | `1.1 \| 1.10`
| `eq`     | `true`

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

