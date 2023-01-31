<!-- import: assert -->

# assert

Assertions.

For each assertion function, if the checked condition is not true, the current execution is aborted.

- Use: import

<!-- index -->

| Operation             | Description
|-----------------------|-----------------------
| [eq](#eq)             | Assert equals
| [f](#f)               | Assert false
| [t](#t)               | Assert true


## eq

Checks that `a` is equal to `b` and aborts execution if not.

    ( a:Val b:Val -- )

Example:

<!-- test: eq -->

| Input           | Stack
|-----------------|-------------
| `1 2`           | `1 \| 2`
| `try assert.eq` | `assertion failed: 1 == 2 \| false`


## f

Checks that `a` is false and aborts execution if not.

    ( a:Bool -- )

Example:

<!-- test: f -->

| Input           | Stack
|-----------------|-------------
| `'true'`        | `true`
| `try assert.f`  | `assertion failed: not false \| false`


## t

Checks that `a` is true and aborts execution if not.

    ( a:Bool -- )

Example:

<!-- test: t -->

| Input           | Stack
|-----------------|-------------
| `'false'`       | `false`
| `try assert.t`  | `assertion failed: not true \| false`



