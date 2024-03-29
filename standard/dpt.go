package standard

import "github.com/twpayne/go-nmea"

type DPT struct {
	nmea.Address
	Depth   float64
	Offset  nmea.Optional[float64]
	Maximum nmea.Optional[float64]
}

func ParseDPT(addr string, tok *nmea.Tokenizer) (*DPT, error) {
	var dpt DPT
	dpt.Address = nmea.NewAddress(addr)
	dpt.Depth = tok.CommaFloat()
	dpt.Offset = tok.CommaOptionalFloat()
	dpt.Maximum = tok.CommaOptionalFloat()
	tok.EndOfData()
	return &dpt, tok.Err()
}
