package ublox

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type Time struct {
	nmea.Address
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
	t.Address = nmea.NewAddress(addr)
	timeOfDay := tok.CommaTimeOfDay()
	date := tok.CommaDate()
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
