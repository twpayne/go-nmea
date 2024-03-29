package standard

import "github.com/twpayne/go-nmea"

type GGA struct {
	nmea.Address
	TimeOfDay                        nmea.TimeOfDay
	Lat                              float64
	Lon                              float64
	FixQuality                       int
	NumberOfSatellites               int
	HDOP                             float64
	Alt                              float64
	HeightOfGeoidAboveWGS84Ellipsoid float64
	TimeSinceLastDGPSUpdate          nmea.Optional[int]
	DGPSReferenceStationID           string
}

func ParseGGA(addr string, tok *nmea.Tokenizer) (*GGA, error) {
	var gga GGA
	gga.Address = nmea.NewAddress(addr)
	gga.TimeOfDay = nmea.ParseCommaTimeOfDay(tok)
	gga.Lat = tok.CommaLatDegMinCommaHemi()
	gga.Lon = tok.CommaLonDegMinCommaHemi()
	gga.FixQuality = tok.CommaUnsignedInt()
	gga.NumberOfSatellites = tok.CommaUnsignedInt()
	gga.HDOP = tok.CommaUnsignedFloat()
	gga.Alt = tok.CommaFloatCommaUnit('M')
	gga.HeightOfGeoidAboveWGS84Ellipsoid = tok.CommaFloatCommaUnit('M')
	gga.TimeSinceLastDGPSUpdate = tok.CommaOptionalUnsignedInt()
	gga.DGPSReferenceStationID = tok.CommaString()
	tok.EndOfData()
	return &gga, tok.Err()
}
