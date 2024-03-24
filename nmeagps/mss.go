package nmeagps

import "github.com/twpayne/go-nmea"

type MSS struct {
	address            Address
	SignalStrength     int
	SignalToNoiseRatio int
	BeaconFrequencyKHz float64
	BeaconBitRate      int
	ChannelNumber      nmea.Optional[int]
}

func ParseMSS(addr string, tok *nmea.Tokenizer) (*MSS, error) {
	var mss MSS
	mss.address = NewAddress(addr)
	mss.SignalStrength = tok.CommaInt()
	mss.SignalToNoiseRatio = tok.CommaInt()
	mss.BeaconFrequencyKHz = tok.CommaFloat()
	mss.BeaconBitRate = tok.CommaUnsignedInt()
	mss.ChannelNumber = tok.CommaOptionalUnsignedInt()
	tok.EndOfData()
	return &mss, tok.Err()
}

func (mss MSS) Address() nmea.Addresser {
	return mss.address
}
