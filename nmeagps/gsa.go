package nmeagps

import "github.com/twpayne/go-nmea"

type GSA struct {
	address  Address
	OpMode   byte
	NavMode  int
	SVIDs    []nmea.Optional[int]
	PDOP     nmea.Optional[float64]
	HDOP     nmea.Optional[float64]
	VDOP     nmea.Optional[float64]
	SystemID nmea.Optional[int]
}

func ParseGSA(addr string, tok *nmea.Tokenizer) (*GSA, error) {
	var gsa GSA
	gsa.address = NewAddress(addr)
	gsa.OpMode = tok.CommaOneByteOf("AM")
	gsa.NavMode = tok.CommaUnsignedInt()
	for i := 0; i < 12; i++ {
		id := tok.CommaOptionalUnsignedInt()
		gsa.SVIDs = append(gsa.SVIDs, id)
	}
	gsa.PDOP = tok.CommaOptionalUnsignedFloat()
	gsa.HDOP = tok.CommaOptionalUnsignedFloat()
	gsa.VDOP = tok.CommaOptionalUnsignedFloat()
	if !tok.AtEndOfData() {
		gsa.SystemID = tok.CommaOptionalUnsignedInt()
	}
	tok.EndOfData()
	return &gsa, tok.Err()
}

func (gsa GSA) Address() nmea.Addresser {
	return gsa.address
}
