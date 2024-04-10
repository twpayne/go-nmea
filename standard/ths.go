package standard

import "github.com/twpayne/go-nmea"

type THS struct {
	nmea.Address
	HeadingTrue   float64
	ModeIndicator byte
}

func ParseTHS(addr string, tok *nmea.Tokenizer) (*THS, error) {
	var ths THS
	ths.Address = nmea.NewAddress(addr)
	ths.HeadingTrue = tok.CommaUnsignedFloat()
	ths.ModeIndicator = tok.CommaOneByteOf("AEMSV")
	tok.EndOfData()
	return &ths, tok.Err()
}
