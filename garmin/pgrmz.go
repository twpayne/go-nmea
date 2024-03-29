package garmin

import "github.com/twpayne/go-nmea"

type PGRMZ struct {
	nmea.Address
	AltFeet float64
	FixType int
}

func ParsePGRMZ(addr string, tok *nmea.Tokenizer) (*PGRMZ, error) {
	var z PGRMZ
	z.Address = nmea.NewAddress(addr)
	z.AltFeet = tok.CommaFloat()
	tok.CommaOneByteOf("Ff")
	z.FixType = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &z, tok.Err()
}
