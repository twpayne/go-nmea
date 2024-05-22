package samsung

import "github.com/twpayne/go-nmea"

type PSAMSA struct {
	nmea.Address
	TimeOfDay nmea.Optional[nmea.TimeOfDay]
	Lat       nmea.Optional[float64]
	Lon       nmea.Optional[float64]
	Unknown1  nmea.Optional[int] // maybe altitude over ground?
	Unknown2  struct{}
	Unknown3  nmea.Optional[int]
	Unknown4  struct{}
	Unknown5  struct{}
	Unknown6  struct{}
	Unknown7  struct{}
	Unknown8  struct{}
}

func ParsePSAMSA(addr string, tok *nmea.Tokenizer) (*PSAMSA, error) {
	var psamsa PSAMSA
	psamsa.Address = nmea.NewAddress(addr)
	psamsa.TimeOfDay = tok.CommaOptionalTimeOfDay()
	psamsa.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	psamsa.Lon = tok.CommaOptionalLonDegMinCommaHemi()
	psamsa.Unknown1 = tok.CommaOptionalIntCommaUnit('M')
	psamsa.Unknown2 = tok.CommaEmpty()
	psamsa.Unknown3 = tok.CommaOptionalInt()
	if !tok.AtEndOfData() {
		psamsa.Unknown4 = tok.CommaEmpty()
		psamsa.Unknown5 = tok.CommaEmpty()
		psamsa.Unknown6 = tok.CommaEmpty()
		psamsa.Unknown7 = tok.CommaEmpty()
		psamsa.Unknown8 = tok.CommaEmpty()
	}
	tok.EndOfData()
	return &psamsa, tok.Err()
}
