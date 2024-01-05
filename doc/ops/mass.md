<!-- Document generated by "gen-doc"; DO NOT EDIT -->
# mass

Mass conversions

| Operation  | Description
|------------|---------------
| [`g-oz`](#g-oz) | Grams to ounces
| [`kg-lb`](#kg-lb) | Kilograms to pounds
| [`lb-kg`](#lb-kg) | Pounds to kilograms
| [`oz-g`](#oz-g) | Ounces to grams


## g-oz

Convert *p0* in grams to ounces.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: g-oz -->

| Input                  | Stack
|------------------------|---------------
| `2834.95 g-oz 2 round` | `100`

## kg-lb

Convert *p0* in kilograms to pounds.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: kg-lb -->

| Input                 | Stack
|-----------------------|---------------
| `45.36 kg-lb 2 round` | `100`

## lb-kg

Convert *p0* in pounds to kilograms.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: lb-kg -->

| Input               | Stack
|---------------------|---------------
| `100 lb-kg 2 round` | `45.36`

## oz-g

Convert *p0* in ounces to grams.

```
( p0:Decimal -- Decimal )
```

Example:

<!-- test: oz-g -->

| Input              | Stack
|--------------------|---------------
| `100 oz-g 2 round` | `2834.95`