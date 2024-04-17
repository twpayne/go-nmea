package standard

import "github.com/twpayne/go-nmea"

type GGA struct {
	nmea.Address
	TimeOfDay                        nmea.Optional[nmea.TimeOfDay]
	Lat                              nmea.Optional[float64]
	Lon                              nmea.Optional[float64]
	FixQuality                       int
	NumberOfSatellites               nmea.Optional[int]
	HDOP                             nmea.Optional[float64]
	Alt                              nmea.Optional[float64]
	HeightOfGeoidAboveWGS84Ellipsoid nmea.Optional[float64]
	TimeSinceLastDGPSUpdate          nmea.Optional[int]
	DGPSReferenceStationID           string
}

func ParseGGA(addr string, tok *nmea.Tokenizer) (*GGA, error) {
	var gga GGA
	gga.Address = nmea.NewAddress(addr)
	gga.TimeOfDay = tok.CommaOptionalTimeOfDay()
	gga.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	gga.Lon = tok.CommaOptionalLonDegMinCommaHemi()
	gga.FixQuality = tok.CommaUnsignedInt()
	gga.NumberOfSatellites = tok.CommaOptionalUnsignedInt()
	gga.HDOP = tok.CommaOptionalUnsignedFloat()
	gga.Alt = tok.CommaOptionalFloatCommaUnit('M')
	gga.HeightOfGeoidAboveWGS84Ellipsoid = tok.CommaOptionalFloatCommaUnit('M')
	gga.TimeSinceLastDGPSUpdate = tok.CommaOptionalUnsignedInt()
	gga.DGPSReferenceStationID = tok.CommaString()
	tok.EndOfData()
	return &gga, tok.Err()
}
