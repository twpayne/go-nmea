package samsung

import "github.com/twpayne/go-nmea"

type PSAMID struct {
	nmea.Address
	Unknown1 string
	Unknown2 struct{}
}

func ParsePSAMID(addr string, tok *nmea.Tokenizer) (*PSAMID, error) {
	var psamid PSAMID
	psamid.Address = nmea.NewAddress(addr)
	psamid.Unknown1 = tok.CommaString()
	psamid.Unknown2 = tok.CommaEmpty()
	tok.EndOfData()
	return &psamid, tok.Err()
}
