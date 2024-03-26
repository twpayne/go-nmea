package gps

import "github.com/twpayne/go-nmea"

type VLW struct {
	address               Address
	TotalWaterDistanceNM  nmea.Optional[float64]
	WaterDistanceNM       nmea.Optional[float64]
	TotalGroundDistanceNM nmea.Optional[float64]
	GroundDistanceNM      nmea.Optional[float64]
}

func ParseVLW(addr string, tok *nmea.Tokenizer) (*VLW, error) {
	var vlw VLW
	vlw.address = NewAddress(addr)
	vlw.TotalWaterDistanceNM = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('N')
	vlw.WaterDistanceNM = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('N')
	vlw.TotalGroundDistanceNM = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('N')
	vlw.GroundDistanceNM = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('N')
	tok.EndOfData()
	return &vlw, tok.Err()
}

func (vlw VLW) Address() nmea.Addresser {
	return vlw.address
}
