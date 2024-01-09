## seq - examples

Even numbers:

<!-- test: even-numbers -->

| Input                 | Stack
|-----------------------|-------------
| `1 6 seq`             | `1 \| 2 \| 3 \| 4 \| 5 \| 6`
| `[2 mod 0 eq] filter` | `2 \| 4 \| 6`

Powers of two:

<!-- test: powers-of-two -->

| Input                 | Stack
|-----------------------|-------------
| `1 8 seq`             | `1 \| 2 \| 3 \| 4 \| 5 \| 6 \| 7 \| 8`
| `[2 swap pow] map`    | `2 \| 4 \| 8 \| 16 \| 32 \| 64 \| 128 \| 256`
