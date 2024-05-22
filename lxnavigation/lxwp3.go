package lxnavigation

import (
	"github.com/twpayne/go-nmea"
)

type LXWP3 struct {
	nmea.Address
	AltOffset       float64
	Mode            int
	Filter          float64
	Reserved1       struct{}
	TELevel         float64
	IntegrationTime float64
	Range           float64
	Silence         float64
	SwitchMode      int
	Speed           float64
	Unknown         struct{} // undocumented
	PolarName       string
	Reserved2       struct{}
}

func ParseLXWP3(addr string, tok *nmea.Tokenizer) (*LXWP3, error) {
	var lxwp3 LXWP3
	lxwp3.Address = nmea.NewAddress(addr)
	lxwp3.AltOffset = tok.CommaFloat()
	lxwp3.Mode = tok.CommaInt()
	lxwp3.Filter = tok.CommaFloat()
	lxwp3.Reserved1 = tok.CommaIgnore()
	lxwp3.TELevel = tok.CommaFloat()
	lxwp3.IntegrationTime = tok.CommaUnsignedFloat()
	lxwp3.Range = tok.CommaFloat()
	lxwp3.Silence = tok.CommaFloat()
	lxwp3.SwitchMode = tok.CommaInt()
	lxwp3.Speed = tok.CommaFloat()
	lxwp3.Unknown = tok.CommaIgnore()
	lxwp3.PolarName = tok.CommaString()
	lxwp3.Reserved2 = tok.CommaIgnore()
	tok.EndOfData()
	return &lxwp3, tok.Err()
}
