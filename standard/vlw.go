package standard

import "github.com/twpayne/go-nmea"

type VLW struct {
	nmea.Address
	TotalWaterDistanceNM  nmea.Optional[float64]
	WaterDistanceNM       nmea.Optional[float64]
	TotalGroundDistanceNM nmea.Optional[float64]
	GroundDistanceNM      nmea.Optional[float64]
}

func ParseVLW(addr string, tok *nmea.Tokenizer) (*VLW, error) {
	var vlw VLW
	vlw.Address = nmea.NewAddress(addr)
	vlw.TotalWaterDistanceNM = tok.CommaOptionalFloatCommaUnit('N')
	vlw.WaterDistanceNM = tok.CommaOptionalFloatCommaUnit('N')
	if !tok.AtEndOfData() {
		vlw.TotalGroundDistanceNM = tok.CommaOptionalFloatCommaUnit('N')
		vlw.GroundDistanceNM = tok.CommaOptionalFloatCommaUnit('N')
	}
	tok.EndOfData()
	return &vlw, tok.Err()
}
