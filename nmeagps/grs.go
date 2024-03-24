package nmeagps

import "github.com/twpayne/go-nmea"

type GRS struct {
	address   Address
	TimeOfDay TimeOfDay
	Mode      int
	Residuals []nmea.Optional[float64]
	SystemID  int
	SignalID  int
}

func ParseGRS(addr string, tok *nmea.Tokenizer) (*GRS, error) {
	var grs GRS
	grs.address = Address(addr)
	grs.TimeOfDay = ParseCommaTimeOfDay(tok)
	grs.Mode = tok.CommaUnsignedInt()
	for i := 0; i < 12; i++ {
		residual := tok.CommaOptionalFloat()
		grs.Residuals = append(grs.Residuals, residual)
	}
	grs.SystemID = tok.CommaHex()
	grs.SignalID = tok.CommaHex()
	tok.EndOfData()
	return &grs, tok.Err()
}

func (grs GRS) Address() nmea.Address {
	return grs.address
}
