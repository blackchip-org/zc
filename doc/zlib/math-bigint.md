<!-- mod: math.bigint -->

# math.bigint

Basic mathematical operations with integers.

- Use: import

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [abs](#abs)             | Absolute value
| [add](#add)             | Addition
| [ceil](#ceil)           | Ceiling
| [div](#div)             | Division
| [floor](#floor)         | Floor
| [mod](#mod)             | Modulus
| [mul](#mul)             | Multiplication
| [neg](#neg)             | Negation
| [pow](#pow)             | Exponentiation
| [rem](#rem)             | Remainder
| [sign](#sign)           | Sign
| [sub](#sub)             | Subtraction


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:BigInt -- abs:BigInt )

Example:

<!-- test: abs -->

| Input             | Stack
|-------------------|-------------
| `-6`              | `-6`
| `math.bigint.abs` | `6`


## add

Adds the value of `b` to `a`.

    ( a:BigInt b:BigInt -- add:BigInt )

Example:

<!-- test: add -->

| Input             | Stack
|-------------------|-------------
| `6`               | `6`
| `2`               | `6 \| 2`
| `math.bigint.add` | `8`


## ceil

The identity operation with integers.

    ( a:BigInt -- a:BigInt )

Example:

<!-- test: ceil -->

| Input               | Stack
|---------------------|-------------
| `6`                 | `6`
| `math.bigint.floor` | `6`


## div

Divides the value of `a` by `b`.

    ( a:BigInt b:BigInt -- div:BigInt )

Example:

<!-- test: div -->

| Input             | Stack
|-------------------|-------------
| `6`               | `6`
| `2`               | `6 \| 2`
| `math.bigint.div` | `3`


## floor

The identity operation with integers.

    ( a:BigInt -- a:BigInt )

Example:

<!-- test: floor -->

| Input               | Stack
|---------------------|-------------
| `6`                 | `6`
| `math.bigint.floor` | `6`


## mod

The modulus when `a` is divided by `b`.

    ( a:BigInt b:BigInt -- mod:BigInt )

Example:

<!-- test: mod -->

| Input             | Stack
|-------------------|-------------
| `-7`              | `-7`
| `2`               | `-7 \| 2`
| `math.bigint.mod` | `1`


## neg

Changes the sign of `a`.

    ( a:BigInt -- neg:BigInt )

Example:

<!-- test: neg -->

| Input             | Stack
|-------------------|-------------
| `-6`              | `-6`
| `math.bigint.neg` | `6`
| `math.bigint.neg` | `-6`


## mul

Multiplies `a` by `b`.

    ( a:BigInt b:BigInt -- mul:Big )

Example:

<!-- test: mul -->

| Input             | Stack
|-------------------|-------------
| `6`               | `6`
| `2`               | `6 \| 2`
| `math.bigint.mul` | `12`


## pow

Raises `a` to the power of `b`.

    ( a:BigInt b:BigInt -- pow:BigInt )

Example:

<!-- test: pow -->

| Input             | Stack
|-------------------|-------------
| `6`               | `6`
| `2`               | `6 \| 2`
| `math.bigint.pow` | `36`


## rem

The remainder when `a` is divided by `b`.

    ( a:BigInt b:BigInt -- rem:BigInt )

Example:

<!-- test: rem -->

| Input             | Stack
|-------------------|-------------
| `-7`              | `-7`
| `2`               | `-7 \| 2`
| `math.bigint.rem` | `-1`


## sign

If:

* `a` is negative: `-1`
* `a` is positive: `1`
* `a` is zero: `0`

```
( a:BigInt -- sign:BigInt )
```

Example:

<!-- test: sign -->

| Input                 | Stack
|-----------------------|-------------
| `-6 math.bigint.sign` | `-1`
| `clear`               |
| `6 math.bigint.sign`  | `1`
| `clear`               |
| `0 math.bigint.sign`  | `0`


## sub

Subtracts `b` from `a`.

    ( a:BigInt b:BigInt -- sub:BigInt )


Example:

<!-- test: sub -->

| Input             | Stack
|-------------------|-------------
| `6`               | `6`
| `2`               | `6 \| 2`
| `math.bigint.sub` | `4`