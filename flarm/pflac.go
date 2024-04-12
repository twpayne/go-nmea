package flarm

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

type PFLACAnswer struct {
	nmea.Address
	ConfigurationItem string
	Values            []string
}

type PFLACError struct {
	nmea.Address
}

func ParsePFLAC(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		switch s := tok.CommaString(); s {
		case "ERROR":
			var pflacError PFLACError
			pflacError.Address = nmea.NewAddress(addr)
			tok.EndOfData()
			return &pflacError, tok.Err()
		default:
			var pflacAnswer PFLACAnswer
			pflacAnswer.Address = nmea.NewAddress(addr)
			pflacAnswer.ConfigurationItem = s
			for !tok.AtEndOfData() {
				value := tok.CommaString()
				pflacAnswer.Values = append(pflacAnswer.Values, value)
			}
			tok.EndOfData()
			return &pflacAnswer, tok.Err()
		}
	default:
		panic(fmt.Errorf("%c: unknown query type", queryType))
	}
}
