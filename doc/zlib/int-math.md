# int-math

Basic mathematical operations with integers.

- Use: import

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


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:BigInt -- abs:BigInt )

Example:

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `abs`   | `6`


## add

Adds the value of `b` to `a`.

    ( a:BigInt b:BigInt -- add:BigInt )

Aliases: `a`, `+`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`


## ceil

The identity operation with integers.

    ( a:BigInt -- a:BigInt )

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `floor` | `6`


## div

Divides the value of `a` by `b`.

    ( a:BigInt b:BigInt -- div:BigInt )

Aliases: `d`, `/`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`


## floor

The identity operation with integers.

    ( a:BigInt -- a:BigInt )

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `floor` | `6`


## mod

The modulus when `a` is divided by `b`.

    ( a:BigInt b:BigInt -- mod:BigInt )

Example:

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `mod`   | `1`


## neg

Changes the sign of `a`.

    ( a:BigInt -- neg:BigInt )

Example:

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `neg`   | `6`
| `neg`   | `-6`


## mul

Multiplies `a` by `b`.

    ( a:BigInt b:BigInt -- mul:Big )

Aliases: `m`, `*`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `m`     | `12`


## pow

Raises `a` to the power of `b`.

    ( a:BigInt b:BigInt -- pow:BigInt )

Alias: `**`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `pow`   | `36`


## rem

The remainder when `a` is divided by `b`.

    ( a:BigInt b:BigInt -- rem:BigInt )

Example:

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
( a:BigInt -- sign:BigInt )
```

Example:

| Input     | Stack
|-----------|-------------|
| `-6 sign` | `-1`
| `clear`   |
| `6 sign`  | `1`
| `clear`   |
| `0 sign`  | `0`


## sub

Subtracts `b` from `a`.

    ( a:BigInt b:BigInt -- sub:BigInt )

Aliases: `s`, `-`

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`