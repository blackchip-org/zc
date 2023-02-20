<!-- import: assert -->

# assert

Assertions

For each assertion function, if the checked condition is not true, the current execution is aborted.

<!-- index -->

| Operation             | Description
|-----------------------|-----------------------
| [eq](#eq)             | Assert equals
| [false](#false)       | Assert false
| [true](#true)         | Assert true


## eq

Checks that `a` is equal to `b` and aborts execution if not.

    ( a:Val b:Val -- )

Example:

<!-- test: eq -->

| Input           | Stack
|-----------------|-------------
| `1 2`           | `1 \| 2`
| `try assert.eq` | `assertion failed: 1 == 2 \| false`


## false

Checks that `a` is false and aborts execution if not.

    ( a:Bool -- )

Example:

<!-- test: false -->

| Input               | Stack
|---------------------|-------------
| `'true'`            | `true`
| `try assert.false`  | `assertion failed: not false \| false`


## true

Checks that `a` is true and aborts execution if not.

    ( a:Bool -- )

Example:

<!-- test: true -->

| Input              | Stack
|--------------------|-------------
| `'false'`          | `false`
| `try assert.true`  | `assertion failed: not true \| false`



