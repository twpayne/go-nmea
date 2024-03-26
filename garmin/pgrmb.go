package garmin

import "github.com/twpayne/go-nmea"

type PGRMB struct {
	address                            nmea.Address
	BeaconTuneFrequencyKHz             float64
	BeaconBitRate                      int
	BeaconSNR                          nmea.Optional[int]
	BeaconDataQuality                  nmea.Optional[int]
	DistanceToBeaconReferenceStationKM nmea.Optional[int]
	BeaconReceiverCommunicationStatus  nmea.Optional[int]
	DGPSFixSource                      byte
	DGPSMode                           byte
}

func ParsePGRMB(addr string, tok *nmea.Tokenizer) (*PGRMB, error) {
	var b PGRMB
	b.address = nmea.NewAddress(addr)
	b.BeaconTuneFrequencyKHz = tok.CommaFloat()
	b.BeaconBitRate = tok.CommaUnsignedInt()
	b.BeaconSNR = tok.CommaOptionalUnsignedInt()
	b.BeaconDataQuality = tok.CommaOptionalUnsignedInt()
	b.DistanceToBeaconReferenceStationKM = tok.CommaOptionalUnsignedInt()
	tok.CommaLiteralByte('K')
	b.BeaconReceiverCommunicationStatus = tok.CommaOptionalUnsignedInt()
	b.DGPSFixSource = tok.CommaOneByteOf("NRW")
	b.DGPSMode = tok.CommaOneByteOf("ANRW")
	tok.EndOfData()
	return &b, tok.Err()
}

func (b PGRMB) Address() nmea.Addresser {
	return b.address
}
