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

	gpsAddressRx = regexp.MustCompile(`\AG[A-Z]([A-Z]{3})\z`)

	gpsSentenceParserMap = nmea.SentenceParserMap{
		"DTM": nmea.MakeSentenceParser(ParseDTM),
		"GBS": nmea.MakeSentenceParser(ParseGBS),
		"GGA": nmea.MakeSentenceParser(ParseGGA),
		"GLL": nmea.MakeSentenceParser(ParseGLL),
		"GNS": nmea.MakeSentenceParser(ParseGNS),
		"GRS": nmea.MakeSentenceParser(ParseGRS),
		"GSA": nmea.MakeSentenceParser(ParseGSA),
		"GST": nmea.MakeSentenceParser(ParseGST),
		"GSV": nmea.MakeSentenceParser(ParseGSV),
		"MSS": nmea.MakeSentenceParser(ParseMSS),
		"RMC": nmea.MakeSentenceParser(ParseRMC),
		"THS": nmea.MakeSentenceParser(ParseTHS),
		"TXT": nmea.MakeSentenceParser(ParseTXT),
		"VLW": nmea.MakeSentenceParser(ParseVLW),
		"VTG": nmea.MakeSentenceParser(ParseVTG),
		"ZDA": nmea.MakeSentenceParser(ParseZDA),
	}
)

func SentenceParserFunc(addr string) nmea.SentenceParser {
	match := gpsAddressRx.FindStringSubmatch(addr)
	if match != nil {
		if sentenceParser := gpsSentenceParserMap[match[1]]; sentenceParser != nil {
			return sentenceParser
		}
	}
	return nil
}
