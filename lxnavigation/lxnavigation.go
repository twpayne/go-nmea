// Package lxnavigation parses LX Navigation NMEA sentences.
//
// See https://downloads.lxnavigation.com/downloads/manuals/LX_CP_R5.pdf.
package lxnavigation

import "github.com/twpayne/go-nmea"

var sentenceParserMap = nmea.SentenceParserMap{
	"LXWP0": nmea.MakeSentenceParser(ParseLXWP0),
	"LXWP1": nmea.MakeSentenceParser(ParseLXWP1),
	"LXWP2": nmea.MakeSentenceParser(ParseLXWP2),
	"LXWP3": nmea.MakeSentenceParser(ParseLXWP3),
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
