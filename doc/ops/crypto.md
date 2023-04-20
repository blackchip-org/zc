# crypto

Cryptographic ciphers

<!-- index -->

| Operation          | Description
|--------------------|-------------------------
| [rot-13](#rot-13)  | Rotate characters by 13


## rot-13

Rotate all characters in string `s` by 13 spaces.

    ( s:Str -- rotate:Str )

Example:

<!-- test: rot-13 -->

| Input               | Stack
|---------------------|-------------
| `'Behind the tree!` | `Behind the tree!`
| `rot-13`            | `Oruvaq gur gerr!`
| `rot-13`            | `Behind the tree!`
