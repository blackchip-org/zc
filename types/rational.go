package types

type Rational interface {
	Value
	Add(Rational) Rational
}

type RationalImpl struct {
	Parse func(string) (Rational, error)
	New   func(int, int) Rational
}

var implRational = RationalImpl{
	Parse: func(string) (Rational, error) { return nil, ErrNotSupported },
	New:   func(int, int) Rational { return nil },
}

func UseRational(impl RationalImpl) {
	implRational = impl
}

func ParseRational(s string) (Rational, error) {
	return implRational.Parse(s)
}

func MustParseRational(s string) Rational {
	r, err := implRational.Parse(s)
	if err != nil {
		panic(err)
	}
	return r
}

func NewRational(a int, b int) Rational {
	return implRational.New(a, b)
}

type RationalType struct{}

func (t RationalType) String() string { return "Rational" }

func (t RationalType) Parse(s string) (Rational, error) {
	return implRational.Parse(s)
}

func (t RationalType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t RationalType) Format(r Rational) string {
	return r.String()
}

func (t RationalType) Value(r Rational) Value {
	return r
}

func (t RationalType) Native(v Value) Rational {
	return v.Native().(Rational)
}
