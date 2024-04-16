package nmea

import (
	"fmt"
	"time"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func (d Date) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (t *Tokenizer) CommaDate() Date {
	t.Comma()
	return t.Date()
}

func (t *Tokenizer) CommaOptionalDate() Optional[Date] {
	t.Comma()
	return t.OptionalDate()
}

func (t *Tokenizer) Date() Date {
	day := t.DecimalDigits(2)
	month := time.Month(t.DecimalDigits(2))
	year := 1900 + t.DecimalDigits(2)
	if year < 1993 {
		year += 100
	}
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (t *Tokenizer) OptionalDate() Optional[Date] {
	if t.err != nil {
		return Optional[Date]{}
	}
	if t.pos == len(t.data) {
		return Optional[Date]{}
	}
	if t.data[t.pos] == ',' {
		return Optional[Date]{}
	}
	return NewOptional(t.Date())
}
