# errors

1 add
[add] not enough arguments, expected 2

'a' 'b' add
[add] no operation for String, String

1 2 3 [1 add] fold
[1 add] invalid function: does not reduce

## macros

def undo 'foo'
invalid name

def 0 'foo'
invalid name

def /foo
macro not defined: /foo