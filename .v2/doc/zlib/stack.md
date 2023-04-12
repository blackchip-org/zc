# stack

<!-- eval: use stack -->

Stack manipulations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [at](@at)               |          | Take stack element at index
| [clear](#clear)         | `c`      | Clear
| [drop](#drop)           |          | Drop top item
| [down](#down)           | `dn`     | Rotate stack by moving items downward
| [dup](#dup)             |          | Duplicate top item
| [empty](#empty)         |          | True if the stack is empty
| [n](#n)                 |          | Number of items on the stack
| [reverse](#reverse)     | `rev`    | Reverse items on the stack
| [swap](#swap)           | `sw`     | Swap the top two items on the stack
| [take](#take)           |          | Take top items from the stack
| [top](#top)             |          | Keep the top of the stack and discard the rest


## at

Take the `n`th stack element, where 0 is the bottom of the stack,
and discard the rest.

    ( ...:Val n:Int -- at:Val )

Example:

<!-- test: at -->

| Input         | Stack
|---------------|-------------|
| `1 2 3 4 5`   | `1 \| 2 \| 3 \| 4 \| 5`
| `1 at`        | `2`


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


## down

Rotate items on the stack by moving downward.

In the interactive calculator, the top of the stack is towards the bottom of the terminal so downward means seeing all items moves toward the bottom.
The top of the stack wraps around to be the bottom of the stack.

    ( ...:Val -- ...:Val )

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


## empty

True if the stack is empty

    ( ...:Any -- ...:Any empty:Bool )


Example:

<!-- test: empty -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `empty`       | `1 \| false`
| `clear`       |
| `empty`       | `true`


## n

Number of items on the stack.

    ( ...:Val -- ...:Val a:Int )

Example:

<!-- test: n -->

| Input             | Stack
|-------------------|-------------|
| `'a' 'b' 'c' 'd'` | `a \| b \| c \| d`
| `n`               | `a \| b \| c \| d \| 4`


## reverse

Reverses the elements on the stack.

    ( ...:Val -- ...:Val )

Alias: `rev`

Example:

<!-- test: reverse -->

| Input             | Stack
|-------------------|-------------|
| `1 2 3 4 5`       | `1 \| 2 \| 3 \| 4 \| 5`
| `reverse`         | `5 \| 4 \| 3 \| 2 \| 1`


## swap

Swap the first two items on the stack

    ( a:Val b:Val -- b:Val a:Val )

Alias: `sw`

<!-- test: swap -->

| Input             | Stack
|-------------------|-------------|
| `1 2`             | `1 \| 2`
| `swap`            | `2 \| 1`
| `swap`            | `1 \| 2`


## take

Take the top `n` elements from the stack and discard the rest.

    ( ...:Val n:Int -- take...:Val )

Example:

<!-- test: take -->

| Input             | Stack
|-------------------|-------------|
| `1 2 3 4 5`       | `1 \| 2 \| 3 \| 4 \| 5`
| `2 take`          | `4 \| 5`


## top

Keep the top of the stack and discard the rest

    ( ...:Val top:Val -- top:Val )

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

    ( ...:Val -- ...:Val )

Example:

<!-- test: up -->

| Input         | Stack
|---------------|-------------|
| `1 2 3`       | `1 \| 2 \| 3`
| `up`          | `2 \| 3 \| 1`
| `up`          | `3 \| 1 \| 2`

