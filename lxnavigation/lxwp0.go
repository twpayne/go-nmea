package lxnavigation

import "github.com/twpayne/go-nmea"

type LXWP0 struct {
	nmea.Address
	IsLoggerRunning bool
	TrueAirspeedKPH nmea.Optional[float64]
	Alt             nmea.Optional[float64]
	Vario           []float64
	Heading         nmea.Optional[int]
	WindDirection   nmea.Optional[float64]
	WindSpeedKMH    nmea.Optional[float64]
}

func ParseLXWP0(addr string, tok *nmea.Tokenizer) (*LXWP0, error) {
	var lxwp0 LXWP0
	lxwp0.Address = nmea.NewAddress(addr)
	if tok.CommaOneByteOf("YN") == 'Y' {
		lxwp0.IsLoggerRunning = true
	}
	lxwp0.TrueAirspeedKPH = tok.CommaOptionalUnsignedFloat()
	lxwp0.Alt = tok.CommaOptionalFloat()
	lxwp0.Vario = make([]float64, 0, 6)
	for i := 0; i < 6; i++ {
		if vario := tok.CommaOptionalFloat(); vario.Valid {
			lxwp0.Vario = append(lxwp0.Vario, vario.Value)
		}
	}
	lxwp0.Heading = tok.CommaOptionalInt()
	lxwp0.WindDirection = tok.CommaOptionalUnsignedFloat()
	lxwp0.WindSpeedKMH = tok.CommaOptionalUnsignedFloat()
	tok.EndOfData()
	return &lxwp0, tok.Err()
}
