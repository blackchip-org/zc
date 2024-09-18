A `Bool` is a value that is either true or false.

An item on the stack can be parsed as a boolean if it is equal to `true`
or `false` when all characters are converted to lowercase. Operations
are defined for `true` and `false` that simply return that string.

Example:

<!-- test: TypesBool -->

| Input           | Stack
|-----------------|-------------
| `true true and` | `true`
| `'FALSE' and`   | `false`
