# math

Basic mathematical operations.

- Prelude: cli, dev
- Use: include

| Operation               | Alias    | Description
|-------------------------|----------|------------
| [add](#add)             | `a`, `+` | Addition
| [div](#div)             | `d`, `/` | Division
| [mul](#mul)             | `m`, `*` | Multiplication
| [sub](#sub)             | `s`, `-` | Subtraction


## add

Adds two numbers, `a + b`, and puts the result on the stack.

    ( a:Num b:Num -- Num )

Aliases: `a`, `+`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `a`     | `8`

## div

Divides two numbers, `a / b` and puts the result on the stack.

    ( a:Num b:Num -- Num )

Aliases: `d`, `/`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `d`     | `3`

## mul

Multiplies two numbers, `a * b`, and puts the result on the stack.

    ( a:Num b:Num -- Num )

Aliases: `m`, `*`

Example:

| Input   | Stack
|---------|-------------|
| `6`     | `6`
| `2`     | `6 \| 2`
| `m`     | `12`

## sub

Subtracts two numbers, `a - b`, and puts the result on the stack.

    ( a:Num b:Num -- Num )

Aliases: `s`, `-`

| Input         | Stack
|---------------|-------------|
| `6`           | `6`
| `2`           | `6 \| 2`
| `s`           | `4`