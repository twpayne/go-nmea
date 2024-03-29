package nmea

import "time"

type Date struct {
	Day   int
	Month time.Month
	Year  int
}

func ParseCommaDate(tok *Tokenizer) Date {
	tok.Comma()
	return ParseDate(tok)
}

func ParseDate(tok *Tokenizer) Date {
	day := tok.DecimalDigits(2)
	month := time.Month(tok.DecimalDigits(2))
	year := 1900 + tok.DecimalDigits(2)
	if year < 1993 {
		year += 100
	}
	return Date{
		Day:   day,
		Month: month,
		Year:  year,
	}
}
