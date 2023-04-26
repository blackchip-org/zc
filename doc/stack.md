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

An operation that produces no values to the stack is notated as:

    ( BigInt -- )

An operation that does not modify the stack is shown by:

    ( -- )

Parameters can be named by placing a name and a `:` before the type:

    ( real:Float imag:Float -- Complex )

If a parameter doesn't have a meaningful name, it may use names such
as `p0`, `p1`, etc. depending on its location on the stack:

    ( p0:BigInt p1:BigInt -- BigInt )

If an operation returns more than one value, the values may be also be named:

    ( x1:Float y1:Float -- x2:Float y2:Float )

When a there is a name that is the same on both sides of the `--`, that
indicates that the operation does not change the value. In the following
example, the *y* value is changed but the *x* value is not:

    ( x:Float y1:Float -- x:Float y2:Float )

A suffix of `*` is used when the parameter consumes the remaining elements of
the stack. if any. For example, the stack notation for the `sum` operation is:

    ( T* -- T )

A `...` is a placeholder for all other items in the stack.
