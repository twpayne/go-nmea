package standard

import "github.com/twpayne/go-nmea"

type DBT struct {
	nmea.Address
	DepthFeet    float64
	Depth        float64
	DepthFathoms float64
}

func ParseDBT(addr string, tok *nmea.Tokenizer) (*DBT, error) {
	var dbt DBT
	dbt.Address = nmea.NewAddress(addr)
	dbt.DepthFeet = tok.CommaFloatCommaUnit('f')
	dbt.Depth = tok.CommaFloatCommaUnit('M')
	dbt.DepthFathoms = tok.CommaFloatCommaUnit('F')
	tok.EndOfData()
	return &dbt, tok.Err()
}
