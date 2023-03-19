package scanner

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
	return func(s *Scanner) {
		if def.Sign(s.This) {
			s.Keep()
		}
		seenDecSep := false
		exponent := false
		for {
			if def.DecSep(s.This) {
				if seenDecSep {
					break
				}
				seenDecSep = true
				s.Keep()
			} else if def.Exponent(s.This) {
				exponent = true
				s.Keep()
				break
			} else if def.Digit(s.This) {
				s.Keep()
			} else {
				break
			}
		}
		if !exponent {
			return
		}
		if def.Sign(s.This) {
			s.Keep()
		}
		for def.Digit(s.This) {
			s.Keep()
		}
	}
}
