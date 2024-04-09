package standard

import "github.com/twpayne/go-nmea"

type RMB struct {
	nmea.Address
	DataStatus             byte
	CrossTrackError        float64
	CrossTrackErrorDir     byte
	OriginWaypointID       string
	DestinationWaypointID  string
	DestinationWaypointLat float64
	DestinationWaypointLon float64
	RangeNM                float64
	BearingTrue            float64
	ClosingVelocityKN      float64
	ArrivalStatus          byte
	ModeIndicator          byte
}

func ParseRMB(addr string, tok *nmea.Tokenizer) (*RMB, error) {
	var rmb RMB
	rmb.Address = nmea.NewAddress(addr)
	rmb.DataStatus = tok.CommaOneByteOf("AV")
	rmb.CrossTrackError = tok.CommaUnsignedFloat()
	rmb.CrossTrackErrorDir = tok.CommaOneByteOf("LR")
	rmb.OriginWaypointID = tok.CommaString()
	rmb.DestinationWaypointID = tok.CommaString()
	rmb.DestinationWaypointLat = tok.CommaLatDegMinCommaHemi()
	rmb.DestinationWaypointLon = tok.CommaLonDegMinCommaHemi()
	rmb.RangeNM = tok.CommaUnsignedFloat()
	rmb.BearingTrue = tok.CommaFloat()
	rmb.ClosingVelocityKN = tok.CommaFloat()
	rmb.ArrivalStatus = tok.CommaOneByteOf("AV")
	rmb.ModeIndicator = tok.CommaOneByteOf("ADEMN")
	tok.EndOfData()
	return &rmb, tok.Err()
}
