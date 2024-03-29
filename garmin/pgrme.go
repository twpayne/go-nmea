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
	e.HorizontalPositionError = tok.CommaFloatCommaUnit('M')
	e.VerticalPositionError = tok.CommaFloatCommaUnit('M')
	e.PositionError = tok.CommaFloatCommaUnit('M')
	tok.EndOfData()
	return &e, tok.Err()
}
