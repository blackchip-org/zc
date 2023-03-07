# geo

<!-- eval: use geo -->

Geo-spatial calculations

    use geo

<!-- index -->

| Operation                     | Alias | Description
|-------------------------------|-------|----------------
| [transform](#transform)       | `tf`  | Transform coordinates


## transform

Transform coordinate `(a, b)` in coordinate system `source` to coordinate
system `target`. The order of the coordinates is defined by the coordinate
system and it may be `(lat, lon)` or `(x, y)`.

    ( a:Float b:Float source:Str target:Str -- ta:Float tb:Float )


Alias: `tf`

Example:

<!-- test: transform -->

| Input                           | Stack
|---------------------------------|---------------------
| `import epsg`                   | *imported epsg*
| `39.203611 -76.856944`          | `39.203611 \| -76.856944`
| `epsg.wgs-84`                   | `39.203611 \| -76.856944 \| EPSG:4326`
| `18n epsg.utm`                  | `39.203611 \| -76.856944 \| EPSG:4326 \| EPSG:32618`
| `transform`                     | `339660.125593429 \| 4341014.551927999`




