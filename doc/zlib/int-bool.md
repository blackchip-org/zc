<!-- mod: int-bool -->

# int-bool

Boolean operations using integer math.

- Use: import

<!-- index -->

| Operation         | Description
|-------------------|----------------
| [eq](#eq)         | Equals
| [gt](#gt)         | Greater than
| [gte](#gte)       | Greater than or equals
| [lt](#lt)         | Less than
| [lte](#lte)       | Less than or equals
| [neq](#neq)       | Not Equals


## eq

`true` if `a` and `b` are equal, otherwise `false`.

    ( a:BigInt b:BigInt -- eq:Bool )

Example:

<!-- test: eq -->

| Input         | Stack
|---------------|------------------
| `'1234`       | `1234`
| `1,234`       | `1234 \| 1,234`
| `int-bool.eq` | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:BigInt b:BigInt -- gt:Bool )

Example:

<!-- test: gt -->

| Input               | Stack
|---------------------|-------------
| `1 0 int-bool.gt`   | `true`
| `clear`             |
| `0 0 int-bool.gt`   | `false`
| `clear`             |
| `-1 0 int-bool.gt`  | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:BigInt b:BigInt -- gt:Bool )

Example:

<!-- test: gte -->

| Input               | Stack
|---------------------|-------------
| `1 0 int-bool.gte`  | `true`
| `clear`             |
| `0 0 int-bool.gte`  | `true`
| `clear`             |
| `-1 0 int-bool.gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:BigInt b:BigInt -- lt:Bool )

Example:

<!-- test: lt-->

| Input               | Stack
|---------------------|-------------
| `1 0 int-bool.lt`   | `false`
| `clear`             |
| `0 0 int-bool.lt`   | `false`
| `clear`             |
| `-1 0 int-bool.lt`  | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:BigInt b:BigInt -- lte:Bool )

Example:

<!-- test: lte -->

| Input               | Stack
|---------------------|-------------
| `1 0 int-bool.lte`  | `false`
| `clear`             |
| `0 0 int-bool.lte`  | `true`
| `clear`             |
| `-1 0 int-bool.lte` | `true`


## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:BigInt b:BigInt -- neq:Bool )

Example:

<!-- test: neq -->

| Input              | Stack
|--------------------|-------------
| `1 2 int-bool.neq` | `true`
| `clear`            |
| `2 2 int-bool.neq` | `false`


