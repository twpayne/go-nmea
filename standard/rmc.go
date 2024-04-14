package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type RMC struct {
	nmea.Address
	Time              time.Time
	Status            byte
	Lat               nmea.Optional[float64]
	Lon               nmea.Optional[float64]
	SpeedOverGroundKN nmea.Optional[float64]
	CourseOverGround  nmea.Optional[float64]
	MagneticVariation nmea.Optional[float64]
	ModeIndicator     nmea.Optional[byte]
	NavStatus         nmea.Optional[byte]
}

func ParseRMC(addr string, tok *nmea.Tokenizer) (*RMC, error) {
	var rmc RMC
	rmc.Address = nmea.NewAddress(addr)
	timeOfDay := nmea.ParseCommaTimeOfDay(tok)
	rmc.Status = tok.CommaOneByteOf("AV")
	rmc.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	rmc.Lon = tok.CommaOptionalLonDegMinCommaHemi()
	rmc.SpeedOverGroundKN = tok.CommaOptionalUnsignedFloat()
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
		rmc.ModeIndicator = nmea.NewOptional(tok.CommaOneByteOf("ADEMN"))
	}
	if !tok.AtEndOfData() {
		rmc.NavStatus = nmea.NewOptional(tok.CommaOneByteOf("V"))
	}
	tok.EndOfData()
	return &rmc, tok.Err()
}
