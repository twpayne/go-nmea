package standard

import "github.com/twpayne/go-nmea"

type GST struct {
	nmea.Address
	TimeOfDay   nmea.TimeOfDay
	RangeRMS    float64
	MajorStdDev nmea.Optional[float64]
	MinorStdDev nmea.Optional[float64]
	Orientation nmea.Optional[float64]
	LatStdDev   float64
	LonStdDev   float64
	AltStdDev   float64
}

func ParseGST(addr string, tok *nmea.Tokenizer) (*GST, error) {
	var gst GST
	gst.Address = nmea.NewAddress(addr)
	gst.TimeOfDay = tok.CommaTimeOfDay()
	gst.RangeRMS = tok.CommaUnsignedFloat()
	gst.MajorStdDev = tok.CommaOptionalUnsignedFloat()
	gst.MinorStdDev = tok.CommaOptionalUnsignedFloat()
	gst.Orientation = tok.CommaOptionalFloat()
	gst.LatStdDev = tok.CommaUnsignedFloat()
	gst.LonStdDev = tok.CommaUnsignedFloat()
	gst.AltStdDev = tok.CommaUnsignedFloat()
	return &gst, tok.Err()
}
