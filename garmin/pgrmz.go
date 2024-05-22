package garmin

import "github.com/twpayne/go-nmea"

type PGRMZ struct {
	nmea.Address
	AltFeet float64
	FixType int
}

func ParsePGRMZ(addr string, tok *nmea.Tokenizer) (*PGRMZ, error) {
	var pgrmz PGRMZ
	pgrmz.Address = nmea.NewAddress(addr)
	pgrmz.AltFeet = tok.CommaFloat()
	tok.CommaOneByteOf("Ff")
	pgrmz.FixType = tok.CommaUnsignedInt()
	tok.EndOfData()
	return &pgrmz, tok.Err()
}
