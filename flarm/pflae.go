package flarm

import "github.com/twpayne/go-nmea"

type PFLAEAnswer struct {
	nmea.Address
	Severity  nmea.Optional[int]
	ErrorCode nmea.Optional[int]
	Message   nmea.Optional[string]
}

func ParsePFLAE(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		return ParsePFLAEAnswer(addr, tok)
	default:
		return nil, &UnknownQueryTypeError{
			QueryType: queryType,
		}
	}
}

func ParsePFLAEAnswer(addr string, tok *nmea.Tokenizer) (*PFLAEAnswer, error) {
	var pflaeAnswer PFLAEAnswer
	pflaeAnswer.Address = nmea.NewAddress(addr)
	if !tok.AtEndOfData() {
		pflaeAnswer.Severity = nmea.NewOptional(tok.CommaUnsignedInt())
	}
	if !tok.AtEndOfData() {
		pflaeAnswer.ErrorCode = nmea.NewOptional(tok.CommaHex())
	}
	if !tok.AtEndOfData() {
		pflaeAnswer.Message = tok.CommaOptionalString()
	}
	tok.EndOfData()
	return &pflaeAnswer, tok.Err()
}
