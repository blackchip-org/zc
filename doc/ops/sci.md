# sci

Advanced math operations found on scientific calculators.

<!-- index -->

| Operation               | Alias     | Description
|-------------------------|-----------|-----------
| [abs](#abs)             |           | Absolute value
| [acos](#acos)           |           | Inverse cosine
| [acosh](#acosh)         |           | Inverse hyperbolic cosine
| [asin](#asin)           |           | Inverse sine
| [asinh](#asinh)         |           | Inverse hyperbolic sine
| [atan](#atan)           |           | Inverse tangent
| [atanh](#atanh)         |           | Inverse hyperbolic tangent
| [ceil](#ceil)           |           | Ceiling
| [complex](#complex)     |           | Complex number
| [cos](#cos)             |           | Cosine
| [cosh](#cosh)           |           | Hyperbolic cosine
| [e](#e)                 |           | Natural logarithm base
| [exp](#exp)             |           | Natural exponential
| [floor](#floor)         |           | Floor
| [log](#log)             |           | Natural logarithm
| [log10](#log10)         |           | Decimal logarithm
| [log2](#log2)           |           | Binary logarithm
| [pi](#pi)               | `π`       | Circumference to diameter ratio
| [scientific-notation](#scientific-notation) | `sn` | Scientific notation
| [sin](#sin)             |           | Sine
| [sinh](#sinh)           |           | Hyperbolic sine
| [tan](#tan)             |           | Tangent
| [tanh](#tanh)           |           | Hyperbolic tangent


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:BigInt   -- abs:BigInt );   or
    ( a:Decimal  -- abs:Decimal );  or
    ( a:Float    -- abs:Float );    or
    ( a:Rational -- abs:Rational );

Example:

<!-- test: abs -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `abs`   | `6`

The distance of `a` from zero in the complex plane.

    ( a:Complex -- abs:Complex )

Example:

<!-- test: abs-complex -->

| Input   | Stack
|---------|-------------|
| `3+4i`  | `3+4i`
| `abs`   | `5`

## acos

Inverse cosine in radians

    ( x:Float -- acos:Float )

Example:

<!-- test: acos -->

| Input           | Stack
|-----------------|-------------|
| `0.5 acos`      | `1.0471975511965976`


## acosh

Inverse hyperbolic cosine in radians

    ( x:Float -- acosh:Float )

Example:

<!-- test: acosh -->

| Input           | Stack
|-----------------|-------------|
| `2 acosh`       | `1.3169578969248166`


## asin

Inverse sine in radians

    ( x:Float -- asin:Float )

Example:

<!-- test: asin -->

| Input           | Stack
|-----------------|-------------|
| `0.5 asin`      | `0.5235987755982989`


## asinh

Inverse hyperbolic sine in radians

    ( x:Float -- asinh:Float )

Example:

<!-- test: asinh -->

| Input           | Stack
|-----------------|-------------|
| `2 asinh`       | `1.4436354751788103`


## atan

Inverse tangent in radians

    ( x:Float -- atan:Float )

Example:

<!-- test: atan -->

| Input           | Stack
|-----------------|-------------|
| `0.5 atan`      | `0.4636476090008061`


## atanh

Inverse hyperbolic tangent in radians

    ( x:Float -- atanh:Float )

Example:

<!-- test: atanh -->

| Input           | Stack
|-----------------|-------------|
| `0.5 atanh`     | `0.5493061443340548`

## ceil

The nearest integer value greater than or equal to `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or

Example:

<!-- test: ceil -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `ceil`  | `7`

## complex

A complex number from a real `r` and an imaginary `i` numbers.

    ( r:Float i:Float -- r0:Complex )

Example:

<!-- test: complex -->

| Input     | Stack
|-----------|-------------|
| `6`       | `6`
| `12`      | `6 \| 12`
| `complex` | `6+12i`

## cos

Cosine in radians

    ( x:Float -- cos:Float )

Example:

<!-- test: cos -->

| Input           | Stack
|-----------------|-------------|
| `2 cos`         | `-0.4161468365471424`

## cosh

Hyperbolic cosine in radians

    ( x:Float -- cosh:Float )

Example:

<!-- test: cosh -->

| Input           | Stack
|-----------------|-------------|
| `2 cosh`        | `3.7621956910836314`

## e

Euler's number, the natural logarithm base.

    ( -- e:Float )

Example:

<!-- test: e -->

| Input           | Stack
|-----------------|-------------|
| `e`             | `2.718281828459045`

## exp

Natural exponential

    ( x:Float -- exp:Float )

Example:

<!-- test: exp -->

| Input           | Stack
|-----------------|-------------|
| `2 exp`         | `7.38905609893065`

## floor

The nearest integer value less than or equal to `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or

Example:

<!-- test: floor -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `floor` | `6`

## log

Natural logarithm

    ( x:Float -- log:Float )

Example:

<!-- test: log -->

| Input           | Stack
|-----------------|-------------|
| `8 log`         | `2.0794415416798357`

## log10

Decimal logarithm

    ( x:Float -- log10:Float )

Example:

<!-- test: log10 -->

| Input           | Stack
|-----------------|-------------|
| `50 log10`      | `1.6989700043360187`

## log2

Binary logarithm

    ( x:Float -- log10:Float )

Example:

<!-- test: log2 -->

| Input           | Stack
|-----------------|-------------|
| `250 log2`      | `7.965784284662087`

## pi

Circumference to diameter ratio of a circle

    ( -- pi:Float )

Alias: `π`

<!-- test: pi -->

| Input           | Stack
|-----------------|-------------|
| `pi`            | `3.141592653589793`

## sin

Sine in radians

    ( x:Float -- sin:Float )

Example:

<!-- test: sin -->

| Input           | Stack
|-----------------|-------------|
| `2 sin`         | `0.9092974268256816`

## sinh

Hyperbolic sine in radians

    ( x:Float -- sinh:Float )

Example:

<!-- test: sinh -->

| Input           | Stack
|-----------------|-------------|
| `2 sinh`        | `3.626860407847019`

## tan

Tangent in radians

    ( x:Float -- tan:Float )

Example:

<!-- test: tan -->

| Input           | Stack
|-----------------|-------------|
| `2 tan`         | `-2.185039863261519`

## tanh

Hyperbolic tangent in radians

    ( x:Float -- tanh:Float )

Example:

<!-- test: tanh -->

| Input           | Stack
|-----------------|-------------|
| `2 tanh`        | `0.9640275800758169`
