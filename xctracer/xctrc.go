package xctracer

import (
	"time"

	"github.com/twpayne/go-nmea"
)

type XCTRC struct {
	nmea.Address
	Time                         time.Time
	Lat                          float64
	Lon                          float64
	Alt                          float64
	SpeedOverGround              float64
	CourseOverGround             float64
	ClimbRate                    float64
	WindSpeed                    nmea.Optional[float64]
	WindDirection                nmea.Optional[float64]
	TotalEnergyCompensationVario nmea.Optional[float64]
	RawPressure                  float64
	BatteryPercent               int
}

func ParseXCTRC(addr string, tok *nmea.Tokenizer) (*XCTRC, error) {
	var xctrc XCTRC
	xctrc.Address = nmea.NewAddress(addr)
	year := tok.CommaUnsignedInt()
	month := tok.CommaUnsignedInt()
	day := tok.CommaUnsignedInt()
	hour := tok.CommaUnsignedInt()
	minute := tok.CommaUnsignedInt()
	second := tok.CommaUnsignedInt()
	centisecond := tok.CommaUnsignedInt()
	xctrc.Time = time.Date(year, time.Month(month), day, hour, minute, second, centisecond*1e7, time.UTC)
	xctrc.Lat = tok.CommaFloat()
	xctrc.Lon = tok.CommaFloat()
	xctrc.Alt = tok.CommaFloat()
	xctrc.SpeedOverGround = tok.CommaFloat()
	xctrc.CourseOverGround = tok.CommaFloat()
	xctrc.ClimbRate = tok.CommaFloat()
	xctrc.WindSpeed = tok.CommaOptionalFloat()
	xctrc.WindDirection = tok.CommaOptionalFloat()
	xctrc.TotalEnergyCompensationVario = tok.CommaOptionalFloat()
	xctrc.RawPressure = tok.CommaFloat()
	xctrc.BatteryPercent = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &xctrc, tok.Err()
}
