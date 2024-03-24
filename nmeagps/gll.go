package nmeagps

import "github.com/twpayne/go-nmea"

type GLL struct {
	address Address
	Lat     float64
	Lon     float64
	TimeOfDay
	Status  byte
	PosMode byte
}

func ParseGLL(addr string, tok *nmea.Tokenizer) (*GLL, error) {
	var gll GLL
	gll.address = Address(addr)
	gll.Lat = commaLatDegMinCommaHemi(tok)
	gll.Lon = commaLonDegMinCommaHemi(tok)
	gll.TimeOfDay = ParseCommaTimeOfDay(tok)
	gll.Status = tok.CommaOneByteOf(statuses)
	gll.PosMode = tok.CommaOneByteOf(posModes)
	return &gll, tok.Err()
}

func (gll GLL) Address() nmea.Address {
	return gll.address
}
