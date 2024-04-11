package flarm

import "github.com/twpayne/go-nmea"

type PFLAL struct {
	nmea.Address
	DebugMessage string
}

func ParsePFLAL(addr string, tok *nmea.Tokenizer) (*PFLAL, error) {
	var pflal PFLAL
	pflal.Address = nmea.NewAddress(addr)
	pflal.DebugMessage = tok.CommaString()
	tok.EndOfData()
	return &pflal, tok.Err()
}
