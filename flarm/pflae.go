package flarm

import "github.com/twpayne/go-nmea"

type PFLAEAnswer struct {
	nmea.Address
	Severity  int
	ErrorCode int
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
		return nil, tok.Err()
	}
}

func ParsePFLAEAnswer(addr string, tok *nmea.Tokenizer) (*PFLAEAnswer, error) {
	var pflaeAnswer PFLAEAnswer
	pflaeAnswer.Address = nmea.NewAddress(addr)
	pflaeAnswer.Severity = tok.CommaUnsignedInt()
	pflaeAnswer.ErrorCode = tok.CommaHex()
	if !tok.AtEndOfData() {
		pflaeAnswer.Message = tok.CommaOptionalString()
	}
	tok.EndOfData()
	return &pflaeAnswer, tok.Err()
}
