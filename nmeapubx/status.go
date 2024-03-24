package nmeapubx

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
	address           Address
	N                 int
	SatelliteStatuses []SatelliteStatus
}

func ParseStatus(addr string, tok *nmea.Tokenizer) (*Status, error) {
	var s Status
	s.address = NewAddress(addr)
	s.N = tok.CommaUnsignedInt()
	s.SatelliteStatuses = make([]SatelliteStatus, 0, s.N)
	for i := 0; i < s.N; i++ {
		var ss SatelliteStatus
		ss.SVID = tok.CommaUnsignedInt()
		ss.Status = tok.CommaOneByteOf("-Ue")
		ss.Az = tok.CommaOptionalUnsignedInt()
		ss.El = tok.CommaOptionalUnsignedInt()
		ss.CNO = tok.CommaUnsignedInt()
		ss.Lck = tok.CommaUnsignedInt()
		s.SatelliteStatuses = append(s.SatelliteStatuses, ss)
	}
	tok.EndOfData()
	return &s, tok.Err()
}

func (s Status) Address() nmea.Address {
	return s.address
}
