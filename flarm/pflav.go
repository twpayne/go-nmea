package flarm

import (
	"regexp"

	"github.com/twpayne/go-nmea"
)

var (
	hardwareVersionRx = regexp.MustCompile(`\A\d\.\d\d`)
	softwareVersionRx = regexp.MustCompile(`\A\d\d?\.\d\d`)
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
		return nil, tok.Err()
	}
}

func ParsePFLAVAnswer(addr string, tok *nmea.Tokenizer) (*PFLAVAnswer, error) {
	var pflavAnswer PFLAVAnswer
	pflavAnswer.Address = nmea.NewAddress(addr)
	if match := tok.CommaRegexp(hardwareVersionRx); match != nil {
		pflavAnswer.HardwareVersion = string(match[0])
	}
	if match := tok.CommaRegexp(softwareVersionRx); match != nil {
		pflavAnswer.SoftwareVersion = string(match[0])
	}
	pflavAnswer.ObstacleVersion = tok.CommaOptionalString()
	return &pflavAnswer, tok.Err()
}
