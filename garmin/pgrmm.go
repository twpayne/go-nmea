package garmin

import "github.com/twpayne/go-nmea"

type PGRMM struct {
	nmea.Address
	Datum string
}

func ParsePGRMM(addr string, tok *nmea.Tokenizer) (*PGRMM, error) {
	var m PGRMM
	m.Address = nmea.NewAddress(addr)
	m.Datum = tok.CommaString()
	tok.EndOfData()
	return &m, tok.Err()
}
