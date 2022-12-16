# fixed-bool

Boolean operations using fixed-point math.

- Use: import

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

| Input            | Stack
|------------------|------------------
| `'1234.56'`      | `1234.56`
| `1,234.56`       | `1234.56 \| 1,234.56`
| `fixed-bool.eq`  | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- gt:Bool )

Example:

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 fixed-bool.gt`  | `true`
| `clear`                  |
| `2.2 2.2 fixed-bool.gt`  | `false`
| `clear`                  |
| `1.1 2.2 fixed-bool.gt`  | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- gt:Bool )

Example:

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 fixed-bool.gte` | `true`
| `clear`                  |
| `2.2 2.2 fixed-bool.gte` | `true`
| `clear`                  |
| `1.1 2.2 fixed-bool.gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- lt:Bool )

Example:

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 fixed-bool.lt`  | `false`
| `clear`                  |
| `2.2 2.2 fixed-bool.lt`  | `false`
| `clear`                  |
| `1.1 2.2 fixed-bool.lt`  | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Fixed b:Fixed -- lte:Bool )

Example:

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 fixed-bool.lte` | `false`
| `clear`                  |
| `2.2 2.2 fixed-bool.lte` | `true`
| `clear`                  |
| `1.1 2.2 fixed-bool.lte` | `true`



## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Fixed b:Fixed -- neq:Bool )

| Input                    | Stack
|--------------------------|-------------
| `3.3 2.2 fixed-bool.neq` | `true`
| `clear`                  |
| `2.2 2.2 fixed-bool.neq` | `false`

