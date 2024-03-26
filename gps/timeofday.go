package gps

// FIXME what happens to second on GPS leap second?

import (
	"fmt"
	"time"

	"github.com/twpayne/go-nmea"
)

type TimeOfDay struct {
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
}

func ParseCommaTimeOfDay(tok *nmea.Tokenizer) TimeOfDay {
	tok.Comma()
	return ParseTimeOfDay(tok)
}

func ParseTimeOfDay(tok *nmea.Tokenizer) TimeOfDay {
	hour := tok.DecimalDigits(2)
	min := tok.DecimalDigits(2)
	sec, nsec := secondPointNanosecond(tok)
	return TimeOfDay{
		Hour:       hour,
		Minute:     min,
		Second:     sec,
		Nanosecond: nsec,
	}
}

func (t TimeOfDay) String() string {
	// FIXME strip trailing zeros
	return fmt.Sprintf("%02d:%02d:%02d.%09d", t.Hour, t.Minute, t.Second, t.Nanosecond)
}

func (t TimeOfDay) SinceMidnight() time.Duration {
	return time.Duration(t.Hour)*time.Hour +
		time.Duration(t.Minute)*time.Minute +
		time.Duration(t.Second)*time.Second +
		time.Duration(t.Nanosecond)*time.Nanosecond
}

func (t TimeOfDay) Valid() bool {
	if t.Hour < 0 || 23 < t.Hour {
		return false
	}
	if t.Minute < 0 || 59 < t.Minute {
		return false
	}
	// FIXME what happens to second on GPS leap second?
	if t.Second < 0 || 60 < t.Second {
		return false
	}
	if t.Nanosecond < 0 || 99999999 < t.Nanosecond {
		return false
	}
	return true
}

func secondPointNanosecond(tok *nmea.Tokenizer) (int, int) {
	sec := tok.DecimalDigits(2)
	numerator, denominator := tok.OptionalPointDecimal()
	for denominator < 1000000000 {
		numerator *= 10
		denominator *= 10
	}
	if denominator > 1000000000 {
		factor := denominator / 1000000000
		numerator = (numerator + factor/2) / factor
	}
	return sec, numerator
}
