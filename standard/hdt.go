package standard

import "github.com/twpayne/go-nmea"

type HDT struct {
	nmea.Address
	HeadingTrue float64
}

func ParseHDT(addr string, tok *nmea.Tokenizer) (*HDT, error) {
	var hdt HDT
	hdt.Address = nmea.NewAddress(addr)
	hdt.HeadingTrue = tok.CommaFloatCommaUnit('T')
	tok.EndOfData()
	return &hdt, tok.Err()
}
