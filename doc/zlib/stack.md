<!-- use: stack -->

# stack

Stack manipulations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [clear](#clear)         | `c`      | Clear
| [drop](#drop)           |          | Drop top item
| [dup](#dup)             |          | Duplicate top item
| [n](#n)                 |          | Number of items on the stack


## clear

Remove all items on the stack.

    ( ...:Val -- )

Example:

<!-- test: clear -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `clear`       |


## drop

Remove the top item from the stack.

    ( a:Val -- )

Example:

<!-- test: drop -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `drop`        | `1`


## dup

Duplicate the top item on the stack.

    ( a:Val -- a:Val a:Val )

Example:

<!-- test: dup -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `dup`         | `1 \| 1`


## n

Number of items on the stack.

    ( ...:Val -- ...:Val a:Int )

Example:

<!-- test: n -->

| Input         | Stack
|---------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `n`               | `a \| b \| c \| d \| 4`



