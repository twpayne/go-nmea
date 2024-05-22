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
	var pgrmh PGRMH
	pgrmh.Address = nmea.NewAddress(addr)
	pgrmh.DataStatus = tok.CommaOneByteOf("Av")
	pgrmh.VerticalSpeedFeetPerMinute = tok.CommaInt()
	pgrmh.VNAVProfileErrorFeet = tok.CommaInt()
	pgrmh.VerticalSpeedToVNAVTargetFeetPerMinute = tok.CommaInt()
	pgrmh.VerticalSpeedToNextWaypointFeetPerMinute = tok.CommaInt()
	pgrmh.ApproximateHeightAboveTerrainFeet = tok.CommaInt()
	pgrmh.DesiredTrack = tok.CommaInt()
	pgrmh.CourseOfNextRouteLeg = tok.CommaInt()
	tok.EndOfData()
	return &pgrmh, tok.Err()
}
