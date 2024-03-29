package garmin

import "github.com/twpayne/go-nmea"

type PGRME struct {
	nmea.Address
	HorizontalPositionError float64
	VerticalPositionError   float64
	PositionError           float64
}

func ParsePGRME(addr string, tok *nmea.Tokenizer) (*PGRME, error) {
	var e PGRME
	e.Address = nmea.NewAddress(addr)
	e.HorizontalPositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	e.VerticalPositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	e.PositionError = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('M')
	tok.EndOfData()
	return &e, tok.Err()
}
