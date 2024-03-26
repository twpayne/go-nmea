package garmin

import "github.com/twpayne/go-nmea"

type PGRMM struct {
	address nmea.Address
	Datum   string
}

func ParsePGRMM(addr string, tok *nmea.Tokenizer) (*PGRMM, error) {
	var m PGRMM
	m.address = nmea.NewAddress(addr)
	m.Datum = tok.CommaString()
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PGRMM) Address() nmea.Addresser {
	return m.address
}
