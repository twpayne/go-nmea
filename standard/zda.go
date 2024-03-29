package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type ZDA struct {
	nmea.Address
	Time                 time.Time
	LocalTimeZoneHours   int
	LocalTimeZoneMinutes int
}

func ParseZDA(addr string, tok *nmea.Tokenizer) (*ZDA, error) {
	var zda ZDA
	zda.Address = nmea.NewAddress(addr)
	timeOfDay := nmea.ParseCommaTimeOfDay(tok)
	day := tok.CommaUnsignedInt()
	month := time.Month(tok.CommaUnsignedInt())
	year := tok.CommaUnsignedInt()
	zda.Time = time.Date(year, month, day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	zda.LocalTimeZoneHours = tok.CommaInt()
	zda.LocalTimeZoneMinutes = tok.CommaInt()
	tok.EndOfData()
	return &zda, tok.Err()
}