package garmin

import "github.com/twpayne/go-nmea"

type PGRMV struct {
	address           nmea.Address
	TrueEastVelocity  float64
	TrueNorthVelocity float64
	UpVelocity        float64
}

func ParsePGRMV(addr string, tok *nmea.Tokenizer) (*PGRMV, error) {
	var v PGRMV
	v.address = nmea.NewAddress(addr)
	v.TrueEastVelocity = tok.CommaFloat()
	v.TrueNorthVelocity = tok.CommaFloat()
	v.UpVelocity = tok.CommaFloat()
	tok.EndOfData()
	return &v, tok.Err()
}

func (v PGRMV) Address() nmea.Addresser {
	return v.address
}
