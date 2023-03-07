# color

<!-- eval: use color -->

Color calculations

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [cmyk-rgb](#cmyk-rgb)   | Convert from CMYK to RGB color space
| [rbg-cmyk](#rgb-cmyk)   | Convert from RGB to CMYK color space
| [sample](#sample)       | Render a sample of a color


## cmyk-rgb

Convert color CMYK color `c`, `m`, `y`, `k`, to a RGB color space as `r`, `b`,
`g`

    ( c:Uint8 m:Uint8 y:Uint8 k:Uint8 -- r:Uint8 g:Uint8 b:Uint8 )

Example:

<!-- test: cmyk-rgb -->

| Input           | Stack
|-----------------|-------------
| `0 127 191 127` | `0 \| 127 \| 191 \| 127`
| `cmyk-rgb`      | `128 \| 64 \| 32`


## rgb-cmyk

Convert color RGB color `r`, `b`', `g` to a CMYK color space as `c`, `m`,
`y`, `k`.

    ( r:Uint8 g:Uint8 b:Uint8 -- c:Uint8 m:Uint8 y:Uint8 k:Uint8 )

Example:

<!-- test: rgb-cmyk -->

| Input           | Stack
|-----------------|-------------
| `128 64 32`     | `128 \| 64 \| 32`
| `rgb-cmyk`      | `0 \| 127 \| 191 \| 127`

## sample

Render five spaces with a background color of `r`, `g`, and `b`. The
terminal in use must be able to support 24-bit ANSI color codes.

    ( r:Uint8 g:Uint8 b:Uint8 -- sample:Str )

