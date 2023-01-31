<!-- import: math.fixed -->

# math.fixed

Basic mathematical operations with fixed-point numbers.

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
| [sign](#sign)           | Sign
| [sub](#sub)             | Subtraction


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:Fixed -- abs:Fixed )

Example:

<!-- test: abs -->

| Input            | Stack
|------------------|-------------
| `-6.6`           | `-6.6`
| `math.fixed.abs` | `6.6`


## add

Adds the value of `b` to `a`.

    ( a:Fixed b:Fixed -- add:Fixed )

Example:

<!-- test: add -->

| Input            | Stack
|------------------|-------------
| `6.6`            | `6.6`
| `2.2`            | `6.6 \| 2.2`
| `math.fixed.add` | `8.8`


## ceil

The nearest integer value greater than or equal to `a`.

    ( a:Fixed -- ceil:Fixed )

Example:

<!-- test: ceil -->

| Input             | Stack
|-------------------|-------------
| `6.12`            | `6.12`
| `math.fixed.ceil` | `7`


## div

Divides the value of `a` by `b`.

    ( a:Fixed b:Fixed -- div:Fixed )

Example:

<!-- test: div -->

| Input            | Stack
|------------------|-------------
| `6.6`            | `6.6`
| `2.2`            | `6.6 \| 2.2`
| `math.fixed.div` | `3`


## floor

The nearest integer value less than or equal to `a`.

    ( a:Fixed -- floor:Fixed )

Example:

<!-- test: floor -->

| Input              | Stack
|--------------------|-------------
| `6.12`             | `6.12`
| `math.fixed.floor` | `6`


## mod

The modulus when `a` is divided by `b`.

    ( a:Fixed b:Fixed -- mod:Fixed )

Example:

<!-- test: mod -->

| Input            | Stack
|------------------|-------------
| `-7.7`           | `-7.7`
| `2`              | `-7.7 \| 2`
| `math.fixed.mod` | `-1.7`


## neg

Changes the sign of `a`.

    ( a:Fixed -- neg:Fixed )

Example:

<!-- test: neg -->

| Input            | Stack
|------------------|-------------
| `-6.6`           | `-6.6`
| `math.fixed.neg` | `6.6`
| `math.fixed.neg` | `-6.6`


## mul

Multiplies `a` by `b`.

    ( a:Fixed b:Fixed -- Num )

Example:

<!-- test: mul -->

| Input            | Stack
|------------------|-------------
| `6.6`            | `6.6`
| `2.2`            | `6.6 \| 2.2`
| `math.fixed.mul` | `14.52`


## pow

Raises `a` to the power of `b`.

    ( a:Fixed b:Fixed -- Num )

Example:

<!-- test: pow -->

| Input            | Stack
|------------------|-------------
| `6.6`            | `6.6`
| `2`              | `6.6 \| 2`
| `math.fixed.pow` | `43.56`


## sign

If:

* `a` is negative: `-1`
* `a` is positive: `1`
* `a` is zero: `0`

```
( a:Fixed -- sign:Fixed )
```

Example:

<!-- test: sign -->

| Input                  | Stack
|------------------------|-------------
| `-6.6 math.fixed.sign` | `-1`
| `clear`                |
| `6.6 math.fixed.sign`  | `1`
| `clear`                |
| `0.0 math.fixed.sign`  | `0`


## sub

Subtracts `b` from `a`.

    ( a:Fixed b:Fixed -- sub:Fixed )

<!-- test: sub -->

| Input            | Stack
|------------------|-------------
| `6.6`            | `6.6`
| `2.2`            | `6.6 \| 2.2`
| `math.fixed.sub` | `4.4`