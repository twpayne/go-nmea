package nmeagps

import "github.com/twpayne/go-nmea"

type VTG struct {
	address                  Address
	TrueCourseOverGround     float64
	MagneticCourseOverGround nmea.Optional[float64]
	SpeedOverGroundKN        float64
	SpeedOverGroundKPH       float64
	ModeIndicator            byte
}

func ParseVTG(addr string, tok *nmea.Tokenizer) (*VTG, error) {
	var vtg VTG
	vtg.address = Address(addr)
	vtg.TrueCourseOverGround = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('T')
	vtg.MagneticCourseOverGround = tok.CommaOptionalUnsignedFloat()
	tok.CommaLiteralByte('M')
	vtg.SpeedOverGroundKN = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('N')
	vtg.SpeedOverGroundKPH = tok.CommaUnsignedFloat()
	tok.CommaLiteralByte('K')
	vtg.ModeIndicator = tok.CommaOneByteOf(posModes)
	tok.EndOfData()
	return &vtg, tok.Err()
}

func (vtg VTG) Address() nmea.Address {
	return vtg.address
}
