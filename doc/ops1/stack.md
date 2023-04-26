# stack

Stack manipulations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [clear](#clear)         | `c`      | Clear
| [drop](#drop)           |          | Drop top item
| [down](#down)           | `dn`     | Rotate stack by moving items downward
| [dup](#dup)             |          | Duplicate top item
| [n](#n)                 |          | Number of items on the stack
| [reverse](#reverse)     | `rev`    | Reverse items on the stack
| [swap](#swap)           | `sw`     | Swap the top two items on the stack
| [take](#take)           |          | Take top items from the stack
| [top](#top)             |          | Keep the top of the stack and discard the rest
| [up](#up)               |          | Rotate items on the stack by moving upward


## clear

Remove all items on the stack.

    ( ... -- )

Example:

<!-- test: clear -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `clear`       |

## down

Rotate items on the stack by moving downward.

In the interactive calculator, the top of the stack is towards the bottom of the terminal so downward means seeing all items moves toward the bottom.
The top of the stack wraps around to be the bottom of the stack.

    ( ... p0:Val -- p0:Val ... )

Alias: `dn`

Example:

<!-- test: down -->

| Input         | Stack
|---------------|-------------|
| `1 2 3`       | `1 \| 2 \| 3`
| `down`        | `3 \| 1 \| 2`
| `down`        | `2 \| 3 \| 1`

## drop

Remove the top item from the stack.

    ( Val -- )

Example:

<!-- test: drop -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `drop`        | `1`

## dup

Duplicate the top item on the stack.

    ( p0:Val -- p0:Val p0:Val )

Example:

<!-- test: dup -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `dup`         | `1 \| 1`

## n

Number of items on the stack.

    ( -- Int )

Example:

<!-- test: n -->

| Input             | Stack
|-------------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `n`               | `a \| b \| c \| d \| 4`

## reverse

Reverses the elements on the stack.

    ( p0:Val ... pn:Val -- pn:Val ... p0:Val )

Alias: `rev`

Example:

<!-- test: reverse -->

| Input             | Stack
|-------------------|-------------|
| `1 2 3 4 5`       | `1 \| 2 \| 3 \| 4 \| 5`
| `reverse`         | `5 \| 4 \| 3 \| 2 \| 1`

## swap

Swap the first two items on the stack

    ( p0:Val p1:Val -- p1:Val p0:Val )

Alias: `sw`

<!-- test: swap -->

| Input             | Stack
|-------------------|-------------|
| `1 2`             | `1 \| 2`
| `swap`            | `2 \| 1`
| `swap`            | `1 \| 2`

## take

Take the top *n* elements from the stack and discard the rest.

    ( ... Val* n:Int -- Val* )

Example:

<!-- test: take -->

| Input             | Stack
|-------------------|-------------|
| `1 2 3 4 5`       | `1 \| 2 \| 3 \| 4 \| 5`
| `2 take`          | `4 \| 5`

## top

Keep the top of the stack and discard the rest

    ( ... p0:Val -- p0:Val )

Example:

<!-- test: top -->

| Input             | Stack
|-------------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `top`             | `d`

## up

Rotate items on the stack by moving upward.

In the interactive calculator, the top of the stack is towards the bottom of
the terminal so upwards means seeing all items move toward the top. The
bottom of the stack wraps around to be the top of the stack.

    ( p0:Val ... -- ... p0:Val )

Example:

<!-- test: up -->

| Input         | Stack
|---------------|-------------|
| `1 2 3`       | `1 \| 2 \| 3`
| `up`          | `2 \| 3 \| 1`
| `up`          | `3 \| 1 \| 2`


