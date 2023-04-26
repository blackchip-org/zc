# stat

Statistical operations.

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [sum](#sum)             |          | Sum

## sum

The sum of all items on the stack.

    ( T* -- T )

Where *T* is one of
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Example:

<!-- test: sum -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `3`           | `1 \| 2 \| 3`
| `4`           | `1 \| 2 \| 3 \| 4`
| `sum`         | `10`
