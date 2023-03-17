# math.decimal

<!-- eval: import math.decimal -->

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

    ( a:Decimal -- abs:Decimal )

Example:

<!-- test: abs -->

| Input              | Stack
|--------------------|-------------
| `-6.6`             | `-6.6`
| `math.decimal.abs` | `6.6`


## add

Adds the value of `b` to `a`.

    ( a:Decimal b:Decimal -- add:Decimal )

Example:

<!-- test: add -->

| Input              | Stack
|--------------------|-------------
| `6.6`              | `6.6`
| `2.2`              | `6.6 \| 2.2`
| `math.decimal.add` | `8.8`


## ceil

The nearest integer value greater than or equal to `a`.

    ( a:Decimal -- ceil:Decimal )

Example:

<!-- test: ceil -->

| Input               | Stack
|---------------------|-------------
| `6.12`              | `6.12`
| `math.decimal.ceil` | `7`


## div

Divides the value of `a` by `b`.

    ( a:Decimal b:Decimal -- div:Decimal )

Example:

<!-- test: div -->

| Input              | Stack
|--------------------|-------------
| `6.6`              | `6.6`
| `2.2`              | `6.6 \| 2.2`
| `math.decimal.div` | `3`


## floor

The nearest integer value less than or equal to `a`.

    ( a:Decimal -- floor:Decimal )

Example:

<!-- test: floor -->

| Input                | Stack
|----------------------|-------------
| `6.12`               | `6.12`
| `math.decimal.floor` | `6`


## mod

The modulus when `a` is divided by `b`.

    ( a:Decimal b:Decimal -- mod:Decimal )

Example:

<!-- test: mod -->

| Input              | Stack
|--------------------|-------------
| `-7.7`             | `-7.7`
| `2`                | `-7.7 \| 2`
| `math.decimal.mod` | `-1.7`


## neg

Changes the sign of `a`.

    ( a:Decimal -- neg:Decimal )

Example:

<!-- test: neg -->

| Input              | Stack
|--------------------|-------------
| `-6.6`             | `-6.6`
| `math.decimal.neg` | `6.6`
| `math.decimal.neg` | `-6.6`


## mul

Multiplies `a` by `b`.

    ( a:Decimal b:Decimal -- Num )

Example:

<!-- test: mul -->

| Input              | Stack
|--------------------|-------------
| `6.6`              | `6.6`
| `2.2`              | `6.6 \| 2.2`
| `math.decimal.mul` | `14.52`


## pow

Raises `a` to the power of `b`.

    ( a:Decimal b:Decimal -- Num )

Example:

<!-- test: pow -->

| Input              | Stack
|--------------------|-------------
| `6.6`              | `6.6`
| `2`                | `6.6 \| 2`
| `math.decimal.pow` | `43.56`


## sign

If:

* `a` is negative: `-1`
* `a` is positive: `1`
* `a` is zero: `0`

```
( a:Decimal -- sign:Decimal )
```

Example:

<!-- test: sign -->

| Input                    | Stack
|--------------------------|-------------
| `-6.6 math.decimal.sign` | `-1`
| `clear`                  |
| `6.6 math.decimal.sign`  | `1`
| `clear`                  |
| `0.0 math.decimal.sign`  | `0`


## sub

Subtracts `b` from `a`.

    ( a:Decimal b:Decimal -- sub:Decimal )

<!-- test: sub -->

| Input              | Stack
|--------------------|-------------
| `6.6`              | `6.6`
| `2.2`              | `6.6 \| 2.2`
| `math.decimal.sub` | `4.4`