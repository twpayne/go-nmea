package garmin

import "github.com/twpayne/go-nmea"

type PGRMT struct {
	nmea.Address
	ProductModelAndSoftwareVersion string
	ROMChecksumTest                nmea.Optional[byte]
	ReceiverFailureDiscrete        nmea.Optional[byte]
	StoredDataLost                 nmea.Optional[byte]
	RealTimeClockLost              nmea.Optional[byte]
	OscillatorDriftDiscrete        nmea.Optional[byte]
	DataCollectionDiscrete         nmea.Optional[byte]
	GPSSensorTemperature           nmea.Optional[int]
	GPSSensorConfigurationData     nmea.Optional[byte]
}

func ParsePGRMT(addr string, tok *nmea.Tokenizer) (*PGRMT, error) {
	var t PGRMT
	t.Address = nmea.NewAddress(addr)
	t.ProductModelAndSoftwareVersion = tok.CommaString()
	t.ROMChecksumTest = tok.CommaOptionalOneByteOf("FP")
	t.ReceiverFailureDiscrete = tok.CommaOptionalOneByteOf("FP")
	t.StoredDataLost = tok.CommaOptionalOneByteOf("LR")
	t.RealTimeClockLost = tok.CommaOptionalOneByteOf("LR")
	t.OscillatorDriftDiscrete = tok.CommaOptionalOneByteOf("FP")
	t.DataCollectionDiscrete = tok.CommaOptionalLiteralByte('C')
	t.GPSSensorTemperature = tok.CommaOptionalInt()
	t.GPSSensorConfigurationData = tok.CommaOptionalOneByteOf("LR")
	tok.EndOfData()
	return &t, tok.Err()
}
