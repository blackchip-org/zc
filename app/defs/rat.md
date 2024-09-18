A `Rational`value is a number that has a numerator *n* and a denominator *d* in
the form of *n*`/`*d*. A whole number can prefix a rational using a ` `, `_` or
`-` character. Examples:

<!-- test: TypesRational -->

| Input                   | Stack
|-------------------------|-------------
| `c 1/2 1/4 add`         | `3/4`
| `c 2_1/2 3_1/4 add`     | `5 3/4`
| `c [2 1/2] [3 1/4] add` | `5 3/4`
| `c 2-1/2 3-1/4 add`     | `5 3/4`
