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
	var pgrmt PGRMT
	pgrmt.Address = nmea.NewAddress(addr)
	pgrmt.ProductModelAndSoftwareVersion = tok.CommaString()
	pgrmt.ROMChecksumTest = tok.CommaOptionalOneByteOf("FP")
	pgrmt.ReceiverFailureDiscrete = tok.CommaOptionalOneByteOf("FP")
	pgrmt.StoredDataLost = tok.CommaOptionalOneByteOf("LR")
	pgrmt.RealTimeClockLost = tok.CommaOptionalOneByteOf("LR")
	pgrmt.OscillatorDriftDiscrete = tok.CommaOptionalOneByteOf("FP")
	pgrmt.DataCollectionDiscrete = tok.CommaOptionalLiteralByte('C')
	pgrmt.GPSSensorTemperature = tok.CommaOptionalInt()
	pgrmt.GPSSensorConfigurationData = tok.CommaOptionalOneByteOf("LR")
	tok.EndOfData()
	return &pgrmt, tok.Err()
}
