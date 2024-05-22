package ublox

import "github.com/twpayne/go-nmea"

type Position struct {
	nmea.Address
	MsgID              int
	TimeOfDay          nmea.TimeOfDay
	Lat                float64
	Lon                float64
	AltRef             float64
	NavStat            string
	HorizAcc           float64
	VertAcc            float64
	SpeedOverGroundKPH float64
	CourseOverGround   float64
	VertVel            float64
	DiffAge            nmea.Optional[float64]
	HDOP               float64
	VDOP               float64
	TDOP               float64
	NumSVs             int
	Reserved           struct{}
	DR                 int
}

func ParsePosition(addr string, tok *nmea.Tokenizer) (*Position, error) {
	var p Position
	p.Address = nmea.NewAddress(addr)
	p.TimeOfDay = tok.CommaTimeOfDay()
	p.Lat = tok.CommaLatDegMinCommaHemi()
	p.Lon = tok.CommaLonDegMinCommaHemi()
	p.AltRef = tok.CommaFloat()
	p.NavStat = tok.CommaString()
	p.HorizAcc = tok.CommaUnsignedFloat()
	p.VertAcc = tok.CommaUnsignedFloat()
	p.SpeedOverGroundKPH = tok.CommaUnsignedFloat()
	p.CourseOverGround = tok.CommaUnsignedFloat()
	p.VertVel = tok.CommaFloat()
	p.DiffAge = tok.CommaOptionalUnsignedFloat()
	p.HDOP = tok.CommaUnsignedFloat()
	p.VDOP = tok.CommaUnsignedFloat()
	p.TDOP = tok.CommaUnsignedFloat()
	p.NumSVs = tok.CommaUnsignedInt()
	p.Reserved = tok.CommaIgnore()
	p.DR = tok.CommaInt()
	tok.EndOfData()
	return &p, tok.Err()
}
