package flarm

import "github.com/twpayne/go-nmea"

// A PFLAI contains IGC readout or trigger or triggers an IGC pilot event.
type PFLAI struct {
	nmea.Address
	Value   string
	Result  string
	Message nmea.Optional[string]
}

func ParsePFLAI(addr string, tok *nmea.Tokenizer) (*PFLAI, error) {
	var pflai PFLAI
	pflai.Address = nmea.NewAddress(addr)
	pflai.Value = tok.CommaString()
	pflai.Result = tok.CommaString()
	if pflai.Result == "ERROR" && !tok.AtEndOfData() {
		pflai.Message = nmea.NewOptional(tok.CommaString())
	}
	tok.EndOfData()
	return &pflai, tok.Err()
}
