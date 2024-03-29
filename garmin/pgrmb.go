package garmin

import "github.com/twpayne/go-nmea"

type PGRMB struct {
	nmea.Address
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
	b.Address = nmea.NewAddress(addr)
	b.BeaconTuneFrequencyKHz = tok.CommaFloat()
	b.BeaconBitRate = tok.CommaUnsignedInt()
	b.BeaconSNR = tok.CommaOptionalUnsignedInt()
	b.BeaconDataQuality = tok.CommaOptionalUnsignedInt()
	b.DistanceToBeaconReferenceStationKM = tok.CommaOptionalIntCommaUnit('K')
	b.BeaconReceiverCommunicationStatus = tok.CommaOptionalUnsignedInt()
	b.DGPSFixSource = tok.CommaOneByteOf("NRW")
	b.DGPSMode = tok.CommaOneByteOf("ANRW")
	tok.EndOfData()
	return &b, tok.Err()
}
