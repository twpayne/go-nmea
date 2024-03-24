package nmeagps

import "github.com/twpayne/go-nmea"

type GGA struct {
	address                          Address
	TimeOfDay                        TimeOfDay
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
	gga.address = Address(addr)
	gga.TimeOfDay = ParseCommaTimeOfDay(tok)
	gga.Lat = commaLatDegMinCommaHemi(tok)
	gga.Lon = commaLonDegMinCommaHemi(tok)
	gga.FixQuality = tok.CommaUnsignedInt()
	gga.NumberOfSatellites = tok.CommaUnsignedInt()
	gga.HDOP = tok.CommaUnsignedFloat()
	gga.Alt = commaAltitudeCommaM(tok)
	gga.HeightOfGeoidAboveWGS84Ellipsoid = commaAltitudeCommaM(tok)
	gga.TimeSinceLastDGPSUpdate = tok.CommaOptionalUnsignedInt()
	gga.DGPSReferenceStationID = tok.CommaString()
	tok.EndOfData()
	return &gga, tok.Err()
}

func (gga GGA) Address() nmea.Address {
	return gga.address
}
