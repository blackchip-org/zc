<!-- mod: prog -->

# prog

Programmer's calculator.

- Use: include

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [and](#and)             | Bitwise and
| [bin](#bin)             | Convert to binary
| [bit](#bit)             | Bit value
| [dec](#dec)             | Convert to decimal
| [hex](#hex)             | Convert to hexadecimal
| [len](#len)             | Length in bits
| [lsh](#lsh)             | Left shift
| [not](#not)             | Bitwise not
| [oct](#oct)             | Convert to octal
| [or](#or)               | Bitwise or
| [rsh](#rsh)             | Right shift
| [xor](#xor)             | Bitwise exclusive or


## and

The bitwise and of `a` and `b`.

    ( a:BigInt b:BigInt -- and:BigInt )

Example:

<!-- test: and -->

| Input       | Stack
|-------------|-------------|
| `0b1100`    | `0b1100`
| `0b1010`    | `0b1100 \| 0b1010`
| `and`       | `0b1000`

## bin

Convert the value of `a` to a binary number.

    ( a:BigInt -- dec:BigInt )

Example:

<!-- test: bin -->

| Input       | Stack
|-------------|-------------|
| `0xf`       | `0xf`
| `bin`       | `0b1111`


## bit

The value of the `b`th bit of `a`.

    ( a:BigInt b:Int -- bit:Uint )

Example:

<!-- test: bit -->

| Input       | Stack
|-------------|-------------|
| `0b100`     | `0b100`
| `2 bit`     | `1`


## dec

Convert the value of `a` to a decimal number.

    ( a:BigInt -- dec:BigInt )

Example:

<!-- test: dec -->

| Input       | Stack
|-------------|-------------|
| `0xf`       | `0xf`
| `dec`       | `15`


## hex

Convert the value of `a` to a hexadecimal number.

    ( a:BigInt -- dec:BigInt )

Example:

<!-- test: hex -->

| Input       | Stack
|-------------|-------------|
| `0b1111`    | `0b1111`
| `hex`       | `0xf`


## len

The length of `a` in bits.

    ( a:BigInt -- len:Int )

Example:

<!-- test: len -->

| Input       | Stack
|-------------|-------------|
| `0b1111`    | `0b1111`
| `len`       | `4`

## lsh

Shifts all bits in `a` to the left by `b`

    ( a:BigInt b:Uint -- lsh:BigInt )

Example:

<!-- test: lsh -->

| Input       | Stack
|-------------|-------------|
| `0b10`      | `0b10`
| `2 lsh`     | `0b1000`


## not

The bitwise not of `a` and `b`.

    ( a:BigInt b:BigInt -- and:BigInt )

Example:

<!-- test: not -->

| Input       | Stack
|-------------|-------------|
| `0b101`     | `0b101`
| `not`       | `-0b110`


## oct

Convert the value of `a` to an octal number.

    ( a:BigInt -- dec:BigInt )

Example:

<!-- test: oct -->

| Input       | Stack
|-------------|-------------|
| `0b1111`    | `0b1111`
| `oct`       | `0o17`


## or

The bitwise or of `a` and `b`.

    ( a:BigInt b:BigInt -- and:BigInt )

Example:

<!-- test: or -->

| Input       | Stack
|-------------|-------------|
| `0b1100`    | `0b1100`
| `0b1010`    | `0b1100 \| 0b1010`
| `or`        | `0b1110`


## rsh

Shifts all bits in `a` to the right by `b`

    ( a:BigInt b:Uint -- lsh:BigInt )

Example:

<!-- test: rsh -->

| Input       | Stack
|-------------|-------------|
| `0b1000`    | `0b1000`
| `2 rsh`     | `0b10`


## xor

The bitwise exclusive of `a` and `b`.

    ( a:BigInt b:BigInt -- and:BigInt )

Example:

<!-- test: xor -->

| Input       | Stack
|-------------|-------------|
| `0b1100`    | `0b1100`
| `0b1010`    | `0b1100 \| 0b1010`
| `xor`       | `0b110`





