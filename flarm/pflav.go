package flarm

import (
	"github.com/twpayne/go-nmea"
)

type PFLAVAnswer struct {
	nmea.Address
	HardwareVersion string
	SoftwareVersion string
	ObstacleVersion nmea.Optional[string]
}

func ParsePFLAV(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		return ParsePFLAVAnswer(addr, tok)
	default:
		return nil, &UnknownQueryTypeError{
			QueryType: queryType,
		}
	}
}

func ParsePFLAVAnswer(addr string, tok *nmea.Tokenizer) (*PFLAVAnswer, error) {
	var pflavAnswer PFLAVAnswer
	pflavAnswer.Address = nmea.NewAddress(addr)
	pflavAnswer.HardwareVersion = tok.CommaString()
	pflavAnswer.SoftwareVersion = tok.CommaString()
	pflavAnswer.ObstacleVersion = tok.CommaOptionalString()
	return &pflavAnswer, tok.Err()
}
