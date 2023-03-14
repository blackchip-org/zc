package fscan

type RuneClass func(rune) bool

func Rune(r1 rune) RuneClass {
	return func(r0 rune) bool {
		return r0 == r1
	}
}

func Rune2(r1 rune, r2 rune) RuneClass {
	return func(r0 rune) bool {
		return r0 == r1 || r0 == r2
	}
}

func RuneRange(from rune, to rune) RuneClass {
	return func(r0 rune) bool {
		return r0 >= from && r0 <= to
	}
}

func Or(classes ...RuneClass) RuneClass {
	return func(r0 rune) bool {
		for _, class := range classes {
			if class(r0) {
				return true
			}
		}
		return false
	}
}

var (
	Always     = func(r rune) bool { return r != End }
	Comma      = Rune(',')
	Digit01    = Rune2('0', '1')
	Digit07    = RuneRange('0', '7')
	Digit09    = RuneRange('0', '9')
	Digit0F    = Or(Digit09, LCharAF, UCharAF)
	ExponentE  = Rune2('E', 'e')
	LCharAF    = RuneRange('a', 'f')
	Never      = func(r rune) bool { return false }
	Period     = Rune('.')
	PlusMinus  = Rune2('+', '-')
	Space      = Rune(' ')
	UCharAF    = RuneRange('A', 'F')
	Underscore = Rune('_')
)
