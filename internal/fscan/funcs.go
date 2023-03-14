package fscan

type Func func(*Scanner)

func WhileFunc(while RuneClass) Func {
	return func(s *Scanner) {
		for while(s.This) {
			s.Keep()
		}
	}
}

func UntilFunc(until RuneClass) Func {
	return func(s *Scanner) {
		for !until(s.This) && s.This != End {
			s.Keep()
		}
	}
}

func NumberFunc(def NumberDef) Func {
	return SepNumberFunc(def, NoSepDef)
}

func SepNumberFunc(def NumberDef, sep NumberSepDef) Func {
	return func(s *Scanner) {
		if def.Sign.Is(s.This) {
			def.Sign.Action(s)
		}
		seenDecSep := false
		exponent := false
		for {
			if sep.Left.Is(s.This) && !seenDecSep && def.Digit.Is(s.Behind) && def.Digit.Is(s.Ahead) {
				sep.Left.Action(s)
			} else if sep.Right.Is(s.This) && seenDecSep && def.Digit.Is(s.Behind) && def.Digit.Is(s.Ahead) {
				sep.Right.Action(s)
			} else if def.DecSep.Is(s.This) {
				if seenDecSep {
					break
				}
				seenDecSep = true
				def.DecSep.Action(s)
			} else if def.Exponent.Is(s.This) {
				exponent = true
				def.Exponent.Action(s)
				break
			} else if def.Digit.Is(s.This) {
				def.Digit.Action(s)
			} else {
				break
			}
		}
		if !exponent {
			return
		}
		if def.Sign.Is(s.This) {
			def.Sign.Action(s)
		}
		for def.Digit.Is(s.This) {
			def.Digit.Action(s)
		}
	}
}
