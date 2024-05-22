package lxnavigation

import "github.com/twpayne/go-nmea"

type LXWP1 struct {
	nmea.Address
	DeviceName      string
	SerialNumber    int
	SoftwareVersion string
	HardwareVersion string
}

func ParseLXWP1(addr string, tok *nmea.Tokenizer) (*LXWP1, error) {
	var lxwp1 LXWP1
	lxwp1.Address = nmea.NewAddress(addr)
	lxwp1.DeviceName = tok.CommaString()
	lxwp1.SerialNumber = tok.CommaInt()
	lxwp1.SoftwareVersion = tok.CommaString()
	lxwp1.HardwareVersion = tok.CommaString()
	tok.EndOfData()
	return &lxwp1, tok.Err()
}
