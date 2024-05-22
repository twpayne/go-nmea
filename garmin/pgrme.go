package garmin

import "github.com/twpayne/go-nmea"

type PGRME struct {
	nmea.Address
	HorizontalPositionError float64
	VerticalPositionError   float64
	PositionError           float64
}

func ParsePGRME(addr string, tok *nmea.Tokenizer) (*PGRME, error) {
	var pgrme PGRME
	pgrme.Address = nmea.NewAddress(addr)
	pgrme.HorizontalPositionError = tok.CommaFloatCommaUnit('M')
	pgrme.VerticalPositionError = tok.CommaFloatCommaUnit('M')
	pgrme.PositionError = tok.CommaFloatCommaUnit('M')
	tok.EndOfData()
	return &pgrme, tok.Err()
}
