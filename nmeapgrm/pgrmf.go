package nmeapgrm

import (
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeagps"
)

type PGRMF struct {
	address            nmea.Address
	GPSWeekNumber      int
	GPSSeconds         int
	Time               time.Time
	LeapSeconds        int
	Lat                float64
	Lon                float64
	Mode               byte
	FixType            int
	SpeedOverGroundKPH int
	CourseOverGround   int
	PDOP               int
	TDOP               int
}

func ParsePGRMF(addr string, tok *nmea.Tokenizer) (*PGRMF, error) {
	var f PGRMF
	f.address = nmea.NewAddress(addr)
	f.GPSWeekNumber = tok.CommaUnsignedInt()
	f.GPSSeconds = tok.CommaUnsignedInt()
	date := nmeagps.ParseCommaDate(tok)
	timeOfDay := nmeagps.ParseCommaTimeOfDay(tok)
	f.Time = time.Date(date.Year, date.Month, date.Day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	f.LeapSeconds = tok.CommaUnsignedInt()
	f.Lat = nmeagps.ParseCommaLatDegMinCommaHemi(tok)
	f.Lon = nmeagps.ParseCommaLonDegMinCommaHemi(tok)
	f.Mode = tok.CommaOneByteOf("AM")
	f.FixType = tok.CommaUnsignedInt()
	f.SpeedOverGroundKPH = tok.CommaUnsignedInt()
	f.CourseOverGround = tok.CommaUnsignedInt()
	f.PDOP = tok.CommaUnsignedInt()
	f.TDOP = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &f, tok.Err()
}

func (f PGRMF) Address() nmea.Addresser {
	return f.address
}
