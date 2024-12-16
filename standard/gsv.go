package standard

import "github.com/twpayne/go-nmea"

type SatelliteInView struct {
	SVID int
	Elv  nmea.Optional[int]
	Az   nmea.Optional[int]
	CNO  nmea.Optional[int]
}

type GSV struct {
	nmea.Address
	NumMsg           int
	MsgNum           int
	NumSV            int
	SatellitesInView []SatelliteInView
	SignalID         nmea.Optional[int]
}

func ParseGSV(addr string, tok *nmea.Tokenizer) (*GSV, error) {
	var gsv GSV
	gsv.Address = nmea.NewAddress(addr)
	gsv.NumMsg = tok.CommaUnsignedInt()
	gsv.MsgNum = tok.CommaUnsignedInt()
	gsv.NumSV = tok.CommaUnsignedInt()
	if n := min(gsv.NumSV-4*(gsv.MsgNum-1), 4); n > 0 {
		for range n {
			siv := SatelliteInView{}
			siv.SVID = tok.CommaUnsignedInt()
			siv.Elv = tok.CommaOptionalInt()
			siv.Az = tok.CommaOptionalInt()
			siv.CNO = tok.CommaOptionalInt()
			gsv.SatellitesInView = append(gsv.SatellitesInView, siv)
		}
	} else {
		tokFork := tok.Fork()
		tok.Comma()
		tok.Comma()
		tok.Comma()
		tok.Comma()
		if tok.Err() != nil {
			tok = tokFork
		}
	}
	if !tok.AtEndOfData() {
		gsv.SignalID = tok.CommaOptionalUnsignedInt()
	}
	tok.EndOfData()
	return &gsv, tok.Err()
}
