package garmin

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type PGRMF struct {
	nmea.Address
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
	var pgrmf PGRMF
	pgrmf.Address = nmea.NewAddress(addr)
	pgrmf.GPSWeekNumber = tok.CommaUnsignedInt()
	pgrmf.GPSSeconds = tok.CommaUnsignedInt()
	date := tok.CommaDate()
	timeOfDay := tok.CommaTimeOfDay()
	pgrmf.Time = time.Date(date.Year, date.Month, date.Day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	pgrmf.LeapSeconds = tok.CommaUnsignedInt()
	pgrmf.Lat = tok.CommaLatDegMinCommaHemi()
	pgrmf.Lon = tok.CommaLonDegMinCommaHemi()
	pgrmf.Mode = tok.CommaOneByteOf("AM")
	pgrmf.FixType = tok.CommaUnsignedInt()
	pgrmf.SpeedOverGroundKPH = tok.CommaUnsignedInt()
	pgrmf.CourseOverGround = tok.CommaUnsignedInt()
	pgrmf.PDOP = tok.CommaUnsignedInt()
	pgrmf.TDOP = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &pgrmf, tok.Err()
}
