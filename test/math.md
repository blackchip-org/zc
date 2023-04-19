### abs

    -4 abs
    4

    -4.4 abs
    4.4

    -4.4f abs
    4.4

    -1/2 abs
    1/2

    3+4i abs
    5

### and

    0xff 0xf0 and hex
    0xf0

### add

    8 2 add
    10

    1.1 2.2 add
    3.3

    1.1f 2.2f add
    3.3

    1/2 1/4 add
    3/4

    3-3/4 2-1/2 add
    6 1/4

    6+2i 2+6i add
    8+8i

## ceil

    4 ceil
    4

    4.2 ceil
    5

    4.2f ceil
    5

## div

    2 3 div
    0.6666666666666667

    2f 3f div
    0.6666666666666666

    2 0 div
    division by zero

    6+8i 2+2i div
    3.5+0.5i

    1/4 3/4 div
    1/3

    6+8i 0+0i div
    division by zero

## eq

    2 2 eq
    true

    2 3 eq
    false

    2.2 2.2 eq
    true

    2.2f 2.2f eq
    true

    2.3 2.2 eq
    false

    1/2 0.5 eq
    true

    8+8i 8+8i eq
    true

## floor

    4 floor
    4

    4.2 floor
    4

    4.2f floor
    4

## gt

    4 5 gt
    false

    5 5 gt
    false

    6 5 gt
    true

    4.2 5.2 gt
    false

    5.2 5.2 gt
    false

    6.2 5.2 gt
    true

    6.2f 5.2f gt
    true

    1/2 0.25 gt
    true

## gte

    4 5 gte
    false

    5 5 gte
    true

    6 5 gte
    true

    4.2 5.2 gte
    false

    5.2 5.2 gte
    true

    6.2 5.2 gte
    true

    6.2f 5.2f gte
    true

    1/2 0.25 gte
    true

## hex

    43981 hex
    0xabcd

## lt

    4 5 lt
    true

    5 5 lt
    false

    6 5 lt
    false

    4.2 5.2 lt
    true

    5.2 5.2 lt
    false

    6.2 5.2 lt
    false

    6.2f 5.2f lt
    false

    1/2 0.25 lt
    false

## lte

    4 5 lte
    true

    5 5 lte
    true

    6 5 lte
    false

    4.2 5.2 lte
    true

    5.2 5.2 lte
    true

    6.2 5.2 lte
    false

    6.2f 5.2f lte
    false

    1/2 0.25 lte
    false

## mod

    7 2 mod
    1

    7 0 mod
    division by zero

    5.75 0.5 mod
    0.25

    5.75f 0.5f mod
    0.25

    5.75 0 mod
    division by zero

    5.75f 0f mod
    division by zero

## mul

    6 2 mul
    12

    6.6 2.2 mul
    14.52

    1e3 1e2 mul
    100000

    1/2 1/4 mul
    1/8

    6+8i 2+2i mul
    -4+28i

## neg

    -1 neg
    1

    -1.1 neg
    1.1

    -1e30 neg
    1e+30

    -1/2 neg
    1/2

## neq

    2 2 neq
    false

    2 3 neq
    true

    2.2 2.2 neq
    false

    2.2f 2.2f neq
    false

    2.3 2.2 neq
    true

    8+8i 8+8i neq
    false

    1/2 1/2 neq
    false


## sign

    -6 sign
    -1

    -2.2 sign
    -1

    -2.2f sign
    -1

    -1e30 sign
    -1

    -3/4 sign
    -1

## sub

    8 2 sub
    6

    1.1 2.2 sub
    -1.1

    1.1f 2.2f sub
    -1.1

    3/4 1/2 sub
    1/4

    6+2i 2+6i sub
    4-4i
