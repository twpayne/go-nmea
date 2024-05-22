package lxnavigation

import "github.com/twpayne/go-nmea"

type LXWP2 struct {
	nmea.Address
	MacCreadyFactor float64
	LoadFactor      float64
	BugsPercent     int
	PolarA          float64
	PolarB          float64
	PolarC          float64
	VolumePercent   int
}

func ParseLXWP2(addr string, tok *nmea.Tokenizer) (*LXWP2, error) {
	var lxwp2 LXWP2
	lxwp2.Address = nmea.NewAddress(addr)
	lxwp2.MacCreadyFactor = tok.CommaFloat()
	lxwp2.LoadFactor = tok.CommaFloat()
	lxwp2.BugsPercent = tok.CommaInt()
	lxwp2.PolarA = tok.CommaFloat()
	lxwp2.PolarB = tok.CommaFloat()
	lxwp2.PolarC = tok.CommaFloat()
	lxwp2.VolumePercent = tok.CommaInt()
	tok.EndOfData()
	return &lxwp2, tok.Err()
}
