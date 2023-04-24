# rand

Random number generation

<!-- index -->

| Operation                   | Description
|-----------------------------|-----------------------
| [rand](#rand)               | Random float between 0 and 1
| [rand-choice](#rand-choice) | Randomly select an item on the stack
| [rand-int](#rand-int)       | Random integer between 1 and n
| [rand-seed](#rand-seed)     | Random number seed, set
| [rand-seed=](#rand-seed=)   | Random number seed, get
| [roll](#roll)               | Dice roller
| [shuffle](#shuffle)         | Shuffle the stack


## rand

Random float between 0 and 1

    ( -- rand:Float )

Example:

<!-- test: rand -->

| Input         | Stack
|---------------|-------------|
| `0 rand-seed` | *seed set to 0*
| `rand`        | `0.9451961492941164`

## rand-choice

Randomly select an item on the stack

    ( ...:any ?:any ...:any -- ?:any )

Example:

<!-- test: rand-choice -->

| Input          | Stack
|----------------|-------------|
| `2 rand-seed`  | *seed set to 2*
| `1 2 3 4 5 6`  | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `rand-choice`  | `5`

## rand-int

Random integer between 1 and n

    ( n:Int -- rand:Int )

Example:

<!-- test: rand-int -->

| Input         | Stack
|---------------|-------------|
| `0 rand-seed` | *seed set to 0*
| `10 rand-int` | `5`

## rand-seed

Sets the random number seed

    ( n:Int64 -- )

<!-- test: rand-seed -->

| Input         | Stack
|---------------|-------------|
| `1 rand-seed` | *seed set to 1*
| `10 rand-int` | `2`

## seed=

Gets the random number seed

    ( -- n:Int64 )

<!-- test: rand-seed= -->

| Input         | Stack
|---------------|-------------|
| `3 rand-seed` | *seed set to 3*
| `rand-seed=`  | `3`

## roll

Rolls dice as specified by `dice` in standard dice notation. The argument
`dice` may start with the number of dice to roll, followed by the literal
character `d`, and then the number of faces found on each die. For example,
use `3d6` to roll three six sided dice.

    ( dice:Str -- rolls...:Int )


<!-- test: roll -->

| Input           | Stack
|-----------------|-------------|
| `99 rand-seed`  | *seed set to 99*
| `3d6 roll`      | `6 \| 2 \| 1`
| `sum`           | `9`

## shuffle

Shuffle the stack

    ( ...:Any -- ...:Any )

<!-- test: shuffle -->

| Input          | Stack
|----------------|-------------|
| `0 rand-seed`  | *seed set to 0*
| `1 2 3 4 5 6`  | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `shuffle`      | `5 \| 4 \| 1 \| 3 \| 2 \| 6`



