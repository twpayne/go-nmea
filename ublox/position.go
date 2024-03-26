package ublox

import (
	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/gps"
)

type Position struct {
	address            Address
	MsgID              int
	TimeOfDay          gps.TimeOfDay
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
	Reserved           int
	DR                 int
}

func ParsePosition(addr string, tok *nmea.Tokenizer) (*Position, error) {
	var p Position
	p.address = NewAddress(addr)
	p.TimeOfDay = gps.ParseCommaTimeOfDay(tok)
	p.Lat = gps.ParseCommaLatDegMinCommaHemi(tok)
	p.Lon = gps.ParseCommaLonDegMinCommaHemi(tok)
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
	p.Reserved = tok.CommaInt()
	p.DR = tok.CommaInt()
	tok.EndOfData()
	return &p, tok.Err()
}

func (p Position) Address() nmea.Addresser {
	return p.address
}
