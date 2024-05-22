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
	var pgrmb PGRMB
	pgrmb.Address = nmea.NewAddress(addr)
	pgrmb.BeaconTuneFrequencyKHz = tok.CommaFloat()
	pgrmb.BeaconBitRate = tok.CommaUnsignedInt()
	pgrmb.BeaconSNR = tok.CommaOptionalUnsignedInt()
	pgrmb.BeaconDataQuality = tok.CommaOptionalUnsignedInt()
	pgrmb.DistanceToBeaconReferenceStationKM = tok.CommaOptionalIntCommaUnit('K')
	pgrmb.BeaconReceiverCommunicationStatus = tok.CommaOptionalUnsignedInt()
	pgrmb.DGPSFixSource = tok.CommaOneByteOf("NRW")
	pgrmb.DGPSMode = tok.CommaOneByteOf("ANRW")
	tok.EndOfData()
	return &pgrmb, tok.Err()
}
