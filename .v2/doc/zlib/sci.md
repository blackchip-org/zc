# sci

<!-- eval: use sci -->

Scientific calculator

<!-- index -->

| Operation               | Alias     | Description
|-------------------------|-----------|-----------
| [acos](#acos)           |           | Inverse cosine
| [acosh](#acosh)         |           | Inverse hyperbolic cosine
| [asin](#asin)           |           | Inverse sine
| [asinh](#asinh)         |           | Inverse hyperbolic sine
| [atan](#atan)           |           | Inverse tangent
| [atanh](#atanh)         |           | Inverse hyperbolic tangent
| [cos](#cos)             |           | Cosine
| [cosh](#cosh)           |           | Hyperbolic cosine
| [e](#e)                 |           | Natural logarithm base
| [exp](#exp)             |           | Natural exponential
| [log](#log)             |           | Natural logarithm
| [log10](#log10)         |           | Decimal logarithm
| [log2](#log2)           |           | Binary logarithm
| [pi](#pi)               | `π`       | Circumference to diameter ratio
| [scientific-notation](#scientific-notation) | `sn` | Scientific notation
| [sin](#sin)             |           | Sine
| [sinh](#sinh)           |           | Hyperbolic sine
| [tan](#tan)             |           | Tangent
| [tanh](#tanh)           |           | Hyperbolic tangent

## acos

Inverse cosine in radians

    ( x:Float -- acos:Float )

Example:

<!-- test: acos -->

| Input           | Stack
|-----------------|-------------|
| `0.5 acos`      | `1.047197551196598`


## acosh

Inverse hyperbolic cosine in radians

    ( x:Float -- acosh:Float )

Example:

<!-- test: acosh -->

| Input           | Stack
|-----------------|-------------|
| `2 acosh`       | `1.316957896924817`


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
| `2 asinh`       | `1.44363547517881`


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
| `2 cosh`        | `3.762195691083631`


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


## log

Natural logarithm

    ( x:Float -- log:Float )

Example:

<!-- test: log -->

| Input           | Stack
|-----------------|-------------|
| `8 log`         | `2.079441541679836`


## log10

Decimal logarithm

    ( x:Float -- log10:Float )

Example:

<!-- test: log10 -->

| Input           | Stack
|-----------------|-------------|
| `50 log10`      | `1.698970004336019`


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
