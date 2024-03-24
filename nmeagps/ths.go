package nmeagps

import "github.com/twpayne/go-nmea"

type THS struct {
	address       Address
	TrueHeading   float64
	ModeIndicator byte
}

func ParseTHS(addr string, tok *nmea.Tokenizer) (*THS, error) {
	var ths THS
	ths.address = NewAddress(addr)
	ths.TrueHeading = tok.CommaUnsignedFloat()
	ths.ModeIndicator = tok.CommaOneByteOf(modeIndicators)
	tok.EndOfData()
	return &ths, tok.Err()
}

func (ths THS) Address() nmea.Address {
	return ths.address
}
