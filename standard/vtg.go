package standard

import "github.com/twpayne/go-nmea"

type VTG struct {
	nmea.Address
	TrueCourseOverGround     float64
	MagneticCourseOverGround nmea.Optional[float64]
	SpeedOverGroundKN        float64
	SpeedOverGroundKPH       float64
	ModeIndicator            byte
}

func ParseVTG(addr string, tok *nmea.Tokenizer) (*VTG, error) {
	var vtg VTG
	vtg.Address = nmea.NewAddress(addr)
	vtg.TrueCourseOverGround = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('T')
	vtg.MagneticCourseOverGround = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('M')
	vtg.SpeedOverGroundKN = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('N')
	vtg.SpeedOverGroundKPH = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('K')
	vtg.ModeIndicator = tok.CommaOneByteOf("ADEFNR")
	tok.EndOfData()
	return &vtg, tok.Err()
}
