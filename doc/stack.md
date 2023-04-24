# Stack Notation

The documentation uses a notation to show how the stack is modified when
a calculator operation is evaluated. The following is an example of the
notation to indicate that an operation consumes two `BigInt` values from
the stack and pushes one `BigInt` value back to the stack:

    ( BigInt BigInt -- BigInt )

Stack notation is enclosed by parenthesis `()`, items consumed from the
stack are to the left of `--`, and items produced to the stack are on the
right of `--`. The top of the stack is the right-most element.

An operation that consumes no values from the stack is notated as:

    ( -- BigInt )

An operation that produces to values from the stack is notated as:

    ( BigInt -- )

An operation that does not modify the stack is shown by:

    ( -- )

Parameters can be named by placing a name and a `:` before the type:

    ( real:Float imag:Float -- Complex )


A suffix of `*` is used when the parameter consumes the remaining elements of
the stack. if any. For example, the stack notation for the `sum` operation is:

    ( Num* -- Num )

A `...` is a placeholder for all other items in the stack. For example, the
notation for `down` which takes the top element and places it on the bottom
of the stack is as follows:

    ( ... Val -- Val ... )

