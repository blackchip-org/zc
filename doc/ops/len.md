<!-- Document generated by "gen-doc"; DO NOT EDIT -->
# len

Length conversions

| Operation    | Description
|--------------|---------------
| [`ft-m`](#ft-m) | Feet to meters
| [`ft-mi`](#ft-mi) | Miles to feet
| [`ft-yd`](#ft-yd) | Feet to yards
| [`in-mm`](#in-mm) | Inches to millimeters
| [`km-m`](#km-m) | Kilometers to meters
| [`km-mi`](#km-mi) | Kilometers to miles
| [`km-nmi`](#km-nmi) | Kilometers to nautical miles
| [`m-ft`](#m-ft) | Meters to feet
| [`m-km`](#m-km) | Meters to kilometers
| [`m-nmi`](#m-nmi) | Meters to nautical miles
| [`m-yd`](#m-yd) | Meters to yards
| [`mi-ft`](#mi-ft) | Miles to feet
| [`mi-km`](#mi-km) | Miles to kilometers
| [`mi-nmi`](#mi-nmi) | Miles to nautical miles
| [`mm-in`](#mm-in) | Millimeters to inches
| [`nmi-km`](#nmi-km) | Nautical miles to kilometers
| [`nmi-m`](#nmi-m) | Nautical miles to meters
| [`nmi-mi`](#nmi-mi) | Nautical miles to miles
| [`yd-ft`](#yd-ft) | Yards to feet
| [`yd-m`](#yd-m) | Yards to meters


## ft-m

Convert *p0* in feet to meters.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: ft-m -->

| Input       | Stack
|-------------|---------------
| `1000 ft-m` | `304.8 # m`

## ft-mi

Convert *p0* in feet to miles

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: ft-mi -->

| Input        | Stack
|--------------|---------------
| `2640 ft-mi` | `0.5 # mi`

## ft-yd

Convert *p0* in feet to yards.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: ft-yd -->

| Input       | Stack
|-------------|---------------
| `300 ft-yd` | `100 # yd`

## in-mm

Convert *p0* in inches to millimeters.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: in-mm -->

| Input       | Stack
|-------------|---------------
| `100 in-mm` | `2540 # mm`

## km-m

Convert *p0* in kilometers to meters.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: km-m -->

| Input           | Stack
|-----------------|---------------
| `6378.137 km-m` | `6378137 # m`

## km-mi

Convert *p0* in kilometers to miles.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: km-mi -->

| Input               | Stack
|---------------------|---------------
| `100 km-mi 2 round` | `62.14 # mi`

## km-nmi

Convert *p0* in kilometers to nautical miles.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: km-nmi -->

| Input                | Stack
|----------------------|---------------
| `100 km-nmi 2 round` | `54 # nmi`

## m-ft

Convert *p0* in meters to feet.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: m-ft -->

| Input                | Stack
|----------------------|---------------
| `304.8 m-ft 2 round` | `1000 # ft`

## m-km

Convert *p0* in meters to kilometers.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: m-km -->

| Input                     | Stack
|---------------------------|---------------
| `earth-equatorial-radius` | `6378137 # m`
| `m-km                   ` | `6378.137 # km`

## m-nmi

Convert *p0* in meters to nautical miles.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: m-nmi -->

| Input                   | Stack
|-------------------------|---------------
| `100,000 m-nmi 2 round` | `54 # nmi`

## m-yd

Convert *p0* in meters to yards

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: m-yd -->

| Input                | Stack
|----------------------|---------------
| `91.44 m-yd 2 round` | `100 # yd`

## mi-ft

Convert *p0* in miles to feet

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: mi-ft -->

| Input       | Stack
|-------------|---------------
| `0.5 mi-ft` | `2640 # ft`

## mi-km

Convert *p0* in miles to kilometers

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: mi-km -->

| Input               | Stack
|---------------------|---------------
| `100 mi-km 2 round` | `160.93 # km`

## mi-nmi

Convert *p0* in miles to nautical miles

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: mi-nmi -->

| Input                | Stack
|----------------------|---------------
| `100 mi-nmi 2 round` | `86.9 # nmi`

## mm-in

Convert *p0* in millimeters to inches.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: mm-in -->

| Input                | Stack
|----------------------|---------------
| `2540 mm-in 2 round` | `100 # in`

## nmi-km

Convert *p0* in nautical miles to kilometers

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: nmi-km -->

| Input                | Stack
|----------------------|---------------
| `100 nmi-km 2 round` | `185.2 # km`

## nmi-m

Convert *p0* in nautical miles to meters

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: nmi-m -->

| Input       | Stack
|-------------|---------------
| `100 nmi-m` | `185200 # m`

## nmi-mi

Convert *p0* in nautical miles to miles

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: nmi-mi -->

| Input                | Stack
|----------------------|---------------
| `100 nmi-mi 2 round` | `115.08 # mi`

## yd-ft

Convert *p0* in yards to feet

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: yd-ft -->

| Input       | Stack
|-------------|---------------
| `100 yd-ft` | `300 # ft`

## yd-m

Convert *p0* in yards to meters

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: yd-m -->

| Input      | Stack
|------------|---------------
| `100 yd-m` | `91.44 # m`
