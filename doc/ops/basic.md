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

Adds the value of `b` to `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Rational b:Rational -- add:Rational ); or
    ( a:Complex  b:Complex  -- add:Complex )

Aliases: `a`, `+`

Example:

<!-- test: add -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`

## div

Divides the value of `a` by `b`. If `b` is zero, a 'division by zero' error is
set.

    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Rational b:Rational -- add:Rational ); or
    ( a:Complex  b:Complex  -- add:Complex )

Aliases: `d`, `/`

Example:

<!-- test: div -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`

## mod

The modulus when `a` is divided by `b`. If `b` is zero, a 'division by zero'
error is set.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );

Example:

<!-- test: mod -->

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `mod`   | `1`

## neg

Changes the sign of `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Rational b:Rational -- add:Rational )

Example:

<!-- test: neg -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `neg`   | `6`
| `neg`   | `-6`

## mul

Multiplies `a` by `b`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Rational b:Rational -- add:Rational ); or
    ( a:Complex  b:Complex  -- add:Complex )

Aliases: `m`, `*`

Example:

<!-- test: mul -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `m`     | `12`

## pow

Raises `a` to the power of `b`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Complex  b:Complex  -- add:Complex )

Alias: `**`, `^`

Example:

<!-- test: pow -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `pow`   | `36`

## rem

The remainder when `a` is divided by `b`. If `b` is zero, a 'division by zero' error is set.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Float    b:Float    -- add:Float );

Example:

<!-- test: rem -->

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `rem`   | `-1`

## sign

If:

* `a` is negative: `-1`
* `a` is positive: `1`
* `a` is zero: `0`

```
( a:BigInt   b:BigInt   -- add:BigInt );   or
( a:Decimal  b:Decimal  -- add:Decimal );  or
( a:Float    b:Float    -- add:Float );    or
( a:Rational b:Rational -- add:Rational ); or
```

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

The square root of `a0`. Returns a Float if `a0` is equal to or greater than
zero, otherwise returns a Complex.

    ( a0:Float -- r0:Float ); or
    ( a0:Float -- r0:Complex)

<!-- test: sqrt -->

| Input     | Stack
|-----------|-------------|
| `256`     | `256`
| `sqrt`    | `16`

## sub

Subtracts `b` from `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or
    ( a:Rational b:Rational -- add:Rational ); or
    ( a:Complex  b:Complex  -- add:Complex )

Aliases: `s`, `-`

Example:

<!-- test: sub -->

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`
