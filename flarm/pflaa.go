package flarm

import "github.com/twpayne/go-nmea"

type PFLAA struct {
	nmea.Address
	AlarmLevel       int
	RelativeNorth    int
	RelativeEast     nmea.Optional[int]
	RelativeVertical int
	IDType           nmea.Optional[int]
	ID               nmea.Optional[int]
	Track            nmea.Optional[int]
	TurnRate         struct{}
	GroundSpeed      nmea.Optional[int]
	ClimbRate        nmea.Optional[float64]
	AircraftType     int
	NoTrack          nmea.Optional[int]
	Source           nmea.Optional[int]
	RSSI             nmea.Optional[float64]
}

func ParsePFLAA(addr string, tok *nmea.Tokenizer) (*PFLAA, error) {
	var pflaa PFLAA
	pflaa.Address = nmea.NewAddress(addr)
	pflaa.AlarmLevel = tok.CommaUnsignedInt()
	pflaa.RelativeNorth = tok.CommaInt()
	pflaa.RelativeEast = tok.CommaOptionalInt()
	pflaa.RelativeVertical = tok.CommaInt()
	pflaa.IDType = tok.CommaOptionalUnsignedInt()
	pflaa.ID = tok.CommaOptionalHex()
	pflaa.Track = tok.CommaOptionalUnsignedInt()
	pflaa.TurnRate = tok.CommaEmpty()
	pflaa.GroundSpeed = tok.CommaOptionalUnsignedInt()
	pflaa.ClimbRate = tok.CommaOptionalFloat()
	pflaa.AircraftType = tok.CommaHex()
	if !tok.AtEndOfData() {
		pflaa.NoTrack = nmea.NewOptional(tok.CommaUnsignedInt())
	}
	if !tok.AtEndOfData() {
		pflaa.Source = nmea.NewOptional(tok.CommaUnsignedInt())
		pflaa.RSSI = tok.CommaOptionalFloat()
	}
	tok.EndOfData()
	return &pflaa, tok.Err()
}
