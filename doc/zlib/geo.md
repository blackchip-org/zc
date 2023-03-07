# geo

<!-- eval: use geo -->

Geo-spatial calculations

    use geo

<!-- index -->

| Operation                     | Description
|-------------------------------|----------------
| [proj](#proj)                 | Transform coordinates


## proj

Transform coordinate `(a, b)` in coordinate system `source` to coordinate
system `target`. The order of the coordinates is defined by the coordinate
system and it may be `(lat, lon)` or `(x, y)`.

    ( a:Float b:Float source:Str target:Str -- ta:Float tb:Float )


Example:

<!-- test: proj -->

| Input                           | Stack
|---------------------------------|---------------------
| `import epsg`                   | *imported epsg*
| `39.203611 -76.856944`          | `39.203611 \| -76.856944`
| `epsg.wgs-84 epsg.utm-18n proj` | `339660.125593429 \| 4341014.551927999`




