package garmin

import "github.com/twpayne/go-nmea"

type PGRMM struct {
	nmea.Address
	Datum string
}

func ParsePGRMM(addr string, tok *nmea.Tokenizer) (*PGRMM, error) {
	var pgrmm PGRMM
	pgrmm.Address = nmea.NewAddress(addr)
	pgrmm.Datum = tok.CommaString()
	tok.EndOfData()
	return &pgrmm, tok.Err()
}
