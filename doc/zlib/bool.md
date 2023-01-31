<!-- use: bool -->

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

<!-- index -->

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

The logical conjunction of `a` and `b`.

    ( a:Bool b:Bool -- and:Bool )

Example:

<!-- test: and -->

| Input       | Stack
|-------------|-------------|
| `t t and`   | `true`
| `clear`     |
| `t f and`   | `false`
| `clear`     |
| `f f and`   | `false`

## eq

`true` if `a` and `b` are equal, otherwise `false`.

    ( a:Val b:Val -- eq:Bool )

Example:

<!-- test: eq -->

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

<!-- test: false -->

| Input    | Stack
|----------|-------------|
| `false`  | `false`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Val b:Val -- gt:Bool )

Example:

<!-- test: gt -->

| Input      | Stack
|------------|-------------|
| `1 0 gt`   | `true`
| `clear`    |
| `0 0 gt`   | `false`
| `clear`    |
| `-1 0 gt`  | `false`

## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Val b:Val -- gte:Bool )

Example:

<!-- test: false -->

| Input      | Stack
|------------|-------------|
| `1 0 gte`  | `true`
| `clear`    |
| `0 0 gte`  | `true`
| `clear`    |
| `-1 0 gt`  | `false`

## iif

`t` if `cond` is `true`, otherwise `f`.

    ( cond=true  t:Val f:Val -- iif:Val ); or
    ( cond=false t:Val f:Val -- iif:Val )

Example:

<!-- test: iif -->

| Input        | Stack
|--------------|-------------|
| `true`       | `true`
| `'yes' 'no'` | `true \| yes \| no`
| `iif`        | `yes`

## true

Places `true` on the stack.

    ( -- 'true' )

Example:

<!-- test: true -->

| Input    | Stack
|----------|-------------|
| `true`   | `true`

## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Val b:Val -- lt:Bool )

Example:

<!-- test: lt -->

| Input      | Stack
|------------|-------------|
| `1 0 lt`   | `false`
| `clear`    |
| `0 0 lt`   | `false`
| `clear`    |
| `-1 0 lt`  | `true`

## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Val b:Val -- lte:Bool )

Example:

<!-- test: lte -->

| Input      | Stack
|------------|-------------|
| `1 0 lte`  | `false`
| `clear`    |
| `0 0 lte`  | `true`
| `clear`    |
| `-1 0 lte` | `true`

## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Val b:Val -- neq:Bool )

Example:

<!-- test: neq -->

| Input                  | Stack
|------------------------|-------------|
| `123 123 neq`          | `false`
| `clear`                |
| `123 456 neq`          | `true`


## not

`true` if `a` is false, otherwise `false`

    ( a:Bool -- not:Bool )

Example:

<!-- test: not -->

| Input                  | Stack
|------------------------|-------------|
| `true not`             | `false`
| `not`                  | `true`


## or

The logical disjunction of `a` and `b`.

    ( a:Bool b:Bool -- or:Bool )

Example:

<!-- test: or -->

| Input       | Stack
|-------------|-------------|
| `t t or`    | `true`
| `clear`     |
| `t f or`    | `true`
| `clear`     |
| `f f or`    | `false`

