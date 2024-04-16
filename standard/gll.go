package standard

import "github.com/twpayne/go-nmea"

type GLL struct {
	nmea.Address
	Lat       nmea.Optional[float64]
	Lon       nmea.Optional[float64]
	TimeOfDay nmea.TimeOfDay
	Status    byte
	PosMode   byte
}

func ParseGLL(addr string, tok *nmea.Tokenizer) (*GLL, error) {
	var gll GLL
	gll.Address = nmea.NewAddress(addr)
	gll.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	gll.Lon = tok.CommaOptionalLonDegMinCommaHemi()
	gll.TimeOfDay = tok.CommaTimeOfDay()
	gll.Status = tok.CommaOneByteOf("AV")
	gll.PosMode = tok.CommaOneByteOf("ADEFNR")
	return &gll, tok.Err()
}
