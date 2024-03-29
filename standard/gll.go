package standard

import "github.com/twpayne/go-nmea"

type GLL struct {
	nmea.Address
	Lat float64
	Lon float64
	nmea.TimeOfDay
	Status  byte
	PosMode byte
}

func ParseGLL(addr string, tok *nmea.Tokenizer) (*GLL, error) {
	var gll GLL
	gll.Address = nmea.NewAddress(addr)
	gll.Lat = tok.CommaLatDegMinCommaHemi()
	gll.Lon = tok.CommaLonDegMinCommaHemi()
	gll.TimeOfDay = nmea.ParseCommaTimeOfDay(tok)
	gll.Status = tok.CommaOneByteOf(statuses)
	gll.PosMode = tok.CommaOneByteOf(posModes)
	return &gll, tok.Err()
}
