package flarm

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

type PFLAJAnswer struct {
	nmea.Address
	FlightState          int
	FlightRecorderState  int
	TISRADSRClientStatus nmea.Optional[int]
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
		panic(fmt.Errorf("%c: unknown query type", queryType))
	}
}

func ParsePFLAJAnswer(addr string, tok *nmea.Tokenizer) (*PFLAJAnswer, error) {
	var pflajAnswer PFLAJAnswer
	pflajAnswer.Address = nmea.NewAddress(addr)
	pflajAnswer.FlightState = tok.CommaUnsignedInt()
	pflajAnswer.FlightRecorderState = tok.CommaUnsignedInt()
	if !tok.AtEndOfData() {
		pflajAnswer.TISRADSRClientStatus = tok.CommaOptionalUnsignedInt()
	}
	tok.EndOfData()
	return &pflajAnswer, tok.Err()
}
