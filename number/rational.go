package number

type Rational interface {
	Add(Rational) Rational
	String() string
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
