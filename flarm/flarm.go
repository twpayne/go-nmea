// Package flarm parses FLARM NMEA sentences.
//
// See https://www.flarm.com/wp-content/uploads/man/FTD-012-Data-Port-Interface-Control-Document-ICD.pdf.
package flarm

import "github.com/twpayne/go-nmea"

var sentenceParserMap = nmea.SentenceParserMap{
	"PFLAA": nmea.MakeSentenceParser(ParsePFLAA),
	"PFLAE": ParsePFLAE,
	"PFLAC": ParsePFLAC,
	"PFLAI": nmea.MakeSentenceParser(ParsePFLAI),
	"PFLAJ": ParsePFLAJ,
	"PFLAL": nmea.MakeSentenceParser(ParsePFLAL),
	"PFLAN": ParsePFLAN,
	"PFLAO": nmea.MakeSentenceParser(ParsePFLAO),
	"PFLAQ": nmea.MakeSentenceParser(ParsePFLAQ),
	"PFLAU": nmea.MakeSentenceParser(ParsePFLAU),
	"PFLAV": ParsePFLAV,
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
