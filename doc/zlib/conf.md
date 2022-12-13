# conf

Calculator configurations and settings

- Prelude: user, dev
- Use: import

| Operation                   | Description
|-----------------------------|----------------
| [len](#len)                 | Length
| [starts-with](#starts-with) | Starts with


## len

Length of string `a`.

    ( a:Str -- len:Int )

Example:

| Input        | Stack
|--------------|------------------|
| `'abcd`      | `abcd`
| `len`        | `4`