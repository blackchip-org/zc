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

Evaluate the `top` of the stack as if it was input to the calculator.

    ( items:Val* top:Val -- items:Val* )


<!-- test: eval -->

| Input               | Stack
|---------------------|---------------------|
| `'1 2 add`          | `1 2 add`
| `eval`              | `3`


## filter

Filter `items` in the stack where each item true when evaluated by
function `f`.

    ( items...:Val `f`:Lambda -- filtered...:Val )

Example which filters the stack to only keep even numbers:

<!-- test: filter -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5 6`       | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `'2 mod 0 eq`       | `1 \| 2 \| 3 \| 4 \| 5 \| 6 \| 2 mod 0 eq`
| `filter`            | `2 \| 4 \| 6`


## fold

Reduce `items` to a `reduced` value using function `f`.

    ( items...:Val f:Lambda -- reduced:Val )

Alias: `reduce`

Example which sums the numbers in the stack:

<!-- test: fold -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'add`              | `1 \| 2 \| 3 \| 4 \| 5 \| add`
| `fold`              | `15`


## map

Apply function `f` to each value in `items`.

    ( items...:Val f:Lambda -- applied...:Val )

Example which doubles all numbers on thes tack:

<!-- test: map -->

| Input               | Stack
|---------------------|---------------------|
| `1 2 3 4 5`         | `1 \| 2 \| 3 \| 4 \| 5`
| `'2 mul`            | `1 \| 2 \| 3 \| 4 \| 5 \| 2 mul`
| `map`               | `2 \| 4 \| 6 \| 8 \| 10`


## repeat

Repeat execution of function `f` for `n` times.

    ( items...:Val f:Lambda n:Int -- items...:Val )

Example:

<!-- test: repeat -->

| Input               | Stack
|---------------------|---------------------|
| `1`                 | `1`
| `'2 mul`            | `1 \| 2 mul`
| `8 repeat`          | `256`
