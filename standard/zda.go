package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type ZDA struct {
	nmea.Address
	Time                 time.Time
	LocalTimeZoneHours   nmea.Optional[int]
	LocalTimeZoneMinutes nmea.Optional[int]
}

func ParseZDA(addr string, tok *nmea.Tokenizer) (*ZDA, error) {
	var zda ZDA
	zda.Address = nmea.NewAddress(addr)
	timeOfDay := tok.CommaTimeOfDay()
	day := tok.CommaUnsignedInt()
	month := time.Month(tok.CommaUnsignedInt())
	year := tok.CommaUnsignedInt()
	zda.Time = time.Date(year, month, day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	zda.LocalTimeZoneHours = tok.CommaOptionalInt()
	zda.LocalTimeZoneMinutes = tok.CommaOptionalInt()
	tok.EndOfData()
	return &zda, tok.Err()
}
