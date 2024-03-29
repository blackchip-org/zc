# math

## abs

    c -4 abs    -- 4
    c -4.4 abs  -- 4.4
    c -4.4f abs -- 4.4
    c -1/2 abs  -- 1/2
    c 3+4i abs  -- 5

## and

    0xff 0xf0 and hex -- 0xf0

## add

    c 8 2 add           -- 10
    c 1.1 2.2 add       -- 3.3
    c 1.1f 2.2f add     -- 3.3000000000000003
    c 1/2 1/4 add       -- 3/4
    c 3-3/4 2-1/2 add   -- 6 1/4
    c 6+2i 2+6i add     -- 8+8i

## ceil

    c 4 ceil    -- 4
    c 4.2 ceil  -- 5
    c 4.2f ceil -- 5

## div

    c 2 3 div       --  0.6666666666666666666
    c 2f 3f div     -- 0.6666666666666666
    c 2 0 div       !- division by zero: 2 0 div
    c 6+8i 2+2i div -- 3.5+0.5i
    c 1/4 3/4 div   -- 1/3
    c 6+8i 0+0i div !- division by zero: 6+8i 0+0i div

## eq

    c 2 2 eq        -- true
    c 2 3 eq        -- false
    c 2.2 2.2 eq    -- true
    c 2.2f 2.2f eq  -- true
    c 2.3 2.2 eq    -- false
    c 1/2 0.5 eq    -- true
    c 8+8i 8+8i eq  -- true

## floor

    c 4 floor       -- 4
    c 4.2 floor     -- 4
    c 4.2f floor    -- 4

## gt

    c 4 5 gt        -- false
    c 5 5 gt        -- false
    c 6 5 gt        -- true
    c 4.2 5.2 gt    -- false
    c 5.2 5.2 gt    -- false
    c 6.2 5.2 gt    -- true
    c 6.2f 5.2f gt  -- true
    c 1/2 0.25 gt   -- true

## gte

    c 4 5 gte       -- false
    c 5 5 gte       -- true
    c 6 5 gte       -- true
    c 4.2 5.2 gte   -- false
    c 5.2 5.2 gte   -- true
    c 6.2 5.2 gte   -- true
    c 6.2f 5.2f gte -- true
    c 1/2 0.25 gte  -- true

## hex

    43981 hex -- 0xabcd

## lt

    c 4 5 lt        -- true
    c 5 5 lt        -- false
    c 6 5 lt        -- false
    c 4.2 5.2 lt    -- true
    c 5.2 5.2 lt    -- false
    c 6.2 5.2 lt    -- false
    c 6.2f 5.2f lt  -- false
    c 1/2 0.25 lt   -- false

## lte

    c 4 5 lte       -- true
    c 5 5 lte       -- true
    c 6 5 lte       -- false
    c 4.2 5.2 lte   -- true
    c 5.2 5.2 lte   -- true
    c 6.2 5.2 lte   -- false
    c 6.2f 5.2f lte -- false
    c 1/2 0.25 lte  -- false

## mod

    c 7 2 mod           -- 1
    c 7 0 mod           !- division by zero: 7 0 mod
    c 5.75 0.5 mod      -- 0.25
    c 5.75f 0.5f mod    -- 0.25
    c 5.75 0 mod        !- division by zero: 5.75 0 mod
    c 5.75f 0f mod      !- division by zero: 5.75f 0f mod

## mul

    c 6 2 mul       -- 12
    c 6.6 2.2 mul   -- 14.52
    c 1e3 1e2 mul   -- 100000
    c 1/2 1/4 mul   -- 1/8
    c 6+8i 2+2i mul -- -4+28i

## neg

    c -1 neg    -- 1
    c -1.1 neg  -- 1.1
    c -1e30 neg -- 1e30
    c -1/2 neg  -- 1/2

## neq

    c 2 2 eq not       -- false
    c 2 3 eq not       -- true
    c 2.2 2.2 eq not   -- false
    c 2.2f 2.2f eq not -- false
    c 2.3 2.2 eq not   -- true
    c 8+8i 8+8i eq not -- false
    c 1/2 1/2 eq not   -- false

## sign

    c -6 sign       -- -1
    c -2.2 sign     -- -1
    c -2.2f sign    -- -1
    c -1e30 sign    -- -1
    c -3/4 sign     -- -1

## sub

    c 8 2 sub       -- 6
    c 1.1 2.2 sub   -- -1.1
    c 1.1f 2.2f sub -- -1.1
    c 3/4 1/2 sub   -- 1/4
    c 6+2i 2+6i sub -- 4-4i
