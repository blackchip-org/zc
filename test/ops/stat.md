# stat

## average

    c /foobar avg       !- no operation for: 'foobar' 1 div

## factorial

    c -42 fact          !- invalid arguments, cannot be negative: -42 factorial
    c 20 fact           -- 2432902008176640000
    c 100 fact sn       -- 9.332621544394415e157
    c 10,000 fact sn    -- 2.846259681e35659

## variance

    c var-p             --
    c /foobar var-p     !- no operation for: 'foobar' 1 div
