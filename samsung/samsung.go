// Package samsung parses Samsung NMEA sentences.
//
// There is no public documentation for these sentences. The parsers are
// incomplete.
package samsung

import "github.com/twpayne/go-nmea"

var sentenceParserMap = nmea.SentenceParserMap{
	"PSAMCLK":  nmea.MakeSentenceParser(ParsePSAMCLK),
	"PSAMDLOK": nmea.MakeSentenceParser(ParsePSAMDLOK),
	"PSAMID":   nmea.MakeSentenceParser(ParsePSAMID),
	"PSAMSA":   nmea.MakeSentenceParser(ParsePSAMSA),
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
