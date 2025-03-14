// Package flarm parses FLARM NMEA sentences.
//
// See https://www.flarm.com/wp-content/uploads/2024/04/FTD-012-Data-Port-Interface-Control-Document-ICD-7.19.pdf.
// See https://www.flarm.com/wp-content/uploads/2024/07/FTD-014-FLARM-Configuration-Specification-1.17.pdf.
package flarm

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

type UnknownQueryTypeError struct {
	QueryType byte
}

func (e *UnknownQueryTypeError) Error() string {
	return fmt.Sprintf("%c: unknown query type", e.QueryType)
}

var sentenceParserMap = nmea.SentenceParserMap{
	"PFLAA": nmea.MakeSentenceParser(ParsePFLAA),
	"PFLAE": ParsePFLAE,
	"PFLAF": ParsePFLAF,
	"PFLAC": ParsePFLAC,
	"PFLAI": nmea.MakeSentenceParser(ParsePFLAI),
	"PFLAJ": ParsePFLAJ,
	"PFLAL": nmea.MakeSentenceParser(ParsePFLAL),
	"PFLAM": nmea.MakeSentenceParser(ParsePFLAM),
	"PFLAN": ParsePFLAN,
	"PFLAO": nmea.MakeSentenceParser(ParsePFLAO),
	"PFLAQ": nmea.MakeSentenceParser(ParsePFLAQ),
	"PFLAU": nmea.MakeSentenceParser(ParsePFLAU),
	"PFLAV": ParsePFLAV,
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
