package gps

import "github.com/twpayne/go-nmea"

type SatelliteInView struct {
	SVID int
	Elv  nmea.Optional[int]
	Az   nmea.Optional[int]
	CNO  nmea.Optional[int]
}

type GSV struct {
	address          Address
	NumMsg           int
	MsgNum           int
	NumSV            int
	SatellitesInView []SatelliteInView
	SignalID         nmea.Optional[int]
}

func ParseGSV(addr string, tok *nmea.Tokenizer) (*GSV, error) {
	var gsv GSV
	gsv.address = NewAddress(addr)
	gsv.NumMsg = tok.CommaInt()
	gsv.MsgNum = tok.CommaInt()
	gsv.NumSV = tok.CommaInt()
	n := min(gsv.NumSV-4*(gsv.MsgNum-1), 4)
	for i := 0; i < n; i++ {
		siv := SatelliteInView{}
		siv.SVID = tok.CommaInt()
		siv.Elv = tok.CommaOptionalUnsignedInt()
		siv.Az = tok.CommaOptionalUnsignedInt()
		siv.CNO = tok.CommaOptionalUnsignedInt()
		gsv.SatellitesInView = append(gsv.SatellitesInView, siv)
	}
	if !tok.AtEndOfData() {
		gsv.SignalID = tok.CommaOptionalUnsignedInt()
	}
	tok.EndOfData()
	return &gsv, tok.Err()
}

func (gsv GSV) Address() nmea.Addresser {
	return gsv.address
}
