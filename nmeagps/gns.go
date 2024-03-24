package nmeagps

import "github.com/twpayne/go-nmea"

type GNS struct {
	address Address
	TimeOfDay
	Lat         nmea.Optional[float64]
	Lon         nmea.Optional[float64]
	PosMode     []byte
	NumSV       int
	HDOP        nmea.Optional[float64]
	Alt         nmea.Optional[float64]
	Sep         nmea.Optional[float64]
	DiffAge     nmea.Optional[float64]
	DiffStation nmea.Optional[int]
	NavStatus   byte
}

func ParseGNS(addr string, tok *nmea.Tokenizer) (*GNS, error) {
	var gns GNS
	gns.address = NewAddress(addr)
	gns.TimeOfDay = ParseCommaTimeOfDay(tok)
	gns.Lat = commaOptionalLatDegMinCommaHemi(tok)
	gns.Lon = commaOptionalLonDegMinCommaHemi(tok)
	gns.PosMode = []byte(tok.CommaString())
	gns.NumSV = tok.CommaUnsignedInt()
	gns.HDOP = tok.CommaOptionalUnsignedFloat()
	gns.Alt = tok.CommaOptionalFloat()
	gns.Sep = tok.CommaOptionalFloat()
	gns.DiffAge = tok.CommaOptionalUnsignedFloat()
	gns.DiffStation = tok.CommaOptionalUnsignedInt()
	gns.NavStatus = tok.CommaOneByteOf("V")
	tok.EndOfData()
	return &gns, tok.Err()
}

func (gns GNS) Address() nmea.Addresser {
	return gns.address
}
