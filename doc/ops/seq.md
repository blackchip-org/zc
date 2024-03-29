<!-- Document generated by "gen-doc"; DO NOT EDIT -->
# seq

Integer sequences

[Examples](../examples/seq.md)

| Operation               | Description
|-------------------------|---------------
| [`fibonacci, fib`](#fibonacci) | Fibonacci sequence
| [`sequence, seq`](#sequence) | Sequence of integers


## fibonacci

Calculates the *n*th element in the Fibonacci sequence. The value of *n*
must be equal to or greater than zero.

Alias: `fib`

```
( n:Int -- BigInt* )
```

Example:

<!-- test: fibonacci -->

| Input              | Stack
|--------------------|---------------
| `1 5 seq /fib map` | `1 \| 1 \| 2 \| 3 \| 5`

## sequence

Adds the integers from *p0* to *p1* to the stack. If *p0* is greater than
*p1*, the list of integers is in decreasing order

Alias: `seq`

```
( p0:BigInt p1:BigInt -- BigInt* )
```

Example:

<!-- test: sequence -->

| Input       | Stack
|-------------|---------------
| `4 8 seq  ` | `4 \| 5 \| 6 \| 7 \| 8`
| `c 8 4 seq` | `8 \| 7 \| 6 \| 5 \| 4`
