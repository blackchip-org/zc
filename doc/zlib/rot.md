# rot

<!-- eval: use rot -->

Rotation ciphers

Great for decoding geocaching hints.

<!-- index -->

| Operation               | Alias     | Description
|-------------------------|-----------|-----------
| [rotate-13](#rotat-13)  | `rot-13`  | Rotate characters by 13


## rotate-13

Rotate all characters in string `s` by 13 spaces.

    ( s:Str -- rotate:Str )

Example:

<!-- test: rotate-13 -->

| Input               | Stack
|---------------------|-------------
| `'Behind the tree!` | `Behind the tree!`
| `rot-13`            | `Oruvaq gur gerr!`
| `rot-13`            | `Behind the tree!`
