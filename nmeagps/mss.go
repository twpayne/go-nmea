package nmeagps

import "github.com/twpayne/go-nmea"

type MSS struct {
	address            Address
	SignalStrength     int
	SignalToNoiseRatio int
	BeaconFrequency    float64
	BeaconBitRate      int
	ChannelNumber      int
}

func ParseMSS(addr string, tok *nmea.Tokenizer) (*MSS, error) {
	var mss MSS
	mss.address = NewAddress(addr)
	mss.SignalStrength = tok.CommaInt()
	mss.SignalToNoiseRatio = tok.CommaInt()
	mss.BeaconFrequency = 1000 * tok.CommaFloat()
	mss.BeaconBitRate = tok.CommaUnsignedInt()
	mss.ChannelNumber = tok.CommaUnsignedInt()
	// tok.EndOfData() // FIXME remove, see trailing data discipline
	return &mss, tok.Err()
}

func (mss MSS) Address() nmea.Addresser {
	return mss.address
}
