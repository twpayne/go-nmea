package samsung

import "github.com/twpayne/go-nmea"

type PSAMCLK struct {
	nmea.Address
	Unknown1 int
	Unknown2 int
	Unknown3 int
	Unknown4 int
	Unknown5 int
	Unknown6 int
}

func ParsePSAMCLK(addr string, tok *nmea.Tokenizer) (*PSAMCLK, error) {
	var psamclk PSAMCLK
	psamclk.Address = nmea.NewAddress(addr)
	psamclk.Unknown1 = tok.CommaInt()
	psamclk.Unknown2 = tok.CommaInt()
	psamclk.Unknown3 = tok.CommaInt()
	psamclk.Unknown4 = tok.CommaInt()
	psamclk.Unknown5 = tok.CommaInt()
	psamclk.Unknown6 = tok.CommaInt()
	tok.EndOfData()
	return &psamclk, tok.Err()
}
