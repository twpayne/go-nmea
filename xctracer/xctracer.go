// Package xctracer parses XC Tracer sentences.
package xctracer

import "github.com/twpayne/go-nmea"

var sentenceParserMap = nmea.SentenceParserMap{
	"XCTRC": nmea.MakeSentenceParser(ParseXCTRC),
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
