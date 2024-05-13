package standard

import "github.com/twpayne/go-nmea"

type DBS struct {
	nmea.Address
	DepthFeet    nmea.Optional[float64]
	Depth        nmea.Optional[float64]
	DepthFathoms nmea.Optional[float64]
}

func ParseDBS(addr string, tok *nmea.Tokenizer) (*DBS, error) {
	var dbt DBS
	dbt.Address = nmea.NewAddress(addr)
	dbt.DepthFeet = tok.CommaOptionalFloatCommaUnit('f')
	dbt.Depth = tok.CommaOptionalFloatCommaUnit('M')
	dbt.DepthFathoms = tok.CommaOptionalFloatCommaUnit('F')
	tok.EndOfData()
	return &dbt, tok.Err()
}
