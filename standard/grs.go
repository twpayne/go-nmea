package standard

import "github.com/twpayne/go-nmea"

type GRS struct {
	nmea.Address
	TimeOfDay nmea.TimeOfDay
	Mode      int
	Residuals []nmea.Optional[float64]
	SystemID  nmea.Optional[int]
	SignalID  nmea.Optional[int]
}

func ParseGRS(addr string, tok *nmea.Tokenizer) (*GRS, error) {
	var grs GRS
	grs.Address = nmea.NewAddress(addr)
	grs.TimeOfDay = tok.CommaTimeOfDay()
	grs.Mode = tok.CommaUnsignedInt()
	for range 12 {
		residual := tok.CommaOptionalFloat()
		grs.Residuals = append(grs.Residuals, residual)
	}
	if !tok.AtEndOfData() {
		grs.SystemID = tok.CommaOptionalHex()
	}
	if !tok.AtEndOfData() {
		grs.SignalID = tok.CommaOptionalHex()
	}
	tok.EndOfData()
	return &grs, tok.Err()
}
