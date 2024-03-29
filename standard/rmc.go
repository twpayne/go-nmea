package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type RMC struct {
	nmea.Address
	Time              time.Time
	Status            byte
	Lat               float64
	Lon               float64
	SpeedOverGroundKN float64
	CourseOverGround  nmea.Optional[float64]
	MagneticVariation nmea.Optional[float64]
	ModeIndicator     byte
	NavStatus         byte
}

func ParseRMC(addr string, tok *nmea.Tokenizer) (*RMC, error) {
	var rmc RMC
	rmc.Address = nmea.NewAddress(addr)
	timeOfDay := nmea.ParseCommaTimeOfDay(tok)
	rmc.Status = tok.CommaOneByteOf(statuses)
	rmc.Lat = tok.CommaLatDegMinCommaHemi()
	rmc.Lon = tok.CommaLonDegMinCommaHemi()
	rmc.SpeedOverGroundKN = tok.CommaUnsignedFloat()
	rmc.CourseOverGround = tok.CommaOptionalUnsignedFloat()
	date := nmea.ParseCommaDate(tok)
	rmc.Time = time.Date(date.Year, date.Month, date.Day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	rmc.MagneticVariation = tok.CommaOptionalFloat()
	if rmc.MagneticVariation.Valid {
		hemisphere := tok.CommaOneByteOf("EW")
		if rmc.MagneticVariation.Valid && hemisphere == 'W' {
			rmc.MagneticVariation.Value = -rmc.MagneticVariation.Value
		}
	} else {
		tok.Comma()
	}
	if !tok.AtEndOfData() {
		rmc.ModeIndicator = tok.CommaOneByteOf(statuses)
		rmc.NavStatus = tok.CommaOneByteOf("V")
	}
	tok.EndOfData()
	return &rmc, tok.Err()
}
