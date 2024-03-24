package nmeagps

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type RMC struct {
	address           Address
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
	rmc.address = Address(addr)
	timeOfDay := ParseCommaTimeOfDay(tok)
	rmc.Status = tok.CommaOneByteOf(statuses)
	rmc.Lat = commaLatDegMinCommaHemi(tok)
	rmc.Lon = commaLonDegMinCommaHemi(tok)
	rmc.SpeedOverGroundKN = tok.CommaUnsignedFloat()
	rmc.CourseOverGround = tok.CommaOptionalUnsignedFloat()
	tok.Comma()
	day := tok.DecimalDigits(2)
	month := time.Month(tok.DecimalDigits(2))
	year := 1900 + tok.DecimalDigits(2)
	if year < 1993 {
		year += 100
	}
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
	rmc.Time = time.Date(year, month, day, timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond, time.UTC)
	return &rmc, tok.Err()
}

func (rmc *RMC) Address() nmea.Address {
	return rmc.address
}
