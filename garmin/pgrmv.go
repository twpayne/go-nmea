package garmin

import "github.com/twpayne/go-nmea"

type PGRMV struct {
	nmea.Address
	TrueEastVelocity  float64
	TrueNorthVelocity float64
	UpVelocity        float64
}

func ParsePGRMV(addr string, tok *nmea.Tokenizer) (*PGRMV, error) {
	var pgrmv PGRMV
	pgrmv.Address = nmea.NewAddress(addr)
	pgrmv.TrueEastVelocity = tok.CommaFloat()
	pgrmv.TrueNorthVelocity = tok.CommaFloat()
	pgrmv.UpVelocity = tok.CommaFloat()
	tok.EndOfData()
	return &pgrmv, tok.Err()
}
