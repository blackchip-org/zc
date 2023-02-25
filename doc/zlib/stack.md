# stack

<!-- eval: use stack -->

Stack manipulations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [at](@at)               |          | Stack element at
| [clear](#clear)         | `c`      | Clear
| [drop](#drop)           |          | Drop top item
| [dup](#dup)             |          | Duplicate top item
| [n](#n)                 |          | Number of items on the stack
| [top](#top)             |          | Keep the top of the stack and discard the rest


## at

The nth stack element starting at the top of the stack.

    ( ...:any n:any ...:any -- ...:any n:any ...:any n:any )

Example:

<!-- test: at -->

| Input         | Stack
|---------------|-------------|
| `1 2 3 4 5`   | `1 \| 2 \| 3 \| 4 \| 5`
| `0 at`        | `1 \| 2 \| 3 \| 4 \| 5 \| 5`
| `drop 1 at`   | `1 \| 2 \| 3 \| 4 \| 5 \| 4`


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

| Input             | Stack
|-------------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `n`               | `a \| b \| c \| d \| 4`


# top

Keep the top of the stack and discard the rest

    ( ...:Val top:Val -- top:Val )

Example:

<!-- test: top -->

| Input             | Stack
|-------------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `top`             | `d`


