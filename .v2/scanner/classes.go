package scanner

import "unicode"

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
	Always        = func(r rune) bool { return r != EndCh }
	IsCharAF      = Or(IsUpperCharAF, IsLowerCharAF)
	IsCharAZ      = Or(IsUpperCharAZ, IsLowerCharAZ)
	IsCurrency    = func(r rune) bool { return unicode.Is(unicode.Sc, r) }
	IsDigit01     = Rune2('0', '1')
	IsDigit07     = RuneRange('0', '7')
	IsDigit09     = RuneRange('0', '9')
	IsDigit0F     = Or(IsDigit09, IsLowerCharAF, IsUpperCharAF)
	IsEnd         = Rune(EndCh)
	IsLowerCharAF = RuneRange('a', 'f')
	IsLowerCharAZ = RuneRange('a', 'z')
	Never         = func(r rune) bool { return false }
	IsUpperCharAF = RuneRange('A', 'F')
	IsUpperCharAZ = RuneRange('A', 'Z')
)
