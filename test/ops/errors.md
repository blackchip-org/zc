# errors

    1 add               !- add: not enough arguments, expected 2
    'a' 'b' add         !- no operation for: 'a' 'b' add
    1 2 3 [1 add] fold  !- function '1 add' is invalid: does not reduce

## macros

    def undo 'foo'  !- invalid name
    def 0 'foo'     !- invalid name
    def /foo        !- invalid name

