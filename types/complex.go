package types

type Complex interface {
	Value
	Add(Complex) Complex
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

type ComplexType struct{}

func (t ComplexType) String() string { return "Complex" }

func (t ComplexType) Parse(s string) (Complex, error) {
	return implComplex.Parse(s)
}

func (t ComplexType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t ComplexType) Format(c Complex) string {
	return c.String()
}

func (t ComplexType) Value(c Complex) Value {
	return c
}

func (t ComplexType) Native(v Value) Complex {
	return v.Native().(Complex)
}
