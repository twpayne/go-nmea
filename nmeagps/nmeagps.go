// Package nmeagps parses GPS NMEA sentences.
package nmeagps

import (
	"github.com/twpayne/go-nmea"
)

const (
	latHemis       = "NS"
	lonHemis       = "EW"
	modeIndicators = "AEMSV"
	posModes       = "ADEFNR"
	statuses       = "AV"
)

var (
	TalkerIDs = map[string]string{
		"GA": "Galileo",
		"GB": "BeiDou",
		"GL": "GLONASS",
		"GN": "Any combination of GNSS",
		"GP": "GPS, SBAS, QZSS",
	}

	sentenceParserMap = nmea.SentenceParserMap{
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

func SentenceParser(address string) nmea.SentenceParser {
	match := addressRx.FindStringSubmatch(address)
	if match == nil {
		return nil
	}
	return sentenceParserMap[match[1]]
}

func commaAltitudeCommaM(tok *nmea.Tokenizer) float64 {
	alt := tok.CommaFloat()
	tok.CommaLiteralByte('M')
	return alt
}

func commaLatCommaHemi(tok *nmea.Tokenizer) float64 {
	lat := tok.CommaUnsignedFloat()
	if tok.CommaOneByteOf(latHemis) == 'S' {
		lat = -lat
	}
	return lat
}

func ParseCommaLatDegMinCommaHemi(tok *nmea.Tokenizer) float64 {
	tok.Comma()
	return latDegMinCommaHemi(tok)
}

func commaOptionalLatDegMinCommaHemi(tok *nmea.Tokenizer) nmea.Optional[float64] {
	tok.Comma()
	if c, ok := tok.Peek(); !ok || c == ',' {
		tok.Comma()
		return nmea.Optional[float64]{}
	}
	return nmea.NewOptional(latDegMinCommaHemi(tok))
}

func commaLonCommaHemi(tok *nmea.Tokenizer) float64 {
	lon := tok.CommaUnsignedFloat()
	if tok.CommaOneByteOf(lonHemis) == 'W' {
		return -lon
	}
	return lon
}

func ParseCommaLonDegMinCommaHemi(tok *nmea.Tokenizer) float64 {
	tok.Comma()
	return lonDegMinCommaHemi(tok)
}

func commaOptionalLonDegMinCommaHemi(tok *nmea.Tokenizer) nmea.Optional[float64] {
	tok.Comma()
	if c, ok := tok.Peek(); !ok || c == ',' {
		tok.Comma()
		return nmea.Optional[float64]{}
	}
	return nmea.NewOptional(lonDegMinCommaHemi(tok))
}

func latDegMinCommaHemi(tok *nmea.Tokenizer) float64 {
	deg := tok.DecimalDigits(2)
	min := tok.UnsignedFloat()
	lat := float64(deg) + min/60
	if tok.CommaOneByteOf(latHemis) == 'S' {
		lat = -lat
	}
	return lat
}

func lonDegMinCommaHemi(tok *nmea.Tokenizer) float64 {
	deg := tok.DecimalDigits(3)
	min := tok.UnsignedFloat()
	lon := float64(deg) + min/60
	if tok.CommaOneByteOf(lonHemis) == 'W' {
		lon = -lon
	}
	return lon
}
