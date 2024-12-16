package nmea

// FIXME what happens to second on GPS leap second?

import (
	"fmt"
	"time"
)

type TimeOfDay struct {
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
}

func (t TimeOfDay) String() string {
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

func (t *Tokenizer) CommaOptionalTimeOfDay() Optional[TimeOfDay] {
	t.Comma()
	return t.OptionalTimeOfDay()
}

func (t *Tokenizer) CommaTimeOfDay() TimeOfDay {
	t.Comma()
	return t.TimeOfDay()
}

func (t *Tokenizer) OptionalTimeOfDay() Optional[TimeOfDay] {
	if t.err != nil {
		return Optional[TimeOfDay]{}
	}
	if t.pos == len(t.data) {
		return Optional[TimeOfDay]{}
	}
	if t.data[t.pos] == ',' {
		return Optional[TimeOfDay]{}
	}
	return NewOptional(t.TimeOfDay())
}

func (t *Tokenizer) TimeOfDay() TimeOfDay {
	hour := t.DecimalDigits(2)
	minute := t.DecimalDigits(2)
	second, nanosecond := secondPointNanosecond(t)
	return TimeOfDay{
		Hour:       hour,
		Minute:     minute,
		Second:     second,
		Nanosecond: nanosecond,
	}
}

func secondPointNanosecond(tok *Tokenizer) (int, int) {
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
