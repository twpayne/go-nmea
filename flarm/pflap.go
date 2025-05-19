package flarm

import "github.com/twpayne/go-nmea"

// A PFLAP is a ping.
type PFLAP struct {
	nmea.Address
}

func ParsePFLAP(addr string, tok *nmea.Tokenizer) (*PFLAP, error) {
	var pflap PFLAP
	tok.CommaOneByteOf("A")
	tok.EndOfData()
	return &pflap, tok.Err()
}
