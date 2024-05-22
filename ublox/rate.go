package ublox

import "github.com/twpayne/go-nmea"

type Rate struct {
	nmea.Address
	MsgID    string
	RDDC     int
	RUS1     int
	RUS2     int
	RUSB     int
	RSPI     int
	Reserved struct{}
}

func ParseRate(addr string, tok *nmea.Tokenizer) (*Rate, error) {
	var r Rate
	r.Address = nmea.NewAddress(addr)
	r.MsgID = tok.CommaString()
	r.RDDC = tok.CommaUnsignedInt()
	r.RUS1 = tok.CommaUnsignedInt()
	r.RUS2 = tok.CommaUnsignedInt()
	r.RUSB = tok.CommaUnsignedInt()
	r.RSPI = tok.CommaUnsignedInt()
	r.Reserved = tok.CommaIgnore()
	tok.EndOfData()
	return &r, tok.Err()
}
