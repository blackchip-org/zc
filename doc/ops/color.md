<!-- Document generated by "gen-doc"; DO NOT EDIT -->
# color

Color conversions

| Operation                | Description
|--------------------------|---------------
| [`cmyk-rgb`](#cmyk-rgb)  | CMYK to RGB color space
| [`color-sample`](#color-sample) | Render a sample of a color
| [`hsl-rgb`](#hsl-rgb)    | HSL to RGB color space
| [`rgb-cmyk`](#rgb-cmyk)  | RGB to CMYK color space
| [`rgb-hsl`](#rgb-hsl)    | RGB to HSL color space


## cmyk-rgb

Convert the CMYK color *c*, *m*, *y*, *k*, to the RGB color space as *r*, *b*,
*g*.

```
( c:Uint8 m:Uint8 y:Uint8 k:Uint8 -- r:Uint8 g:Uint8 b:Uint8 )
```

Example:

<!-- test: cmyk-rgb -->

| Input           | Stack
|-----------------|---------------
| `0 127 191 127` | `0 \| 127 \| 191 \| 127`
| `cmyk-rgb     ` | `128 \| 64 \| 32`

## color-sample

Render five spaces with a background color of *r*, *g*, and *b*. The
terminal in use must be able to support 24-bit ANSI color codes.

```
( r:Uint8 g:Uint8 b:Uint8 -- Str )
```


## hsl-rgb

Convert the HSL color *h*, *s*, *l* to the RGB color space as *r*, *g*, *b*.

```
( h:Float s:Float l:Float -- r:Uint8 g:Uint8 b:Uint8 )
```

Example:

<!-- test: hsl-rgb -->

| Input          | Stack
|----------------|---------------
| `20 0.6 0.314` | `20 \| 0.6 \| 0.314`
| `hsl-rgb     ` | `128 \| 64 \| 32`

## rgb-cmyk

Convert the RGB color *r*, *b*, *g* to the CMYK color space as *c*, *m*, *y*,
*k*.

```
( r:Uint8 g:Uint8 b:Uint8 -- c:Uint8 m:Uint8 y:Uint8 k:Uint8 )
```

Example:

<!-- test: rgb-cmyk -->

| Input       | Stack
|-------------|---------------
| `128 64 32` | `128 \| 64 \| 32`
| `rgb-cmyk ` | `0 \| 127 \| 191 \| 127`

## rgb-hsl

Convert the RGB color *r*, *g*, *b* to the HSL color space as *h*, *s*, *l*.

```
( r:Uint8 g:Uint8 b:Uint8 -- h:Float s:Float l:Float )
```

Example:

<!-- test: rgb-hsl -->

| Input                   | Stack
|-------------------------|---------------
| `128 64 32            ` | `128 \| 64 \| 32`
| `rgb-hsl [3 round] map` | `20 \| 0.6 \| 0.314`
