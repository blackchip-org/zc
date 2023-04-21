# unit

Units of measure

- [len](#len): Length conversions
- [si](#si): Prefixes in the International System of Units
- [temp](#temp): Temperature conversions

# len

Units of length

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [km-mi](#km-mi)         | Kilometers to miles
| [km-nmi](#km-nmi)       | Kilometers to nautical miles
| [m-nmi](#m-nmi)         | Meters to nautical miles
| [mi-km](#mi-km)         | Miles to kilometers
| [mi-nm](#mi-nmi)        | Miles to nautical miles
| [nmi-km](#nmi-km)       | Nautical miles to kilometers
| [nmi-m](#nmi-m)         | Nautical miles to meters
| [nmi-mi](#nmi-mi)       | Nautical miles to miles

## km-mi

Kilometers to miles

    ( km:Num -- mi:Num )

Example:

<!-- test: km-mi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `km-mi 2 round`  | `62.14`


## km-nmi

Kilometers to nautical miles

    ( km:Num -- nmi:Num )

Example:

<!-- test: km-nmi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `km-nmi 2 round`  | `54`


## m-nmi

Meters to nautical miles

    ( m:Num -- nmi:Num )

Example:

<!-- test: m-nmi -->

| Input            | Stack
|------------------|-------------
| `100,000`        | `100,000`
| `m-nmi 2 round`   | `54`

## mi-km

Miles to kilometers

    ( mi:Num -- km:Num )

Example:

<!-- test: mi-km -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `mi-km 2 round`  | `160.93`


## mi-nmi

Miles to nautical miles

    ( mi:Num -- nmi:Num )

Example:

<!-- test: mi-nmi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `mi-nmi 2 round` | `86.9`


## nmi-km

Nautical miles to kilometers

    ( nmi:Num -- km:Num )

Example:

<!-- test: nmi-km -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nmi-km 2 round` | `185.2`

## nmi-km

Nautical miles to kilometers

    ( nmi:Num -- km:Num )

Example:

<!-- test: nmi-km -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nmi-km 2 round` | `185.2`

## nmi-m

Nautical miles to meters

    ( nmi:Num -- m:Num )

Example:

<!-- test: nmi-m -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nmi-m 2 round`   | `185200`

## nmi-mi

Nautical miles to miles

    ( nmi:Num -- mi:Num )

Example:

<!-- test: nmi-mi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nmi-mi 2 round` | `115.08`

# si

Prefixes for the International System of Units

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [quetta](#quetta)       | Q, 1e30
| [ronna](#ronna)         | R, 1e27
| [yotta](#yotta)         | Y, 1e24
| [zetta](#zetta)         | Z, 1e21
| [exa](#exa)             | E, 1e18
| [peta](#peta)           | P, 1e15
| [tera](#tera)           | T, 1e12
| [giga](#giga)           | G, 1e09
| [mega](#mega)           | M, 1e06
| [kilo](#kilo)           | k, 1e03
| [hecto](#hecto)         | h, 1e02
| [deca](#deca)           | da, 1e01
| [deci](#deci)           | d, 1e-01
| [centi](#centi)         | c, 1e-02
| [milli](#milli)         | m, 1e-03
| [micro](#micro)         | μ, 1e-06
| [nano](#nano)           | n, 1e-09
| [pico](#pico)           | p, 1e-12
| [femto](#femto)         | f, 1e-15
| [atto](#atto)           | a, 1e-18
| [zepto](#zepto)         | z, 1e-21
| [yocto](#yocto)         | y, 1e-24
| [ronto](#ronto)         | r, 1e-27
| [quecto](#quecto)       | q, 1e-30


## quetta

Q, 1e30

    ( -- '1e30' )

Example:

<!-- test: quetta -->

| Input           | Stack
|-----------------|-------------|
| `quetta`        | `1e30`

## ronna

R, 1e27

    ( -- '1e27' )

Example:

<!-- test: ronna -->

| Input          | Stack
|----------------|-------------|
| `ronna`        | `1e27`

## yotta

Y, 1e24

    ( -- '1e24' )

Example:

<!-- test: yotta -->

| Input          | Stack
|----------------|-------------|
| `yotta`        | `1e24`

## zetta

Z, 1e21

    ( -- '1e21' )

Example:

<!-- test: zetta -->

| Input          | Stack
|----------------|-------------|
| `zetta`        | `1e21`

## exa

E, 1e18

    ( -- '1e18' )

Example:

<!-- test: exa -->

| Input         | Stack
|---------------|-------------|
| `exa`         | `1e18`

## peta

P, 1e15

    ( -- '1e15' )

Example:

<!-- test: peta -->

| Input         | Stack
|---------------|-------------|
| `peta`        | `1e15`

## tera

T, 1e12

    ( -- '1e12' )

Example:

<!-- test: tera -->

| Input         | Stack
|---------------|-------------|
| `tera`        | `1e12`

## giga

G, 1e9

    ( -- '1e9' )

Example:

<!-- test: giga -->

| Input         | Stack
|---------------|-------------|
| `giga`        | `1e09`

## mega

M, 1e6

    ( -- '1e6' )

Example:

<!-- test: mega -->

| Input         | Stack
|---------------|-------------|
| `mega`        | `1e06`

## kilo

k, 1e3

    ( -- '1e3' )

Example:

<!-- test: kilo -->

| Input         | Stack
|---------------|-------------|
| `kilo`        | `1e03`

## hecto

h, 1e2

    ( -- '1e2' )

Example:

<!-- test: hecto -->

| Input         | Stack
|---------------|-------------|
| `hecto`       | `1e02`

## deca

da, 1e1

    ( -- '1e1' )

Example:

<!-- test: deca -->

| Input         | Stack
|---------------|-------------|
| `deca`        | `1e01`

## deci

d, 1e-1

    ( -- '1e-1' )

Example:

<!-- test: deci -->

| Input         | Stack
|---------------|-------------|
| `deci`        | `1e-01`

## centi

c, 1e-2

    ( -- '1e-2' )

Example:

<!-- test: deci -->

| Input         | Stack
|---------------|-------------|
| `centi`       | `1e-02`

## milli

d, 1e-3

    ( -- '1e-3' )

Example:

<!-- test: milli -->

| Input         | Stack
|---------------|-------------|
| `milli`       | `1e-03`

## micro

μ, 1e-6

    ( -- '1e-6' )

Example:

<!-- test: micro -->

| Input         | Stack
|---------------|-------------|
| `micro`       | `1e-06`

## nano

n, 1e-9

    ( -- '1e-9' )

Example:

<!-- test: nano -->

| Input         | Stack
|---------------|-------------|
| `nano`        | `1e-09`

## pico

p, 1e-12

    ( -- '1e-12' )

Example:

<!-- test: pico -->

| Input         | Stack
|---------------|-------------|
| `pico`        | `1e-12`

## femto

f, 1e-15

    ( -- '1e-15' )

Example:

<!-- test: femto -->

| Input         | Stack
|---------------|-------------|
| `femto`       | `1e-15`

## atto

a, 1e-18

    ( -- '1e-18' )

Example:

<!-- test: atto -->

| Input         | Stack
|---------------|-------------|
| `atto`        | `1e-18`

## zepto

z, 1e-21

    ( -- '1e-21' )

Example:

<!-- test: zepto -->

| Input         | Stack
|---------------|-------------|
| `zepto`       | `1e-21`

## yocto

y, 1e-24

    ( -- '1e-24' )

Example:

<!-- test: yocto -->

| Input         | Stack
|---------------|-------------|
| `yocto`       | `1e-24`

## ronto

r, 1e-27

    ( -- '1e-27' )

Example:

<!-- test: ronto -->

| Input         | Stack
|---------------|-------------|
| `ronto`       | `1e-27`

## quecto

q, 1e-30

    ( -- '1e-30' )

Example:

<!-- test: quecto -->

| Input         | Stack
|---------------|-------------|
| `quecto`      | `1e-30`


# temp

Temperature conversions

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [c-f](#c-f)             | Celsius to Fahrenheit
| [c-k](#c-k)             | Celsius to Kelvin
| [f-c](#f-c)             | Fahrenheit to Celsius
| [k-c](#k-c)             | Kelvin to Celsius


## c-f

Celsius to Fahrenheit

    ( c:Number -- f:Number )

Example:

<!-- test: c-f -->

| Input         | Stack
|---------------|-------------|
| `20`          | `20`
| `c-f`         | `68`

## c-k

Celsius to Kelvin

    ( c:Number -- k:Number )

Example:

<!-- test: c-k -->

| Input         | Stack
|---------------|-------------|
| `100`         | `100`
| `c-k`         | `373.15`

## f-c

Fahrenheit to Celsius

    ( f:Number -- c:Number )

Example:

<!-- test: f-c -->

| Input         | Stack
|---------------|-------------|
| `68`          | `68`
| `f-c 2 round` | `20`

## k-c

Kelvin to Celsius

    ( k:Number -- c:Number )

Example:

<!-- test: k-c -->

| Input         | Stack
|---------------|-------------|
| `373.15`      | `373.15`
| `k-c`         | `100`





