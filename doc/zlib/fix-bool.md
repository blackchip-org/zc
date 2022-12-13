# fix-bool

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

    ( a:Fix b:Fix -- eq:Bool )

Example:

| Input        | Stack
|--------------|------------------|
| `'1234.56'`  | `1234.56`
| `1,234.56`   | `1234.56 \| 1,234.56`
| `eq`         | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Fix b:Fix -- gt:Bool )

Example:

| Input         | Stack
|---------------|-------------|
| `3.3 2.2 gt`  | `true`
| `clear`       |
| `2.2 2.2 gt`  | `false`
| `clear`       |
| `1.1 2.2 gt`  | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Fix b:Fix -- gt:Bool )

Example:

| Input         | Stack
|---------------|-------------|
| `3.3 2.2 gte` | `true`
| `clear`       |
| `2.2 2.2 gte` | `true`
| `clear`       |
| `1.1 2.2 gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Fix b:Fix -- lt:Bool )

Example:

| Input         | Stack
|---------------|-------------|
| `3.3 2.2 lt`  | `false`
| `clear`       |
| `2.2 2.2 lt`  | `false`
| `clear`       |
| `1.1 2.2 lt`  | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Fix b:Fix -- lte:Bool )

Example:

| Input         | Stack
|---------------|-------------|
| `3.3 2.2 lte` | `false`
| `clear`       |
| `2.2 2.2 lte` | `true`
| `clear`       |
| `1.1 2.2 lte` | `true`


## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Fix b:Fix -- neq:Bool )

| Input         | Stack
|---------------|-------------|
| `3.3 2.2 neq` | `true`
| `clear`       |
| `2.2 2.2 neq` | `false`

