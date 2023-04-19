# sci

Scientific operations.

<!-- index -->

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [abs](#abs)             |          | Absolute value
| [ceil](#ceil)           |          | Ceiling
| [floor](#floor)         |          | Floor


## abs

If `a` is less than zero, the negated value of `a`, otherwise `a`.

    ( a:BigInt   -- abs:BigInt );   or
    ( a:Decimal  -- abs:Decimal );  or
    ( a:Float    -- abs:Float );    or
    ( a:Rational -- abs:Rational );

Example:

<!-- test: abs -->

| Input   | Stack
|---------|-------------|
| `-6`    | `-6`
| `abs`   | `6`

The distance of `a` from zero in the complex plane.

    ( a:Complex -- abs:Complex )

Example:

<!-- test: abs-complex -->

| Input   | Stack
|---------|-------------|
| `3+4i`  | `3+4i`
| `abs`   | `5`

## ceil

The nearest integer value greater than or equal to `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or

Example:

<!-- test: ceil -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `ceil`  | `7`

## floor

The nearest integer value less than or equal to `a`.

    ( a:BigInt   b:BigInt   -- add:BigInt );   or
    ( a:Decimal  b:Decimal  -- add:Decimal );  or
    ( a:Float    b:Float    -- add:Float );    or

Example:

<!-- test: floor -->

| Input   | Stack
|---------|-------------|
| `6.12`  | `6.12`
| `floor` | `6`
