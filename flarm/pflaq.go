package flarm

import "github.com/twpayne/go-nmea"

type PFLAQ struct {
	nmea.Address
	Operation string
	Info      nmea.Optional[string]
	Progress  int
}

func ParsePFLAQ(addr string, tok *nmea.Tokenizer) (*PFLAQ, error) {
	var pflaq PFLAQ
	pflaq.Address = nmea.NewAddress(addr)
	pflaq.Operation = tok.CommaString()
	tokFork := tok.Fork()
	pflaq.Info = nmea.NewOptional(tok.CommaString())
	pflaq.Progress = tok.CommaUnsignedInt()
	tok.EndOfData()
	if tok.Err() != nil {
		tok = tokFork
		pflaq.Info = nmea.Optional[string]{}
		pflaq.Progress = tok.CommaUnsignedInt()
		tok.EndOfData()
	}
	return &pflaq, tok.Err()
}
