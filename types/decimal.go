package types

type Decimal interface {
	Value
	Add(Decimal) Decimal
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

type DecimalType struct{}

func (t DecimalType) String() string { return "Decimal" }

func (t DecimalType) Parse(s string) (Decimal, error) {
	return implDecimal.Parse(cleanNumber(s))
}

func (t DecimalType) ParseValue(s string) (Value, error) {
	v, err := t.Parse(s)
	if err != nil {
		return nil, err
	}
	return t.Value(v), nil
}

func (t DecimalType) Format(d Decimal) string {
	return d.String()
}

func (t DecimalType) Value(d Decimal) Value {
	return d
}

func (t DecimalType) Native(v Value) Decimal {
	return v.Native().(Decimal)
}
