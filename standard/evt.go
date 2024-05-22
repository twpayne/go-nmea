package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type EVT struct {
	nmea.Address
	Day       int
	Month     time.Month
	Year      int
	TimeOfDay nmea.TimeOfDay
	Unknown1  int
	Unknown2  struct{}
	Unknown3  struct{}
	Unknown4  struct{}
}

func ParseEVT(addr string, tok *nmea.Tokenizer) (*EVT, error) {
	var evt EVT
	evt.Address = nmea.NewAddress(addr)
	tok.Comma()
	evt.Day = tok.DecimalDigits(2)
	evt.Month = time.Month(tok.DecimalDigits(2))
	evt.Year = tok.DecimalDigits(4)
	evt.TimeOfDay = tok.CommaTimeOfDay()
	evt.Unknown1 = tok.CommaInt()
	evt.Unknown2 = tok.CommaEmpty()
	evt.Unknown3 = tok.CommaEmpty()
	evt.Unknown4 = tok.CommaEmpty()
	tok.EndOfData()
	return &evt, tok.Err()
}

func (evt *EVT) Time() time.Time {
	return time.Date(
		evt.Year, evt.Month, evt.Day,
		evt.TimeOfDay.Hour, evt.TimeOfDay.Minute, evt.TimeOfDay.Second, evt.TimeOfDay.Nanosecond,
		time.UTC,
	)
}
