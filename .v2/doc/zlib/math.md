# math

<!-- eval: use math -->

Basic mathematical operations.

    use math -- user prelude

<!-- main: abs -->
<!-- main: add -->
<!-- main: ceil -->
<!-- main: div -->
<!-- main: floor -->
<!-- main: mod -->
<!-- main: mul -->
<!-- main: neg -->
<!-- main: pow -->
<!-- main: rem -->
<!-- main: sign -->
<!-- main: sub -->

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [abs](#abs)             |          | Absolute value
| [add](#add)             | `a`, `+` | Addition
| [ceil](#ceil)           |          | Ceiling
| [div](#div)             | `d`, `/` | Division
| [floor](#floor)         |          | Floor
| [mod](#mod)             |          | Modulus
| [mul](#mul)             | `m`, `*` | Multiplication
| [neg](#neg)             |          | Negation
| [pow](#pow)             | `**`     | Exponentiation
| [rem](#rem)             |          | Remainder
| [sign](#sign)           |          | Sign
| [sub](#sub)             | `s`, `-` | Subtraction
| [sum](#sum)             |          | Sum


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:Num -- abs:Num )

Example:

<!-- test: abs -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `abs`   | `6`


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


## ceil

The nearest integer value greater than or equal to `a`.

    ( a:Num -- ceil:Num )

Example:

<!-- test: ceil -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `ceil`  | `7`


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


## floor

The nearest integer value less than or equal to `a`.

    ( a:Num -- floor:Num )

Example:

<!-- test: floor -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `floor`  | `6`


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


## sum

The sum of all items on the stack.

    ( ...:Num -- sum:Num )

Example:

<!-- test: sum -->

| Input         | Stack
|---------------|-------------|
| `1`           | `1`
| `2`           | `1 \| 2`
| `3`           | `1 \| 2 \| 3`
| `4`           | `1 \| 2 \| 3 \| 4`
| `sum`         | `10`
