package standard

import "github.com/twpayne/go-nmea"

type VHW struct {
	nmea.Address
	HeadingTrue     nmea.Optional[float64]
	HeadingMagnetic nmea.Optional[float64]
	SpeedKnots      nmea.Optional[float64]
	SpeedKPH        nmea.Optional[float64]
}

func ParseVHW(addr string, tok *nmea.Tokenizer) (*VHW, error) {
	var vhw VHW
	vhw.Address = nmea.NewAddress(addr)
	vhw.HeadingTrue = tok.CommaOptionalFloatCommaUnit('T')
	vhw.HeadingMagnetic = tok.CommaOptionalFloatCommaUnit('M')
	vhw.SpeedKnots = tok.CommaOptionalFloatCommaUnit('N')
	vhw.SpeedKPH = tok.CommaOptionalFloatCommaUnit('K')
	tok.EndOfData()
	return &vhw, tok.Err()
}
