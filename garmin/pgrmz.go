package garmin

import "github.com/twpayne/go-nmea"

type PGRMZ struct {
	address nmea.Address
	AltFeet float64
	FixType int
}

func ParsePGRMZ(addr string, tok *nmea.Tokenizer) (*PGRMZ, error) {
	var z PGRMZ
	z.address = nmea.NewAddress(addr)
	z.AltFeet = tok.CommaFloat()
	tok.CommaOneByteOf("Ff")
	z.FixType = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &z, tok.Err()
}

func (z PGRMZ) Address() nmea.Addresser {
	return z.address
}
