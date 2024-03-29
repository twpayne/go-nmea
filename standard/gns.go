package standard

import "github.com/twpayne/go-nmea"

type GNS struct {
	nmea.Address
	nmea.TimeOfDay
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
	gns.Address = nmea.NewAddress(addr)
	gns.TimeOfDay = nmea.ParseCommaTimeOfDay(tok)
	gns.Lat = tok.CommaOptionalLatDegMinCommaHemi()
	gns.Lon = tok.CommaOptionalLonDegMinCommaHemi()
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
