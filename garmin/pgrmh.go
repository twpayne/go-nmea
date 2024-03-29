package garmin

import "github.com/twpayne/go-nmea"

type PGRMH struct {
	nmea.Address
	DataStatus                               byte
	VerticalSpeedFeetPerMinute               int
	VNAVProfileErrorFeet                     int
	VerticalSpeedToVNAVTargetFeetPerMinute   int
	VerticalSpeedToNextWaypointFeetPerMinute int
	ApproximateHeightAboveTerrainFeet        int
	DesiredTrack                             int
	CourseOfNextRouteLeg                     int
}

func ParsePGRMH(addr string, tok *nmea.Tokenizer) (*PGRMH, error) {
	var h PGRMH
	h.Address = nmea.NewAddress(addr)
	h.DataStatus = tok.CommaOneByteOf("Av")
	h.VerticalSpeedFeetPerMinute = tok.CommaInt()
	h.VNAVProfileErrorFeet = tok.CommaInt()
	h.VerticalSpeedToVNAVTargetFeetPerMinute = tok.CommaInt()
	h.VerticalSpeedToNextWaypointFeetPerMinute = tok.CommaInt()
	h.ApproximateHeightAboveTerrainFeet = tok.CommaInt()
	h.DesiredTrack = tok.CommaInt()
	h.CourseOfNextRouteLeg = tok.CommaInt()
	tok.EndOfData()
	return &h, tok.Err()
}
