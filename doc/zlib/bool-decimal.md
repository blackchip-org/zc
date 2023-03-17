# bool.decimal

<!-- eval: import bool.decimal -->

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

    ( a:Decimal b:Decimal -- eq:Bool )

Example:

<!-- test: eq -->

| Input             | Stack
|-------------------|------------------
| `1234.56`         | `1234.56`
| `1234.56`         | `1234.56 \| 1234.56`
| `bool.decimal.eq` | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Decimal b:Decimal -- gt:Bool )

Example:

<!-- test: gt -->

| Input                     | Stack
|---------------------------|-------------
| `3.3 2.2 bool.decimal.gt` | `true`
| `clear`                   |
| `2.2 2.2 bool.decimal.gt` | `false`
| `clear`                   |
| `1.1 2.2 bool.decimal.gt` | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Decimal b:Decimal -- gt:Bool )

Example:

<!-- test: gte -->

| Input                      | Stack
|----------------------------|-------------
| `3.3 2.2 bool.decimal.gte` | `true`
| `clear`                    |
| `2.2 2.2 bool.decimal.gte` | `true`
| `clear`                    |
| `1.1 2.2 bool.decimal.gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Decimal b:Decimal -- lt:Bool )

Example:

<!-- test: lt -->

| Input                     | Stack
|---------------------------|-------------
| `3.3 2.2 bool.decimal.lt` | `false`
| `clear`                   |
| `2.2 2.2 bool.decimal.lt` | `false`
| `clear`                   |
| `1.1 2.2 bool.decimal.lt` | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Decimal b:Decimal -- lte:Bool )

Example:

<!-- test: lte -->

| Input                      | Stack
|----------------------------|-------------
| `3.3 2.2 bool.decimal.lte` | `false`
| `clear`                    |
| `2.2 2.2 bool.decimal.lte` | `true`
| `clear`                    |
| `1.1 2.2 bool.decimal.lte` | `true`

## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Decimal b:Decimal -- neq:Bool )

<!-- test: neq -->

| Input                      | Stack
|----------------------------|-------------
| `3.3 2.2 bool.decimal.neq` | `true`
| `clear`                    |
| `2.2 2.2 bool.decimal.neq` | `false`

