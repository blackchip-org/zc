
'add' test
    2 4 add
    6 assert


### global reference

//a/g;

### functions

(fn export by default, use local keyword?)

fn parity /a
    'even' 'odd'
    /a mod 2 eq 0
    iif

fn max /a /b
    /a /b
    /a /b gt
    iif

fn fizz_buzz /n
    if /n 15 mod 0 eq
        'fizz buzz' .
    if /n 5 mod 0 eq
        'buzz' .
    if /n 3 mod 0 eq if
        'fizz' .
    /n

fn pally /str
    /str cp reverse ==

fn sum //a
    while n 1 gt
        add

fn reverse /str
    //a; /str '' split
    //a; empty not while
        //a; pop
    '' join

fn dist /x1 /y1 /x2 /y2
    /x2 /x1 sub
    2 pow
    /y2 /y1 sub
    2 pow
    add sqrt

fn shuffle //x
    //x
    loop n
        0 n rand
        pull







