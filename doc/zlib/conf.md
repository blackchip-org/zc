# conf

Configurations and settings

- Prelude: user, dev
- Use: import

| Operation                   | Description
|-----------------------------|----------------
| [places](#places)           | Sets the number of places after the decimal point
| [places=](#places)          | Gets the number of places after the decimal point


## places

Sets the number of places to use after the decimal point to `a`

    ( a:Int32 -- )

Example:

| Input            | Stack
|------------------|---------------------|
| `2 3 div`        | `0.6666666666666667`
| `clear`          |
| `2 conf.places`  |
| `2 3 div`        | `0.67`


## places=

Gets the number of places after the decimal point

    ( -- places:Int32 )

Example:

| Input            | Stack
|------------------|---------------------|
| `conf.places=`   | `16`

