package nmeapgrm

import "github.com/twpayne/go-nmea"

type PGRME struct {
	address                 nmea.Addresser
	HorizontalPositionError float64
	VerticalPositionError   float64
	PositionError           float64
}

func ParsePGRME(addr string, tok *nmea.Tokenizer) (*PGRME, error) {
	var e PGRME
	e.address = nmea.NewAddress(addr)
	e.HorizontalPositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	e.VerticalPositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	e.PositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	tok.EndOfData()
	return &e, tok.Err()
}

func (e PGRME) Address() nmea.Addresser {
	return e.address
}
