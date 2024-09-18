A floating point value is a number that can either be a:

- `Float`: a double precision floating-point number
- `Float/v`: a variable precision floating-point number
- `Float/sp`: a single precision floating-point number

Standard unqualified names (such as `add`) use fixed-point arithmetic when
available. Use the qualified names (such as `add/f`) to use float-point
arithmetic instead. 

The precision for the variable precision floating-point operations is set
to 113 bits in the mantissa by defualt. Use `prec/v` to change this to a
different value if desired. This type is implemented using the math/big 
pacakge in the standard go library. 



