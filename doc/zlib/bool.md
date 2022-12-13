# bool

Boolean operations.

- Prelude: dev
- Use: include

Comparison operations are parsed as numerics when possible. Therefore the following values are considered equal:

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
| [iif](#iif)       |       | Immediate if
| [true](#true)     | `t`   | True
| [lt](#lt)         |       | Less than
| [lte](#lte)       |       | Less than or equals
| [neq](#neq)       |       | Not Equals
| [not](#not)       |       | Negation
| [or](#or)         |       | Logical disjunction

## and

The logical conjunction of `a` and `b` is placed on the stack.

    ( a:Bool b:Bool -- Bool )

Example:

| Input       | Stack
|-------------|-------------|
| `t t and`   | `true`
| `clear`     |
| `t f and`   | `false`
| `clear`     |
| `f f and`   | `false`

## eq

Places `true` on the stack if `a` and `b` are equal to each other, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

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

If `a` is greater than `b` then `true` is placed on the stack, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 gt`   | `true`
| `clear`    |
| `0 0 gt`   | `false`
| `clear`    |
| `-1 0 gt`  | `false`

## gte

If `a` is greater than or equal to `b` then `true` is placed on the stack, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 gte`  | `true`
| `clear`    |
| `0 0 gte`  | `true`
| `clear`    |
| `-1 0 gt`  | `false`

## iif

If `c` is true then `a` is placed on the stack, otherwise `b`.

    ( c=true  a:Val b:Val -- a:Val ); or
    ( c=false a:Val b:Val -- b:Val )

Example:

| Input        | Stack
|--------------|-------------|
| `true`       | `true`
| `'yes' 'no'` | `true \| yes \| no`
| `iif`        | `yes`

## true

Places `true` on the stack.

    ( -- 'true' )

Example:

| Input    | Stack
|----------|-------------|
| `true`   | `true`

## lt

If `a` is less than `b`, then `true` is placed on the stack, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 lt`   | `false`
| `clear`    |
| `0 0 lt`   | `false`
| `clear`    |
| `-1 0 lt`  | `true`

## lte

If `a` is less than or equal to `b`, then `true` is placed on the stack, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 lte`  | `false`
| `clear`    |
| `0 0 lte`  | `true`
| `clear`    |
| `-1 0 lte` | `true`

## neq

Places `true` on the stack if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Val b:Val -- Bool )

Example:

| Input                  | Stack
|------------------------|-------------|
| `123 123 neq`          | `false`
| `clear`                |
| `123 456 neq`          | `true`

## not

Places `true` on the stack if `a` is false, otherwise `true`

Example:

| Input                  | Stack
|------------------------|-------------|
| `true not`             | `false`
| `not`                  | `true`

## or

The logical disjunction of `a` and `b` is placed on the stack.

    ( a:Bool b:Bool -- Bool )

Example:

| Input       | Stack
|-------------|-------------|
| `t t or`    | `true`
| `clear`     |
| `t f or`    | `true`
| `clear`     |
| `f f or`    | `false`

