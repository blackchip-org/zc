# basic

Basic mathematical operations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [add](#add)             | `a`, `+` | Addition
| [div](#div)             | `d`, `/` | Division
| [mod](#mod)             |          | Modulus
| [mul](#mul)             | `m`, `*` | Multiplication
| [neg](#neg)             |          | Negation
| [pow](#pow)             | `**`     | Exponentiation
| [rem](#rem)             |          | Remainder
| [sign](#sign)           |          | Sign
| [sub](#sub)             | `s`, `-` | Subtraction


## add

Adds the value of `b` to `a`.

    ( a:Num b:Num -- add:Num )

Aliases: `a`, `+`

Example:

<!-- test: add -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`

## div

Divides the value of `a` by `b`.

    ( a:Num b:Num -- div:Num )

Aliases: `d`, `/`

Example:

<!-- test: div -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`

## mod

The modulus when `a` is divided by `b`.

    ( a:Num b:Num -- mod:Num )

Example:

<!-- test: mod -->

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `mod`   | `1`

## neg

Changes the sign of `a`.

    ( a:Num -- neg:Num )

Example:

<!-- test: neg -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `neg`   | `6`
| `neg`   | `-6`

## mul

Multiplies `a` by `b`.

    ( a:Num b:Num -- Num )

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

    ( a:Num b:Num -- Num )

Alias: `**`

Example:

<!-- test: pow -->

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `pow`   | `36`

## rem

The remainder when `a` is divided by `b`.

    ( a:Num b:Num -- rem:Num )

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
( a:Num -- sign:Int )
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

## sub

Subtracts `b` from `a`.

    ( a:Num b:Num -- sub:Num )

Aliases: `s`, `-`

Example:

<!-- test: sub -->

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`
