# basic

Basic mathematical operations.

<!-- index -->

| Operation               | Alias     | Description
|-------------------------|-----------|------------
| [add](#add)             | `a`, `+`  | Addition
| [div](#div)             | `d`, `/`  | Division
| [mod](#mod)             |           | Modulus
| [mul](#mul)             | `m`, `*`  | Multiplication
| [neg](#neg)             |           | Negation
| [pow](#pow)             | `**`, `^` | Exponentiation
| [rem](#rem)             |           | Remainder
| [sign](#sign)           |           | Sign
| [sqrt](#sqrt)           |           | Square root
| [sub](#sub)             | `s`, `-`  | Subtraction


## add

Adds the value of *p1* to *p0*.

    ( p0:T p1:T -- T )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Aliases: `a`, `+`

Example:

<!-- test: add -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`

## div

Divides the value of *p0* by *p1*. If *p1* is zero, a 'division by zero' error
is raised.

    ( p0:T p1:T -- T )

Where *T* is one of:
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Aliases: `d`, `/`

Example:

<!-- test: div -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`

## mod

The modulus when *p0* is divided by *p1*. If *p1* is zero, a 'division by zero'
error is raised.

    ( p0:T p1:T -- T )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`

Example:

<!-- test: mod -->

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `mod`   | `1`

## mul

Multiplies *p0* by *p1*.

    ( p0:T p1:T -- T )

Aliases: `m`, `*`

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Example:

<!-- test: mul -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `m`     | `12`

## neg

Changes the sign of `p0`

    ( p0:T -- T )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`

Example:

<!-- test: neg -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `neg`   | `6`
| `neg`   | `-6`

## pow

Raises *a* to the power of *b*.

    ( a:T b:T -- T )

Where *T* is one of:
- `BigInt`
- `Float`
- `Complex`

Alias: `**`, `^`

Example:

<!-- test: pow -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `pow`   | `36`

## rem

The remainder when *p0* is divided by *p1*. If *p1* is zero, a
'division by zero' error is raised.

    ( p0:T p1:T -- T )

Where *T* is one of:
- `BigInt`
- `Float`

Example:

<!-- test: rem -->

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `rem`   | `-1`

## sign

Returns `-1` if *p0* is negative, `1` if *p0* is positive, or `0` if *p0*
is zero.

    ( p0:T -- Int )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`

Example:

<!-- test: sign -->

| Input     | Stack
|-----------|-------------|
| `-6 sign` | `-1`
| `clear`   |
| `6 sign`  | `1`
| `clear`   |
| `0 sign`  | `0`

## sqrt

The square root of `p0`.

    ( p0:Float -- T )

Where *T* is one of:
- `Float` if *p0* is positive or zero
- `Complex` if *p0* is negative.

<!-- test: sqrt -->

| Input     | Stack
|-----------|-------------|
| `256`     | `256`
| `sqrt`    | `16`

## sub

Subtract *p1* from *p0*.

    ( p0:Num p1:Num -- Num )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Aliases: `s`, `-`

Example:

<!-- test: sub -->

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`
