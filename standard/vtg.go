package standard

import "github.com/twpayne/go-nmea"

type VTG struct {
	nmea.Address
	TrueCourseOverGround     nmea.Optional[float64]
	MagneticCourseOverGround nmea.Optional[float64]
	SpeedOverGroundKN        nmea.Optional[float64]
	SpeedOverGroundKPH       nmea.Optional[float64]
	ModeIndicator            byte
}

func ParseVTG(addr string, tok *nmea.Tokenizer) (*VTG, error) {
	var vtg VTG
	vtg.Address = nmea.NewAddress(addr)
	vtg.TrueCourseOverGround = tok.CommaOptionalFloatCommaUnit('T')
	vtg.MagneticCourseOverGround = tok.CommaOptionalFloatCommaUnit('M')
	vtg.SpeedOverGroundKN = tok.CommaOptionalFloatCommaUnit('N')
	vtg.SpeedOverGroundKPH = tok.CommaOptionalFloatCommaUnit('K')
	vtg.ModeIndicator = tok.CommaOneByteOf("ADEFNR")
	tok.EndOfData()
	return &vtg, tok.Err()
}
