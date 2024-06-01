// Package lk8000 parses LK8000 sentences.
//
// See https://github.com/LK8000/LK8000/blob/master/Docs/LK8EX1.txt.
package lk8000

import "github.com/twpayne/go-nmea"

var sentenceParserMap = nmea.SentenceParserMap{
	"LK8EX1": nmea.MakeSentenceParser(ParseLK8EX1),
}

func SentenceParserFunc(addr string) nmea.SentenceParser {
	return sentenceParserMap[addr]
}
