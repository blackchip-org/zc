# fix-math

Basic mathematical operations with fixed-point numbers.

- Use: import

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [abs](#abs)             |          | Absolute value
| [add](#add)             | `a`, `+` | Addition
| [ceil](#ceil)           |          | Ceiling
| [div](#div)             | `d`, `/` | Division
| [floor](#floor)         |          | Floor
| [inc](#inc)             | `++`     | Increment
| [mod](#mod)             |          | Modulus
| [mul](#mul)             | `m`, `*` | Multiplication
| [neg](#neg)             |          | Negation
| [pow](#pow)             | `**`     | Exponentiation
| [sign](#sign)           |          | Sign
| [sub](#sub)             | `s`, `-` | Subtraction


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:Fix -- abs:Fix )

Example:

| Input   | Stack
|---------|-------------|
| `-6.6`  | `-6.6`
| `abs`   | `6.6`


## add

Adds the value of `b` to `a`.

    ( a:Fix b:Fix -- add:Fix )

Aliases: `a`, `+`

Example:

| Input   | Stack
|---------|-------------|
| `6.6`   | `6.6`
| `2.2`   | `6.6 \| 2.2`
| `a`     | `8.8`


## ceil

The nearest integer value greater than or equal to `a`.

    ( a:Fix -- ceil:Fix )

Example:

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `ceil`  | `7`


## div

Divides the value of `a` by `b`.

    ( a:Fix b:Fix -- div:Fix )

Aliases: `d`, `/`

Example:

| Input   | Stack
|---------|-------------|
| `6.6`   | `6.6`
| `2.2`   | `6.6 \| 2.2`
| `d`     | `3`


## floor

The nearest integer value less than or equal to `a`.

    ( a:Fix -- floor:Fix )

Example:

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `floor` | `6`


## mod

The modulus when `a` is divided by `b`.

    ( a:Fix b:Fix -- mod:Fix )

Example:

| Input   | Stack
|---------|-------------|
| `-7.7`  | `-7.7`
| `2`     | `-7.7 \| 2`
| `mod`   | `-1.7`


## neg

Changes the sign of `a`.

    ( a:Fix -- neg:Fix )

Example:

| Input   | Stack
|---------|-------------|
| `-6.6`  | `-6.6`
| `neg`   | `6.6`
| `neg`   | `-6.6`


## mul

Multiplies `a` by `b`.

    ( a:Fix b:Fix -- Num )

Aliases: `m`, `*`

Example:

| Input   | Stack
|---------|-------------|
| `6.6`   | `6.6`
| `2.2`   | `6.6 \| 2.2`
| `m`     | `14.52`


## pow

Raises `a` to the power of `b`.

    ( a:Fix b:Fix -- Num )

Alias: `**`

Example:

| Input   | Stack
|---------|-------------|
| `6.6`   | `6.6`
| `2`     | `6.6 \| 2`
| `pow`   | `43.56`


## sign

If:

* `a` is negative: `-1`
* `a` is positive: `1`
* `a` is zero: `0`

```
( a:Fix -- sign:Fix )
```

Example:

| Input       | Stack
|-------------|-------------|
| `-6.6 sign` | `-1`
| `clear`     |
| `6.6 sign`  | `1`
| `clear`     |
| `0.0 sign`  | `0`


## sub

Subtracts `b` from `a`.

    ( a:Fix b:Fix -- sub:Fix )

Aliases: `s`, `-`

| Input         | Stack
|---------------|-------------|
| `6.6`         | `6.6`
| `2.2`         | `6.6 \| 2.2`
| `s`           | `4.4`