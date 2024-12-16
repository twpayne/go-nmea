package ublox

import "github.com/twpayne/go-nmea"

type SatelliteStatus struct {
	SVID   int
	Status byte
	Az     nmea.Optional[int]
	El     nmea.Optional[int]
	CNO    int
	Lck    int
}

type Status struct {
	nmea.Address
	N                 int
	SatelliteStatuses []SatelliteStatus
}

func ParseStatus(addr string, tok *nmea.Tokenizer) (*Status, error) {
	var s Status
	s.Address = nmea.NewAddress(addr)
	s.N = tok.CommaUnsignedInt()
	s.SatelliteStatuses = make([]SatelliteStatus, s.N)
	for i := range s.N {
		var ss SatelliteStatus
		ss.SVID = tok.CommaUnsignedInt()
		ss.Status = tok.CommaOneByteOf("-Ue")
		ss.Az = tok.CommaOptionalUnsignedInt()
		ss.El = tok.CommaOptionalUnsignedInt()
		ss.CNO = tok.CommaUnsignedInt()
		ss.Lck = tok.CommaUnsignedInt()
		s.SatelliteStatuses[i] = ss
	}
	tok.EndOfData()
	return &s, tok.Err()
}
