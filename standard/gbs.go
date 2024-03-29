package standard

import "github.com/twpayne/go-nmea"

type GBS struct {
	nmea.Address
	nmea.TimeOfDay
	ErrLat   float64
	ErrLon   float64
	ErrAlt   float64
	SVID     nmea.Optional[int]
	Prob     nmea.Optional[float64]
	Bias     nmea.Optional[float64]
	StdDev   nmea.Optional[float64]
	SystemID nmea.Optional[int]
	SignalID nmea.Optional[int]
}

func ParseGBS(addr string, tok *nmea.Tokenizer) (*GBS, error) {
	var gbs GBS
	gbs.Address = nmea.NewAddress(addr)
	gbs.TimeOfDay = nmea.ParseCommaTimeOfDay(tok)
	gbs.ErrLat = tok.CommaUnsignedFloat()
	gbs.ErrLon = tok.CommaUnsignedFloat()
	gbs.ErrAlt = tok.CommaUnsignedFloat()
	gbs.SVID = tok.CommaOptionalUnsignedInt()
	gbs.Prob = tok.CommaOptionalUnsignedFloat()
	gbs.Bias = tok.CommaOptionalFloat()
	gbs.StdDev = tok.CommaOptionalUnsignedFloat()
	gbs.SystemID = tok.CommaOptionalHex()
	gbs.SignalID = tok.CommaOptionalHex()
	return &gbs, tok.Err()
}
