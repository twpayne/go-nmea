// Package standard parses standard NMEA sentences.
package standard

import (
	"regexp"

	"github.com/twpayne/go-nmea"
)

var (
	TalkerNames = map[string]string{
		"GA": "Galileo",
		"GB": "BeiDou",
		"GL": "GLONASS",
		"GN": "Any combination of GNSS",
		"GP": "GPS, SBAS, QZSS",
	}

	addressRx = regexp.MustCompile(`\A[A-Z]{2}([A-Z]{3})\z`)

	sentenceParserMap = nmea.SentenceParserMap{
		"ALM": nmea.MakeSentenceParser(ParseALM),
		"DBT": nmea.MakeSentenceParser(ParseDBT),
		"DBS": nmea.MakeSentenceParser(ParseDBS),
		"DPT": nmea.MakeSentenceParser(ParseDPT),
		"DTM": nmea.MakeSentenceParser(ParseDTM),
		"EVT": nmea.MakeSentenceParser(ParseEVT),
		"GBS": nmea.MakeSentenceParser(ParseGBS),
		"GGA": nmea.MakeSentenceParser(ParseGGA),
		"GLL": nmea.MakeSentenceParser(ParseGLL),
		"GNS": nmea.MakeSentenceParser(ParseGNS),
		"GRS": nmea.MakeSentenceParser(ParseGRS),
		"GSA": nmea.MakeSentenceParser(ParseGSA),
		"GST": nmea.MakeSentenceParser(ParseGST),
		"GSV": nmea.MakeSentenceParser(ParseGSV),
		"HDT": nmea.MakeSentenceParser(ParseHDT),
		"MLA": nmea.MakeSentenceParser(ParseMLA),
		"MSS": nmea.MakeSentenceParser(ParseMSS),
		"MTW": nmea.MakeSentenceParser(ParseMTW),
		"RMB": nmea.MakeSentenceParser(ParseRMB),
		"RMC": nmea.MakeSentenceParser(ParseRMC),
		"THS": nmea.MakeSentenceParser(ParseTHS),
		"TXT": nmea.MakeSentenceParser(ParseTXT),
		"VHW": nmea.MakeSentenceParser(ParseVHW),
		"VLW": nmea.MakeSentenceParser(ParseVLW),
		"VTG": nmea.MakeSentenceParser(ParseVTG),
		"ZDA": nmea.MakeSentenceParser(ParseZDA),
	}
)

func SentenceParserFunc(addr string) nmea.SentenceParser {
	match := addressRx.FindStringSubmatch(addr)
	if match != nil {
		if sentenceParser := sentenceParserMap[match[1]]; sentenceParser != nil {
			return sentenceParser
		}
	}
	return nil
}
