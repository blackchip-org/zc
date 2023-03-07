# epsg

<!-- eval: import epsg -->

EPSG Geodetic Parameter Dataset

    import epsg

<!-- index -->

| Operation                     | Alias | Description
|-------------------------------|-------|--------
| [utm](#utm)                   |       | EPSG codes for Universal Transverse Mercator (WGS-84)
| [web-mercator](#web-mercator) | wmerc | Web Mercator, EPSG:3857
| [wgs-84](#wgs-84)             |       | World Geodetic System of 1984, EPSG:4326


## utm

The EPGS code for the given UTM zone `z`. The zone should be a number between
1 and 60 inclusive and is followed by a hemisphere designator of `n` or `s`.

    ( z:Str -- epsg:Str )

Example:

<!-- test: utm -->

| Input               | Stack
|---------------------|---------------------|
| `17n epsg.utm`      | `EPSG:32617`


## web-mercator

EPSG:3857

    ( -- epsg:Str )

Alias: `wmerc`

<!-- test: web-mercator -->

| Input               | Stack
|---------------------|---------------------|
| `epsg.web-mercator` | `EPSG:3857`


## wgs-84

EPSG:4326

    ( -- epsg:Str )

<!-- test: wgs-84 -->

| Input               | Stack
|---------------------|---------------------|
| `epsg.wgs-84`       | `EPSG:4326`

