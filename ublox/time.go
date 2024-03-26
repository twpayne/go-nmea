package ublox

import (
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/gps"
)

type Time struct {
	address              Address
	Time                 time.Time
	UTCTimeOfWeek        float64
	UTCWeek              int
	LeapSeconds          int
	LeapSecondsDefault   bool
	ClockBias            int
	ClockDrift           float64
	TimePulseGranularity int
}

func ParseTime(addr string, tok *nmea.Tokenizer) (*Time, error) {
	var t Time
	t.address = NewAddress(addr)
	timeOfDay := gps.ParseCommaTimeOfDay(tok)
	date := gps.ParseCommaDate(tok)
	t.Time = time.Date(date.Year, date.Month, date.Day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	t.UTCTimeOfWeek = tok.CommaUnsignedFloat()
	t.UTCWeek = tok.CommaUnsignedInt()
	t.LeapSeconds = tok.CommaUnsignedInt()
	if b, ok := tok.Peek(); ok && b == 'D' {
		tok.LiteralByte('D')
		t.LeapSecondsDefault = true
	}
	t.ClockBias = tok.CommaUnsignedInt()
	t.ClockDrift = tok.CommaFloat()
	t.TimePulseGranularity = tok.CommaUnsignedInt()
	tok.Comma()
	tok.EndOfData()
	return &t, tok.Err()
}

func (t Time) Address() nmea.Addresser {
	return t.address
}
