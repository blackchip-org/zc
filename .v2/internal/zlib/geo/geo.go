package geo

type dms struct {
	deg  string
	min  string
	sec  string
	hemi rune
}

/*
func parseDMS(calc *zc.Calc, v string) (dms, bool) {
	var c dms
	_, err := calc.ParseFloat(v)
	if err == nil {
		c.deg = v
		return c, true
	}

	def := fscan.UDecRef
	def.DecSep = fscan.NewRule(fscan.Period, fscan.Keep)
	num := fscan.NumberFunc(def)
	s := fscan.NewForString(strings.TrimSpace(v))

	s.ScanWhitespace()
	c.deg = s.NextToken(num)
	if c.deg == "" {
		return c, false
	}
	switch s.This {
	case 'd', 'D', '°':
		s.Next()
	}
	s.ScanWhitespace()
	if fscan.Digit09(s.This) {
		c.min = s.NextToken(num)
		if c.min == "" {
			return c, false
		}
		if s.This != '\'' {
			return c, false
		}
		s.Next()
		s.ScanWhitespace()
		if fscan.Digit09(s.This) {
			c.sec = s.ScanUDec()
			if c.sec == "" {
				return c, false
			}
			if s.This != '"' {
				return c, false
			}
			s.Next()
		}
	}
	s.ScanWhitespace()
	switch s.This {
	case 'N', 'n', 'S', 's', 'W', 'w', 'E', 'e':
		c.hemi = s.This
		s.Next()
	}
	return c, s.IsEnd()
}

func parseDD(calc *zc.Calc, v string) (float64, bool) {
	c, ok := parseDMS(calc, v)
	if !ok {
		return 0, false
	}

	var deg, min, sec float64
	var err error

	deg, err = calc.ParseFloat(c.deg)
	if err != nil {
		return 0, false
	}
	if c.min != "" {
		min, err = calc.ParseFloat(c.min)
		if err != nil {
			return 0, false
		}
		if deg != float64(int(deg)) {
			return 0, false
		}
	}
	if c.sec != "" {
		sec, err = calc.ParseFloat(c.sec)
		if err != nil {
			return 0, false
		}
		if min != float64(int(min)) {
			return 0, false
		}
	}

	var sign float64
	switch c.hemi {
	case 'N', 'n', 'E', 'e':
		sign = 1
	case 'S', 's', 'W', 'w':
		sign = -1
	default:
		sign = 1
	}

	z := sign * (deg + (min / 60.0) + (sec / 3600.0))
	return z, true
}

func popDMS(env *zc.Env) (dms, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return dms{}, err
	}
	c, ok := parseDMS(env.Calc, s)
	if !ok {
		return dms{}, fmt.Errorf("expecting Coordinate but got %v", zc.Quote(s))
	}
	return c, nil
}

func popDD(env *zc.Env) (float64, error) {
	s, err := env.Stack.Pop()
	if err != nil {
		return 0, err
	}
	c, ok := parseDD(env.Calc, s)
	if !ok {
		return 0, fmt.Errorf("expecting Coordinate but got %v", zc.Quote(s))
	}
	return c, nil
}

func formatDMS(c dms) string {
	if c.min == "" && c.sec == "" {
		return fmt.Sprintf("%v°%c", c.deg, c.hemi)
	}
	if c.sec == "" {
		return fmt.Sprintf("%v°%v'%c", c.deg, c.min, c.hemi)
	}
	return fmt.Sprintf("%v°%v'%v\"%c", c.deg, c.min, c.sec, c.hemi)
}

func DecimalDegrees(env *zc.Env) error {
	c, err := popDD(env)
	if err != nil {
		return err
	}
	env.Stack.PushFloat(c)
	return nil
}

func DegreesMinutes(env *zc.Env) error {
	c, err := popDD(env)
	if err != nil {
		return err
	}
	deg := int(c)
	min := math.Mod(c*60, 60)
	env.Stack.Push(fmt.Sprintf("%v°%v'", deg, min))
	return nil
}

func DegreesMinutesSeconds(env *zc.Env) error {
	// c, err := popDD(env)
	// if err != nil {
	// 	return err
	// }
	// deg := int(c)
	// min := math.Mod(c*60, 60)
	// sec := math.Mod(c*3600, 60)
	// env.Stack.Push(fmt.Sprintf("%v°%02f'", deg, min))
	return nil
}

func RoundCoordinate(env *zc.Env) error {
	n, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	dms, err := popDMS(env)
	if err != nil {
		return err
	}
	ns := strconv.Itoa(n)
	if dms.sec != "" {
		f := "%02." + ns + "f"
		dms.sec = fmt.Sprintf(f, env.Calc.MustParseFloat(dms.sec))
	} else if dms.min != "" {
		f := "%02." + ns + "f"
		dms.min = fmt.Sprintf(f, env.Calc.MustParseFloat(dms.min))
	} else {
		f := "%." + ns + "f"
		dms.deg = fmt.Sprintf(f, env.Calc.MustParseFloat(dms.deg))
	}
	env.Stack.Push(formatDMS(dms))
	return nil
}
*/
