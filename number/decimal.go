package number

type Decimal interface {
	Add(Decimal) Decimal
	String() string
}

type DecimalImpl struct {
	Parse func(string) (Decimal, error)
	New   func(float64) Decimal
}

var implDecimal = DecimalImpl{
	Parse: func(string) (Decimal, error) { return nil, ErrNotSupported },
	New:   func(float64) Decimal { return nil },
}

func UseDecimal(impl DecimalImpl) {
	implDecimal = impl
}

func ParseDecimal(s string) (Decimal, error) {
	return implDecimal.Parse(s)
}

func MustParseDecimal(s string) Decimal {
	d, err := implDecimal.Parse(s)
	if err != nil {
		panic(err)
	}
	return d
}

func NewDecimal(f float64) Decimal {
	return implDecimal.New(f)
}
