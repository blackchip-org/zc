# unit (len)

<!-- eval: use unit -->

Length conversions

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [km-mi](#km-mi)         | Kilometers to miles
| [km-nm](#km-nm)         | Kilometers to nautical miles
| [m-nm](#m-nm)           | Meters to nautical miles
| [mi-km](#mi-km)         | Miles to kilometers
| [mi-nm](#mi-nm)         | Miles to nautical miles
| [nm-km](#nm-km)         | Nautical miles to kilometers
| [nm-m](#nm-m)           | Nautical miles to meters
| [nm-mi](#nm-mi)         | Nautical miles to miles


## km-mi

Kilometers to miles

    ( km:Num -- mi:Num )

Example:

<!-- test: km-mi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `km-mi 2 round`  | `62.14`


## km-nm

Kilometers to nautical miles

    ( km:Num -- nm:Num )

Example:

<!-- test: km-nm -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `km-mi 2 round`  | `54`


## m-nm

Meters to nautical miles

    ( m:Num -- nm:Num )

Example:

<!-- test: m-nm -->

| Input            | Stack
|------------------|-------------
| `100,000`        | `100000`
| `m-mi 2 round`   | `54`


## mi-km

Miles to kilometers

    ( mi:Num -- km:Num )

Example:

<!-- test: mi-km -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `mi-km 2 round`  | `160.93`


## mi-nm

Miles to nautical miles

    ( mi:Num -- nm:Num )

Example:

<!-- test: mi-nm -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `mi-nm 2 round`  | `86.9`


## nm-km

Nautical miles to kilometers

    ( nm:Num -- km:Num )

Example:

<!-- test: nm-km -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nm-km 2 round`  | `185.2`


## nm-m

Nautical miles to meters

    ( nm:Num -- m:Num )

Example:

<!-- test: nm-m -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nm-m 2 round`   | `185200`


## nm-mi

Nautical miles to miles

    ( nm:Num -- mi:Num )

Example:

<!-- test: nm-mi -->

| Input            | Stack
|------------------|-------------
| `100`            | `100`
| `nm-mi 2 round`  | `115.08`