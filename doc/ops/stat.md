# stat

Statistical operations.

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [sum](#sum)             |          | Sum

## sum

The sum of all items on the stack.

    ( a:BigInt*   -- add:BigInt );   or
    ( a:Decimal*  -- add:Decimal );  or
    ( a:Float*    -- add:Float );    or
    ( a:Rational* -- add:Rational ); or
    ( a:Complex*  -- add:Complex )

Example:

<!-- test: sum -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `3`           | `1 \| 2 \| 3`
| `4`           | `1 \| 2 \| 3 \| 4`
| `sum`         | `10`
