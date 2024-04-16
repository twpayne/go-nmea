package standard

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type RMC struct {
	nmea.Address
	TimeOfDay         nmea.Optional[nmea.TimeOfDay]
	Status            byte
	Lat               nmea.Optional[float64]
	Lon               nmea.Optional[float64]
	SpeedOverGroundKN nmea.Optional[float64]
	CourseOverGround  nmea.Optional[float64]
	Date              nmea.Optional[nmea.Date]
	MagneticVariation nmea.Optional[float64]
	ModeIndicator     nmea.Optional[byte]
	NavStatus         nmea.Optional[byte]
}

func ParseRMC(addr string, tok *nmea.Tokenizer) (*RMC, error) {
	var rmc RMC
	rmc.Address = nmea.NewAddress(addr)
	rmc.TimeOfDay = tok.CommaOptionalTimeOfDay()
	rmc.Status = tok.CommaOneByteOf("AV")
	rmc.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	rmc.Lon = tok.CommaOptionalLonDegMinCommaHemi()
	rmc.SpeedOverGroundKN = tok.CommaOptionalUnsignedFloat()
	rmc.CourseOverGround = tok.CommaOptionalUnsignedFloat()
	rmc.Date = tok.CommaOptionalDate()
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

func (rmc *RMC) Time() nmea.Optional[time.Time] {
	if !rmc.TimeOfDay.Valid || !rmc.Date.Valid {
		return nmea.Optional[time.Time]{}
	}
	timeOfDay := rmc.TimeOfDay.Value
	date := rmc.Date.Value
	return nmea.NewOptional(time.Date(
		date.Year, date.Month, date.Day,
		timeOfDay.Hour, timeOfDay.Minute, timeOfDay.Second, timeOfDay.Nanosecond,
		time.UTC,
	))
}
