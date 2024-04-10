package standard

import "github.com/twpayne/go-nmea"

type ALM struct {
	nmea.Address
	NumMsg               int
	MsgNum               int
	PRN                  int
	GPSWeek              int
	SVHealth             int
	Eccentricity         int
	AlmanacReferenceTime int
	InclinationAngle     int
	OmegaDot             int
	RootAxis             int
	Omega                int
	AscensionNodeLon     int
	MeanAnomaly          int
	AF0                  int
	AF1                  int
}

func ParseALM(addr string, tok *nmea.Tokenizer) (*ALM, error) {
	var alm ALM
	alm.Address = nmea.NewAddress(addr)
	alm.NumMsg = tok.CommaUnsignedInt()
	alm.MsgNum = tok.CommaUnsignedInt()
	alm.PRN = tok.CommaUnsignedInt()
	alm.GPSWeek = tok.CommaUnsignedInt()
	alm.SVHealth = tok.CommaHex()
	alm.Eccentricity = tok.CommaHex()
	alm.AlmanacReferenceTime = tok.CommaHex()
	alm.InclinationAngle = tok.CommaHex()
	alm.OmegaDot = tok.CommaHex()
	alm.RootAxis = tok.CommaHex()
	alm.Omega = tok.CommaHex()
	alm.AscensionNodeLon = tok.CommaHex()
	alm.MeanAnomaly = tok.CommaHex()
	alm.AF0 = tok.CommaHex()
	alm.AF1 = tok.CommaHex()
	tok.EndOfData()
	return &alm, tok.Err()
}
