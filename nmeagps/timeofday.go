package nmeagps

// FIXME what happens to second on GPS leap second?

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/twpayne/go-nmea"
)

type TimeOfDay struct {
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
}

var secondsRx = regexp.MustCompile(`(\d{2})(?:\.(\d+))?`)

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

func secondPointNanosecond(tok *nmea.Tokenizer) (sec, nsec int) {
	// FIXME avoid use of regexp
	// FIXME nanosecond rounding
	match := tok.Regexp(secondsRx)
	if match == nil {
		return 0, 0
	}
	sec, _ = strconv.Atoi(string(match[1]))
	if len(match[2]) != 0 {
		nsec, _ = strconv.Atoi(string(match[2]))
		for i := 0; i < 9-len(match[2]); i++ {
			nsec *= 10
		}
	}
	return sec, nsec
}
