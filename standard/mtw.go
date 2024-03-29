package standard

import "github.com/twpayne/go-nmea"

type MTW struct {
	nmea.Address
	Temperature float64
}

func ParseMTW(addr string, tok *nmea.Tokenizer) (*MTW, error) {
	var mtw MTW
	mtw.Address = nmea.NewAddress(addr)
	mtw.Temperature = tok.CommaFloatCommaUnit('C')
	tok.EndOfData()
	return &mtw, tok.Err()
}
