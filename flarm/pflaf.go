package flarm

import "github.com/twpayne/go-nmea"

// A PFLAFAnswer contains information about simulated traffic and alarms.
type PFLAFAnswer struct {
	nmea.Address
	ScenarioNumber int
}

type PFLAFError struct {
	nmea.Address
	ErrorType string
}

func ParsePFLAF(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		tokFork := tok.Fork()
		if tok.CommaLiteralString("ERROR"); tok.Err() == nil {
			var pflafError PFLAFError
			pflafError.Address = nmea.NewAddress(addr)
			pflafError.ErrorType = tok.CommaString()
			tok.EndOfData()
			return &pflafError, tok.Err()
		}
		tok = tokFork
		var pflafAnswer PFLAFAnswer
		pflafAnswer.Address = nmea.NewAddress(addr)
		pflafAnswer.ScenarioNumber = tok.CommaUnsignedInt()
		tok.EndOfData()
		return &pflafAnswer, tok.Err()
	default:
		return nil, &UnknownQueryTypeError{
			QueryType: queryType,
		}
	}
}
