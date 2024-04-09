package standard

import "github.com/twpayne/go-nmea"

type MLA struct {
	nmea.Address
	NumMsg                 int
	MsgNum                 int
	PRN                    int
	GPSWeek                int
	SVHealth               int
	Eccentricity           int
	AlmanacReferenceTime   int
	InclinationAngle       int
	OmegaDot               int
	RootAxis               int
	Omega                  int
	AscensionNodeLongitude int
	MeanAnomaly            int
	AF0                    int
	AF1                    int
}

func ParseMLA(addr string, tok *nmea.Tokenizer) (*MLA, error) {
	var mla MLA
	mla.Address = nmea.NewAddress(addr)
	mla.NumMsg = tok.CommaUnsignedInt()
	mla.MsgNum = tok.CommaUnsignedInt()
	mla.PRN = tok.CommaUnsignedInt()
	mla.GPSWeek = tok.CommaUnsignedInt()
	mla.SVHealth = tok.CommaHex()
	mla.Eccentricity = tok.CommaHex()
	mla.AlmanacReferenceTime = tok.CommaHex()
	mla.InclinationAngle = tok.CommaHex()
	mla.OmegaDot = tok.CommaHex()
	mla.RootAxis = tok.CommaHex()
	mla.Omega = tok.CommaHex()
	mla.AscensionNodeLongitude = tok.CommaHex()
	mla.MeanAnomaly = tok.CommaHex()
	mla.AF0 = tok.CommaHex()
	mla.AF1 = tok.CommaHex()
	tok.EndOfData()
	return &mla, tok.Err()
}
