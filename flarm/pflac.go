package flarm

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

type PFLACAnswer struct {
	nmea.Address
	ConfigurationItem string
	Value             string
}

type PFLACError struct {
	nmea.Address
}

func ParsePFLAC(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	result := tok.CommaString()
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch result {
	case "OK":
		return ParsePFLACAnswer(addr, tok)
	case "ERROR":
		return ParsePFLACError(addr, tok)
	default:
		panic(fmt.Errorf("%c: unknown query type", queryType))
	}
}

func ParsePFLACAnswer(addr string, tok *nmea.Tokenizer) (*PFLACAnswer, error) {
	var pflacAnswer PFLACAnswer
	pflacAnswer.Address = nmea.NewAddress(addr)
	pflacAnswer.ConfigurationItem = tok.CommaString()
	pflacAnswer.Value = tok.CommaString()
	tok.EndOfData()
	return &pflacAnswer, tok.Err()
}

func ParsePFLACError(addr string, tok *nmea.Tokenizer) (*PFLACError, error) {
	var pflacError PFLACError
	pflacError.Address = nmea.NewAddress(addr)
	tok.EndOfData()
	return &pflacError, tok.Err()
}
