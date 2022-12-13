# math

Basic mathematical operations.

- Prelude: user, dev
- Use: include

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [abs](#abs)             |          | Absolute value
| [add](#add)             | `a`, `+` | Addition
| [ceil](#ceil)           |          | Ceiling
| [dec](#dec)             | `--`     | Decrement
| [div](#div)             | `d`, `/` | Division
| [floor](#floor)         |          | Floor
| [inc](#inc)             | `++`     | Increment
| [mod](#mod)             |          | Modulus
| [mul](#mul)             | `m`, `*` | Multiplication
| [neg](#neg)             |          | Negation
| [pow](#pow)             | `**`     | Exponentiation
| [rem](#rem)             |          | Remainder
| [sign](#sign)           |          | Sign
| [sub](#sub)             | `s`, `-` | Subtraction


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:Num -- abs:Num )

Example:

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `abs`   | `6`


## add

Adds the value of `b` to `a`.

    ( a:Num b:Num -- add:Num )

Aliases: `a`, `+`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`


## ceil

The nearest integer value greater than or equal to `a`.

    ( a:Num -- ceil:Num )

Example:

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `ceil`  | `7`


## dec

Decrements the value of `a` by `1`.

    ( a:Num -- dec:Num )

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `dec`   | `5`


## div

Divides the value of `a` by `b`.

    ( a:Num b:Num -- div:Num )

Aliases: `d`, `/`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`


## floor

The nearest integer value less than or equal to `a`.

    ( a:Num -- floor:Num )

Example:

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `floor`  | `6`


## inc

Increments the value of `a` by `1`.

    ( a:Num -- inc:Num )

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `inc`   | `7`


## mod

The modulus when `a` is divided by `b`.

    ( a:Num b:Num -- mod:Num )

Example:

| Input   | Stack
|---------|-------------|
| `-7`    | `-7`
| `2`     | `-7 \| 2`
| `mod`   | `1`


## neg

Changes the sign of `a`.

    ( a:Num -- neg:Num )

Example:

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

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `pow`   | `36`


## rem

The remainder when `a` is divided by `b`.

    ( a:Num b:Num -- rem:Num )

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
( a:Num -- sign:Int )
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

    ( a:Num b:Num -- sub:Num )

Aliases: `s`, `-`

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`