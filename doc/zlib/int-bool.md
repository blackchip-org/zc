# int-bool

Boolean operations using integer math.

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

    ( a:Int b:Int -- eq:Bool )

Example:

| Input        | Stack
|--------------|------------------|
| `'1234`      | `1234`
| `1,234`      | `1234 \| 1,234`
| `eq`         | `true`


## gt

`true` if `a` is greater than `b`, otherwise `false`.

    ( a:Int b:Int -- gt:Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 gt`   | `true`
| `clear`    |
| `0 0 gt`   | `false`
| `clear`    |
| `-1 0 gt`  | `false`


## gte

`true` if `a` is greater than or equal to `b`, otherwise `false`.

    ( a:Int b:Int -- gt:Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 gte`  | `true`
| `clear`    |
| `0 0 gte`  | `true`
| `clear`    |
| `-1 0 gte` | `false`


## lt

`true` if `a` is less than `b`, otherwise `false`.

    ( a:Int b:Int -- lt:Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 lt`   | `false`
| `clear`    |
| `0 0 lt`   | `false`
| `clear`    |
| `-1 0 lt`  | `true`


## lte

`true` if `a` is less than or equal to `b`, otherwise `false`.

    ( a:Int b:Int -- lte:Bool )

Example:

| Input      | Stack
|------------|-------------|
| `1 0 lte`  | `false`
| `clear`    |
| `0 0 lte`  | `true`
| `clear`    |
| `-1 0 lte` | `true`


## neq

`true` if `a` and `b` are not equal to each other, otherwise `false`.

    ( a:Int b:Int -- neq:Bool )

| Input         | Stack
|---------------|-------------|
| `1 2 neq`     | `true`
| `clear`       |
| `2 2 neq`     | `false`


