# bool

Boolean operations.

<!-- index -->

| Operation         | Description
|-------------------|---------------
| [and](#and)       | Logical conjunction
| [eq](#eq)         | Equals
| [false](#false)   | False
| [gt](#gt)         | Greater than
| [gte](#gte)       | Greater than or equals
| [true](#true)     | True
| [lt](#lt)         | Less than
| [lte](#lte)       | Less than or equals
| [neq](#neq)       | Not Equals
| [not](#not)       | Negation
| [or](#or)         | Logical disjunction

## and

The logical conjunction of *p0* and *p1*.

    ( p0:Bool p1:Bool -- Bool )

Example:

<!-- test: and -->

| Input               | Stack
|---------------------|-------------|
| `c true true   and` | `true`
| `c true false  and` | `false`
| `c false false and` | `false`

## eq

`true` if *p0* and *p1* are equal, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`
- `Str`

Example:

<!-- test: eq -->

| Input                      | Stack
|----------------------------|-------------|
| `c 1234.56 1,234.56   eq`  | `true`
| `c 1234.56 1234.56000 eq`  | `true`
| `c 1234.56 $1,234.56  eq`  | `true`
| `c 1234.56 +1,234.56  eq`  | `true`

## false

Places `false` on the stack.

    ( -- Str )

Example:

<!-- test: false -->

| Input    | Stack
|----------|-------------|
| `false`  | `false`

## gt

`true` if *p0* is greater than *p1*, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Rational`
- `Float`
- `Str`

Example:

<!-- test: gt -->

| Input        | Stack
|--------------|-------------|
| `c 1  0 gt`  | `true`
| `c 0  0 gt`  | `false`
| `c -1 0 gt`  | `false`

## gte

`true` if *p0* is greater than or equal to *p1*, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Rational`
- `Float`
- `Str`

Example:

<!-- test: false -->

| Input        | Stack
|--------------|-------------|
| `c 1 0  gte` | `true`
| `c 0 0  gte` | `true`
| `c -1 0 gte` | `false`

## true

Places `true` on the stack.

    ( -- Str )

Example:

<!-- test: true -->

| Input    | Stack
|----------|-------------|
| `true`   | `true`

## lt

`true` if *p0* is less than *p1*, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Rational`
- `Float`
- `Str`

Example:

<!-- test: lt -->

| Input        | Stack
|--------------|-------------|
| `c 1 0  lt`  | `false`
| `c 0 0  lt`  | `false`
| `c -1 0 lt`  | `true`

## lte

`true` if *p0* is less than or equal to *p1*, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Rational`
- `Float`
- `Str`

Example:

<!-- test: lte -->

| Input        | Stack
|--------------|-------------|
| `c 1 0  lte` | `false`
| `c 0 0  lte` | `true`
| `c -1 0 lte` | `true`

## neq

`true` if *p0* and *p1* are not equal to each other, otherwise `false`.

    ( p0:T p1:T -- Bool )

Where *T* is one of:
- `BigInt`
- `Decimal`
- `Float`
- `Rational`
- `Complex`

Example:

<!-- test: neq -->

| Input                  | Stack
|------------------------|-------------|
| `123 123 neq`          | `false`
| `clear`                |
| `123 456 neq`          | `true`

## not

`true` if *p0* is `false`, otherwise `false`

    ( p0:Bool -- Bool )

Example:

<!-- test: not -->

| Input                  | Stack
|------------------------|-------------|
| `true not`             | `false`
| `not`                  | `true`

## or

The logical disjunction of *p0* and *p1*.

    ( p0:Bool p1:Bool -- Bool )

Example:

<!-- test: or -->

| Input              | Stack
|--------------------|-------------|
| `c true true   or` | `true`
| `c true false  or` | `true`
| `c false false or` | `false`
