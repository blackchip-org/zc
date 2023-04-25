# hof

Higher order functions

<!-- index -->

| Operation           | Alias    | Description
|---------------------|----------|----------------
| [eval](#eval)       |          | Evaluate top of stack
| [filter](#filter)   |          | Filter items in the stack
| [fold](#fold)       | `reduce` | Reduce items to a single value
| [map](#map)         |          | Apply a function to each item on the stack
| [repeat](#repeat)   |         | Repeat the execution of a function


## eval

Evaluate *expr* as if it was input to the calculator.

    ( Val* expr:Str -- Val* )


<!-- test: eval -->

| Input               | Stack
|---------------------|---------------------|
| `'1 2 add`          | `1 2 add`
| `eval`              | `3`


## filter

Filter the stack by keeping items that are true when evaluated by expression *expr*.

    ( Val* expr:Str -- Val* )

Example which filters the stack to only keep even numbers:

<!-- test: filter -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5 6`       | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `'2 mod 0 eq`       | `1 \| 2 \| 3 \| 4 \| 5 \| 6 \| 2 mod 0 eq`
| `filter`            | `2 \| 4 \| 6`

## fold

Reduce the stack to a single value using the expression *expr*. An
'invalid function' error is raised if *expr* does not reduce.

    ( Val* expr:Str -- Val )

Alias: `reduce`

Example which sums the numbers in the stack:

<!-- test: fold -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'add`              | `1 \| 2 \| 3 \| 4 \| 5 \| add`
| `fold`              | `15`


## map

Apply expression *expr* to each value in the stack.

    ( Val* expr:Str -- Val* )

Example which doubles all numbers on the stack:

<!-- test: map -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'2 mul`            | `1 \| 2 \| 3 \| 4 \| 5 \| 2 mul`
| `map`               | `2 \| 4 \| 6 \| 8 \| 10`


## repeat

Repeat execution of expression *expr* for *n* times.

    ( Val* expr:Str n:Int -- Val* )

Example:

<!-- test: repeat -->

| Input               | Stack
|---------------------|---------------------|
| `1`                 | `1`
| `'2 mul`            | `1 \| 2 mul`
| `8 repeat`          | `256`
