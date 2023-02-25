# unit

<!-- eval: use unit -->

Units of measure

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [c-f](#c-f)             | Celsius to Fahrenheit
| [c-k](#c-k)             | Celsius to Kelvin
| [f-c](#f-c)             | Fahrenheit to Celsius
| [k-c](#k-c)             | Kelvin to Celsius
| [km-mi](#km-mi)         | Kilometers to miles
| [mi-km](#mi-km)         | Miles to kilometers


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


## km-mi

Kilometers to miles

    ( km:Number -- mi:Number )

Example:

<!-- test: km-mi -->

| Input           | Stack
|-----------------|-------------|
| `100`           | `100`
| `km-mi 3 round` | `62.137`


## mi-km

Miles to kilometers

    ( mi:Number -- km:Number )

Example:

<!-- test: mi-km -->

| Input           | Stack
|-----------------|-------------|
| `60`            | `60`
| `mi-km 3 round` | `96.561`


