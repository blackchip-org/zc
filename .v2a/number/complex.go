package number

type Complex interface {
	Add(Complex) Complex
	String() string
}

type ComplexImpl struct {
	Parse func(string) (Complex, error)
	New   func(float64, float64) Complex
}

var implComplex = ComplexImpl{
	Parse: func(string) (Complex, error) { return nil, ErrNotSupported },
	New:   func(float64, float64) Complex { return nil },
}

func UseComplex(impl ComplexImpl) {
	implComplex = impl
}

func ParseComplex(s string) (Complex, error) {
	return implComplex.Parse(s)
}

func MustParseComplex(s string) Complex {
	c, err := implComplex.Parse(s)
	if err != nil {
		panic(err)
	}
	return c
}

func NewComplex(r float64, i float64) Complex {
	return implComplex.New(r, i)
}
