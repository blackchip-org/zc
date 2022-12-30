<!-- mod: bool.fixed -->

# bool.fixed

Boolean operations using fixed-point math.

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

    ( a:Fixed b:Fixed -- eq:Bool )

Example:

<!-- test: eq -->

| Input            | Stack
|------------------|------------------
| `'1234.56'`      | `1234.56`
| `1,234.56`       | `1234.56 \| 1,234.56`
| `bool.fixed.eq`  | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- gt:Bool )

Example:

<!-- test: gt -->

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 bool.fixed.gt`  | `true`
| `clear`                  |
| `2.2 2.2 bool.fixed.gt`  | `false`
| `clear`                  |
| `1.1 2.2 bool.fixed.gt`  | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- gt:Bool )

Example:

<!-- test: gte -->

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 bool.fixed.gte` | `true`
| `clear`                  |
| `2.2 2.2 bool.fixed.gte` | `true`
| `clear`                  |
| `1.1 2.2 bool.fixed.gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- lt:Bool )

Example:

<!-- test: lt -->

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 bool.fixed.lt`  | `false`
| `clear`                  |
| `2.2 2.2 bool.fixed.lt`  | `false`
| `clear`                  |
| `1.1 2.2 bool.fixed.lt`  | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- lte:Bool )

Example:

<!-- test: lte -->

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 bool.fixed.lte` | `false`
| `clear`                  |
| `2.2 2.2 bool.fixed.lte` | `true`
| `clear`                  |
| `1.1 2.2 bool.fixed.lte` | `true`

## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Fixed b:Fixed -- neq:Bool )

<!-- test: neq -->

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 bool.fixed.neq` | `true`
| `clear`                  |
| `2.2 2.2 bool.fixed.neq` | `false`

