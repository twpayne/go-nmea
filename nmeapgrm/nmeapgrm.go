// Package nmeapgrm parses Garmin NMEA sentences.
//
// See https://developer.garmin.com/downloads/legacy/uploads/2015/08/190-00684-00.pdf.
package nmeapgrm

import (
	"github.com/twpayne/go-nmea"
)

var sentenceParserMap = nmea.SentenceParserMap{
	"PGRMB": nmea.MakeSentenceParser(ParsePGRMB),
	"PGRME": nmea.MakeSentenceParser(ParsePGRME),
	"PGRMF": nmea.MakeSentenceParser(ParsePGRMF),
	"PGRMH": nmea.MakeSentenceParser(ParsePGRMH),
	"PGRMM": nmea.MakeSentenceParser(ParsePGRMM),
	"PGRMT": nmea.MakeSentenceParser(ParsePGRMT),
	"PGRMV": nmea.MakeSentenceParser(ParsePGRMV),
	"PGRMZ": nmea.MakeSentenceParser(ParsePGRMZ),
}

func SentenceParser(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
