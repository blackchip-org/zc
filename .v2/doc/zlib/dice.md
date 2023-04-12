# dice

<!-- eval: use dice -->

Dice roller

<!-- index -->

| Operation               | Description
|-------------------------|-----------------------
| [roll](#roll)           | Dice roller


## roll

Rolls dice as specified by `dice` in standard dice notation. The argument
`dice` may start with the number of dice to roll, followed by the literal
character `d`, and then the number of faces found on each die. For example,
use `3d6` to roll three six sided dice.

    ( dice:Str -- rolls...:Int )


<!-- test: roll -->

| Input           | Stack
|-----------------|-------------|
| `import rand`   | *imported rand*
| `99 rand.seed`  | *seed set to 99*
| `3d6 roll`      | `6 \| 2 \| 1`
| `sum`           | `9`


