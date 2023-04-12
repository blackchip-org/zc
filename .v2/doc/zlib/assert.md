# assert

<!-- eval: import assert -->

Type assertions

    import assert -- prelude: dev

<!-- index -->

| Operation            | Description
|---------------------|----------------
| [bigint](#bigint)   | Assert value is a big integer
| [bool](#bool)       | Assert value is a boolean
| [decimal](#decimal) | Assert value is a fixed-point number
| [float](#float)     | Assert value is a floating-point number
| [int](#int)         | Assert value is an integer
| [int32](#int32)     | Assert value is a 32-bit signed integer
| [int64](#int64)     | Assert value is a 64-bit signed integer


## bigint

Assert that `a` is a big integer. If not, an error is raised.

    ( a:BigInt -- )

Example:

<!-- test: bigint -->

| Input               | Stack
|---------------------|------------------
| `120`               | `120`
| `try assert.bigint` | `true`
| `clear`             |
| `'foo`              | `foo`
| `try assert.bigint` | `expecting BigInt but got 'foo' \| false`


## bool

Assert that `a` is a boolean. If not, an error is raised.

    ( a:Bool -- )

Example:

<!-- test: bool -->

| Input               | Stack
|---------------------|------------------
| `false`             | `false`
| `try assert.bool`   | `true`
| `clear`             |
| `'no`               | `no`
| `try assert.bool`   | `expecting Bool but got 'no' \| false`


## decimal

Assert that `a` is a fixed-point number. If not, an error is raised.

    ( a:Decimal -- )

Example:

<!-- test: decimal -->

| Input                | Stack
|----------------------|------------------
| `123.45`             | `123.45`
| `try assert.decimal` | `true`
| `clear`              |
| `12.34.56`           | `12.34.56`
| `try assert.decimal` | `expecting Decimal but got 12.34.56 \| false`


## float

Assert that `a` is a floating-point number. If not, an error is raised.

    ( a:Float -- )

Example:

<!-- test: float -->

| Input               | Stack
|---------------------|------------------
| `123.45`            | `123.45`
| `try assert.float`  | `true`
| `clear`             |
| `12.34.56`          | `12.34.56`
| `try assert.float`  | `expecting Float but got 12.34.56 \| false`


## int

Assert that `a` is an integer. If not, an error is raised.

    ( a:Int -- )

Example:

<!-- test: int -->

| Input               | Stack
|---------------------|------------------
| `123`               | `123`
| `try assert.int`    | `true`
| `clear`             |
| `123.45`            | `123.45`
| `try assert.int`    | `expecting Int but got 123.45 \| false`


## int32

Assert that `a` is an 32-bit signed integer. If not, an error is raised.

    ( a:Int32 -- )

Example:

<!-- test: int -->

| Input                  | Stack
|------------------------|------------------
| `2 32 pow 2 div 1 sub` | `2147483647`
| `try assert.int32`     | `true`
| `clear`                |
| `2 32 pow 2 div`       | `2147483648`
| `try assert.int32`     | `expecting Int32 but got 2147483648 \| false`


## int64

Assert that `a` is an 64-bit signed integer. If not, an error is raised.

    ( a:Int64 -- )

Example:

<!-- test: int -->

| Input                  | Stack
|------------------------|------------------
| `2 64 pow 2 div 1 sub` | `9223372036854775807`
| `try assert.int64`     | `true`
| `clear`                |
| `2 64 pow 2 div`       | `9223372036854775808`
| `try assert.int64`     | `expecting Int64 but got 9223372036854775808 \| false`
