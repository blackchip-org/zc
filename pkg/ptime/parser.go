package ptime

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/ptime/locale"
)

type Parsed struct {
	Weekday     string `json:",omitempty"`
	Year        string `json:",omitempty"`
	Month       string `json:",omitempty"`
	Day         string `json:",omitempty"`
	Hour        string `json:",omitempty"`
	Minute      string `json:",omitempty"`
	Second      string `json:",omitempty"`
	FracSecond  string `json:",omitempty"`
	Period      string `json:",omitempty"`
	Zone        string `json:",omitempty"`
	Offset      string `json:",omitempty"`
	DateSep     string `json:",omitempty"`
	TimeSep     string `json:",omitempty"`
	DateTimeSep string `json:",omitempty"`
	HourSep     string `json:",omitempty"`
}

func (p Parsed) String() string {
	text, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(text)
}

type state int

const (
	unknown state = iota
	parsingDate
	parsingTime
	parsingZone
	done
)

func (s state) String() string {
	switch s {
	case parsingDate:
		return "ParsingDate"
	case parsingTime:
		return "ParsingTime"
	case parsingZone:
		return "ParsingZone"
	case done:
		return "Done"
	}
	return "Unknown"
}

type dateOrder int

const (
	unknownOrder dateOrder = iota
	dayMonthYearOrder
	monthDayYearOrder
	yearMonthDayOrder
	yearDayOrder
)

func (d dateOrder) String() string {
	switch d {
	case dayMonthYearOrder:
		return "day-month-year"
	case monthDayYearOrder:
		return "month-day-year"
	case yearMonthDayOrder:
		return "year-month-day"
	case yearDayOrder:
		return "year-day"
	}
	return "unknown"
}

type Parser struct {
	loc       *locale.Locale
	tokens    []Token
	tok       Token
	idx       int
	parsed    Parsed
	Trace     bool
	state     state
	dateOrder dateOrder
	parseOne  bool
}

func NewParser(l *locale.Locale) *Parser {
	return &Parser{loc: l}
}

func (p *Parser) parse(text string) (Parsed, error) {
	p.trace("state: %v", p.state)
	p.tokens = Scan(text)

	if len(p.tokens) == 0 {
		return Parsed{}, nil
	}

	p.idx = -1
	p.tok = p.tokens[0]
	p.parsed = Parsed{}
	p.dateOrder = unknownOrder

	for p.tok.Type != End {
		var err error
		p.trace("top")
		p.next()
		switch p.tok.Type {
		case Text:
			err = p.parseText()
		case Number:
			err = p.parseNumber()
		case Indicator:
			err = p.parseIndicator()
		}
		if err != nil {
			return p.parsed, err
		}
	}
	return p.parsed, nil
}

func (p *Parser) Parse(text string) (Parsed, error) {
	p.state = unknown
	p.parseOne = false
	return p.parse(text)
}

func (p *Parser) ParseDate(text string) (Parsed, error) {
	p.state = parsingDate
	p.parseOne = true
	return p.parse(text)
}

func (p *Parser) ParseTime(text string) (Parsed, error) {
	p.state = parsingTime
	p.parseOne = true
	return p.parse(text)
}

func (p *Parser) parseText() error {
	if p.state == unknown {
		p.state = parsingDate
	}
	if p.state == parsingDate {
		if p.parsed.Weekday == "" {
			if day, ok := lookupDay(p.loc, p.tok.Val); ok {
				p.trace("is weekday")
				p.parsed.Weekday = day
				return nil
			}
		}
		if p.parsed.Month == "" {
			if mon, ok := lookupMonth(p.loc, p.tok.Val); ok {
				p.trace("is month")
				p.parsed.Month = mon
				return nil
			}
		}
		if inSet(p.tok.Val, p.loc.DateTimeSep) {
			p.trace("is date time separator")
			p.changeState(parsingTime)
			p.parsed.DateTimeSep = p.tok.Val
			return nil
		}
	}
	if p.state == parsingTime {
		if inSet(p.tok.Val, p.loc.TimeSep) {
			return nil
		}
		if p.parsed.Period == "" {
			period, ok := lookupPeriod(p.loc, p.tok.Val)
			if ok {
				p.trace("is period")
				p.parsed.Period = string(period)
				p.changeState(parsingZone)
				return nil
			}
		}
		if _, ok := p.loc.ZoneNamesShort[p.tok.Val]; ok {
			p.changeState(parsingZone)
		}
		if inSet(p.tok.Val, p.loc.UTCFlags) {
			p.trace("is UTC")
			p.parsed.Zone = p.tok.Val
			p.parsed.Offset = "+0000"
			return nil
		}
	}
	if p.state == parsingZone {
		if p.tok.Val == "UTC" && p.lookahead(1).Val == "-" {
			p.trace("is UTC offset")
			p.next()
			sign := p.tok.Val
			if sign != "+" && sign != "-" {
				p.trace("invalid sign in UTC offset: %v", p.tok.Val)
				return nil
			}
			p.next()
			offset, err := strconv.ParseInt(p.tok.Val, 0, 64)
			if err != nil {
				p.trace("invalid offset: %v", p.tok.Val)
				return nil
			}
			fmtOffset := fmt.Sprintf("%v%04d", sign, offset*100)
			if p.parsed.Offset == "" {
				p.parsed.Offset = fmtOffset
			} else {
				if p.parsed.Offset != fmtOffset {
					p.trace("offset mismatch: %v != %v", p.parsed.Offset, fmtOffset)
					return nil
				}
			}
			return nil
		}

		p.trace("is zone")
		var offset string
		var ok bool

		offset, ok = p.loc.ZoneNamesShort[p.tok.Val]
		if !ok {
			p.trace("zone not recognized")
			return nil
		}
		p.parsed.Zone = p.tok.Val
		if p.parsed.Offset != "" && p.parsed.Offset != offset {
			return p.err("time zone '%v' does not match given offset '%v'", p.tok.Val, p.parsed.Offset)
		}
		p.parsed.Offset = offset
		return nil
	}

	return p.err("unexpected text: %v", p.tok.Val)
}

func (p *Parser) parseNumber() error {
	if p.state == unknown {
		la := p.lookahead(1)
		if la.Type == Indicator && (inSet(la.Val, p.loc.TimeSep) || inSet(la.Val, p.loc.HourSep)) {
			p.changeState(parsingTime)
		} else {
			p.changeState(parsingDate)
		}
	}
	if p.state == parsingDate {
		la := p.lookahead(1)
		if la.Type == Indicator && inSet(la.Val, p.loc.TimeSep) {
			p.changeState(parsingTime)

		} else {
			return p.parseNumberDate()
		}
	}
	if p.state == parsingTime {
		return p.parseNumberTime()
	}
	if p.state == parsingZone {
		p.changeState(done)
		return p.parseYear4()
	}
	return p.err("extra number: %v", p.tok.Val)
}

func (p *Parser) parseNumberDate() error {
	sep := p.parsed.DateSep
	if sep == "" {
		la := p.lookahead(1)
		if la.Type == Indicator {
			if inSet(la.Val, p.loc.DateSep) {
				sep = la.Val
			}
		} else {
			sep = " "
		}
		p.trace("DateSep = '%v'", sep)
		p.parsed.DateSep = sep
	}
	return p.parseDate()
}

func (p *Parser) parseNumberTime() error {
	sep := p.parsed.TimeSep
	if sep == "" && p.parsed.HourSep == "" {
		la := p.lookahead(1)
		if la.Val != "" {
			if inSet(la.Val, p.loc.TimeSep) {
				sep = la.Val
			} else {
				if inSet(la.Val, p.loc.HourSep) {
					sep = ""
				}
			}
			p.trace("TimeSep = '%v'", sep)
			p.parsed.TimeSep = sep
		}
	}
	return p.parseTime()
}

func (p *Parser) parseIndicator() error {
	if p.state == parsingDate && p.tok.Val == p.parsed.DateSep {
		p.next()
		return p.parseDate()
	}
	if p.state == parsingTime {
		if p.tok.Val == p.parsed.TimeSep {
			p.next()
			return p.parseTime()
		}
		if p.tok.Val == "-" || p.tok.Val == "+" {
			p.changeState(parsingZone)
		}
	}
	if p.state == parsingZone {
		if p.tok.Val == "-" || p.tok.Val == "+" {
			return p.parseOffset()
		}
	}
	p.trace("discarding")
	return nil
}

func (p *Parser) changeState(newState state) {
	if p.parseOne {
		if newState != parsingZone {
			newState = done
		}
	}
	if p.state != newState {
		p.trace("state: %v -> %v", p.state, newState)
	}
	p.state = newState
}

func (p *Parser) parseDate() error {
	delim := p.parsed.DateSep
	if p.dateOrder == unknownOrder {
		la1 := p.lookahead(1)
		_, la1IsMonth := lookupMonth(p.loc, la1.Val)
		la2 := p.lookahead(2)
		_, la2IsMonth := lookupMonth(p.loc, la2.Val)

		switch {
		case delim == "-" && la2IsMonth:
			p.dateOrder = dayMonthYearOrder
		case delim == "-" && la2.Type == Number && len(la2.Val) == 3:
			p.dateOrder = yearDayOrder
		case delim == "-":
			p.dateOrder = yearMonthDayOrder
		case la1IsMonth:
			p.dateOrder = dayMonthYearOrder
		case p.loc.MonthDayOrder:
			p.dateOrder = monthDayYearOrder
		default:
			p.dateOrder = dayMonthYearOrder
		}
		p.trace("order: %v", p.dateOrder)
	}
	switch p.dateOrder {
	case yearDayOrder:
		return p.parseYearDay()
	case yearMonthDayOrder:
		return p.parseYearMonthDay()
	case dayMonthYearOrder:
		return p.parseDayMonthYear()
	case monthDayYearOrder:
		return p.parseMonthDayYear()
	}
	return p.err("unexpected '%v' in date", p.tok.Val)
}

func (p *Parser) parseYearMonthDay() error {
	if p.parsed.Year == "" {
		return p.parseYear4()
	}
	if p.parsed.Month == "" {
		return p.parseMonth()
	}
	if p.parsed.Day == "" {
		return p.parseDay()
	}
	return p.err("pass parseYearDayMonth")
}

func (p *Parser) parseYearDay() error {
	if p.parsed.Year == "" {
		return p.parseYear4()
	}
	if p.parsed.Day == "" {
		return p.parseOrdinalDay()
	}
	return p.err("pass parseYearDayMonth")
}

func (p *Parser) parseDayMonthYear() error {
	if p.parsed.Day == "" {
		return p.parseDay()
	}
	if p.parsed.Month == "" {
		return p.parseMonth()
	}
	if p.parsed.Year == "" {
		return p.parseYear()
	}
	return p.err("pass parseDayMonth")
}

func (p *Parser) parseMonthDayYear() error {
	if p.parsed.Month == "" {
		return p.parseMonth()
	}
	if p.parsed.Day == "" {
		return p.parseDay()
	}
	if p.parsed.Year == "" {
		return p.parseYear()
	}
	return p.err("pass parseMonthDay")
}

func (p *Parser) parseYear() error {
	p.trace("is year")
	p.parsed.Year = p.tok.Val
	switch len(p.parsed.Year) {
	case 4:
		//
	case 2:
		//
	default:
		return p.err("invalid year: %v", p.parsed.Year)
	}
	return nil
}

func (p *Parser) parseYear4() error {
	p.trace("is year4")
	p.parsed.Year = p.tok.Val
	if len(p.parsed.Year) != 4 {
		return p.err("invalid year: %v", p.parsed.Year)
	}
	return nil
}

func (p *Parser) parseMonth() error {
	p.trace("is month")
	p.parsed.Month = p.tok.Val
	if _, ok := lookupMonth(p.loc, p.tok.Val); ok {
		return nil
	}
	m, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid month: %v", p.tok.Val)
	}
	if m < 1 || m > 12 {
		return p.err("invalid month: %v", p.tok.Val)
	}
	return nil
}

func (p *Parser) parseDay() error {
	p.trace("is day")
	p.parsed.Day = p.tok.Val
	d, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid day: %v", p.tok.Val)
	}
	if d < 1 || d > 31 {
		return p.err("invalid day: %v", p.tok.Val)
	}
	return nil
}

func (p *Parser) parseOrdinalDay() error {
	p.trace("is ordinal day")
	p.parsed.Day = p.tok.Val
	d, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid day: %v", p.tok.Val)
	}
	if d < 1 || d > 365 {
		return p.err("invalid day: %v", p.tok.Val)
	}
	return nil
}

func (p *Parser) parseTime() error {
	if p.parsed.Hour == "" {
		return p.parseHour()
	}
	if p.parsed.Minute == "" {
		return p.parseMinute()
	}
	if p.parsed.Second == "" {
		return p.parseSecond()
	}
	p.changeState(done)
	return p.parseYear4()
}

func (p *Parser) parseHour() error {
	p.trace("is hour")
	p.parsed.Hour = p.tok.Val
	h, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid hour: %v", p.tok.Val)
	}
	if h < 0 || h >= 24 {
		return p.err("invalid hour: %v", p.tok.Val)
	}

	la := p.lookahead(1)
	if inSet(la.Val, p.loc.HourSep) {
		p.trace("HourSep = '%v'", la.Val)
		p.parsed.HourSep = la.Val
		p.next()
	}
	return nil
}

func (p *Parser) parseMinute() error {
	p.trace("is minute")
	p.parsed.Minute = p.tok.Val
	m, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid minute: %v", p.tok.Val)
	}
	if m < 0 || m >= 60 {
		return p.err("invalid minute: %v", p.tok.Val)
	}
	return nil
}

func (p *Parser) parseSecond() error {
	p.trace("is second")
	p.parsed.Second = p.tok.Val
	s, err := strconv.Atoi(p.tok.Val)
	if err != nil {
		return p.err("invalid second: %v", p.tok.Val)
	}
	if s < 0 || s >= 60 {
		return p.err("invalid second: %v", p.tok.Val)
	}

	la := p.lookahead(1)
	if la.Type == Indicator && la.Val == p.loc.DecimalSep {
		p.trace("has fractions")
		p.next()
		p.next()
		p.parsed.FracSecond = p.tok.Val
	}
	return nil
}

func (p *Parser) parseOffset() error {
	p.trace("is offset")
	var parts []string
	if p.tok.Type == Indicator && (p.tok.Val == "+" || p.tok.Val == "-") {
		parts = append(parts, p.tok.Val)
		p.next()
	}
	if p.tok.Type != Number {
		return p.err("expecting offset but got '%v'", p.tok.Val)
	}
	if len(p.tok.Val) == 4 {
		parts = append(parts, p.tok.Val)
	}
	if len(p.tok.Val) == 2 {
		parts = append(parts, p.tok.Val)
		p.next()
		if p.tok.Type != Indicator || p.tok.Val != ":" {
			return p.err("expecting ':' in offset but got '%v'", p.tok.Val)
		}
		p.next()
		if p.tok.Type != Number {
			return p.err("expecting offset minutes but got '%v'", p.tok.Val)
		}
		parts = append(parts, p.tok.Val)
	}
	offset := strings.Join(parts, "")
	if p.parsed.Offset != "" && p.parsed.Offset != offset {
		return p.err("offset mismatch between '%v' and '%v'", offset, p.parsed.Offset)
	}
	p.parsed.Offset = offset
	return nil
}

func (p *Parser) lookahead(n int) Token {
	if n+p.idx >= len(p.tokens) {
		return Token{End, "", 0}
	}
	return p.tokens[n+p.idx]
}

func (p *Parser) next() {
	p.idx++
	if p.idx >= len(p.tokens) {
		p.trace("end")
		p.idx = len(p.tokens)
		p.tok = Token{End, "", 0}
		return
	}
	p.tok = p.tokens[p.idx]
	p.trace("next: %v", p.tok)
}

func (p *Parser) err(format string, a ...any) error {
	return fmt.Errorf(format, a...)
}

func (p *Parser) trace(format string, a ...any) {
	if p.Trace {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

func inSet(text string, domain []string) bool {
	text = strings.ToLower(text)
	for _, v := range domain {
		if text == strings.ToLower(v) {
			return true
		}
	}
	return false
}

func lookupMonth(l *locale.Locale, text string) (string, bool) {
	n, ok := l.MonthNum[l.Key(text)]
	if !ok {
		return "", false
	}
	return l.MonthNamesAbbr[n-1], true
}

func lookupDay(l *locale.Locale, text string) (string, bool) {
	n, ok := l.DayNum[l.Key(text)]
	if !ok {
		return "", false
	}
	return l.DayNamesAbbr[n], true
}

func lookupPeriod(l *locale.Locale, text string) (string, bool) {
	n, ok := l.PeriodNum[l.Key(text)]
	if !ok {
		return "", false
	}
	return l.PeriodNamesAbbr[n][0], true
}
