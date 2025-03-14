package flarm

import (
	"github.com/twpayne/go-nmea"
)

// A PFLAJAnswer contains flight and IGC recording state information.
type PFLAJAnswer struct {
	nmea.Address
	FlightState          int
	FlightRecorderState  int
	TISBADSRClientStatus nmea.Optional[int]
}

func ParsePFLAJ(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		return ParsePFLAJAnswer(addr, tok)
	default:
		return nil, &UnknownQueryTypeError{
			QueryType: queryType,
		}
	}
}

func ParsePFLAJAnswer(addr string, tok *nmea.Tokenizer) (*PFLAJAnswer, error) {
	var pflajAnswer PFLAJAnswer
	pflajAnswer.Address = nmea.NewAddress(addr)
	pflajAnswer.FlightState = tok.CommaUnsignedInt()
	pflajAnswer.FlightRecorderState = tok.CommaUnsignedInt()
	if !tok.AtEndOfData() {
		pflajAnswer.TISBADSRClientStatus = tok.CommaOptionalUnsignedInt()
	}
	tok.EndOfData()
	return &pflajAnswer, tok.Err()
}
