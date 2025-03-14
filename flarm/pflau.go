package flarm

import "github.com/twpayne/go-nmea"

// A PFLAU contains heart beat, status, and basic alarms.
type PFLAU struct {
	nmea.Address
	Rx               int
	Tx               int
	GPS              int
	Power            int
	AlarmLevel       int
	RelativeBearing  nmea.Optional[int]
	AlarmType        int
	RelativeVertical nmea.Optional[int]
	RelativeDistance nmea.Optional[int]
	ID               nmea.Optional[int]
}

func ParsePFLAU(addr string, tok *nmea.Tokenizer) (*PFLAU, error) {
	var pflau PFLAU
	pflau.Address = nmea.NewAddress(addr)
	pflau.Rx = tok.CommaUnsignedInt()
	pflau.Tx = tok.CommaUnsignedInt()
	pflau.GPS = tok.CommaUnsignedInt()
	pflau.Power = tok.CommaUnsignedInt()
	pflau.AlarmLevel = tok.CommaUnsignedInt()
	pflau.RelativeBearing = tok.CommaOptionalInt()
	pflau.AlarmType = tok.CommaUnsignedInt()
	pflau.RelativeVertical = tok.CommaOptionalInt()
	pflau.RelativeDistance = tok.CommaOptionalUnsignedInt()
	if !tok.AtEndOfData() {
		pflau.ID = tok.CommaOptionalHex()
	}
	tok.EndOfData()
	return &pflau, tok.Err()
}
