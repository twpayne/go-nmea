package flarm

import (
	"fmt"

	"github.com/twpayne/go-nmea"
)

var AirportStatuses = map[int]string{
	0: "none: no status information available",
	1: "red: airport closed and should not be operated",
	2: "yellow: airport open but caution is required",
	3: "green: airport is safe to use",
}

// A PFLAMAnswer is part of the messaging feature.
type PFLAMAnswer struct {
	nmea.Address
	Result      string
	Error       string
	MessageType string
	Message     PFLAMMessage
}

// A PFLAMResponse is part of the messaging feature.
type PFLAMResponse struct {
	nmea.Address
	Scheduled   int
	Transmitted int
	FreeSlots   int
}

// A PFLAMUnsolicited is part of the messaging feature.
type PFLAMUnsolicited struct {
	nmea.Address
	IDType      int
	ID          int
	MsgType     int
	MessageType string
	Message     PFLAMMessage
}

type PFLAMMessage interface {
	MessageType() string
}

type PFLAMAircraftCallsignMessage struct {
	Name string
}

type PFLAMAircraftRegistrationMessage struct {
	Name string
}

type PFLAMAircraftTypeMessage struct {
	Name string
}

type PFLAMAirportInformationMessage struct {
	ICAOCode        string
	Lat             float64
	Lon             float64
	AltAMSLFeet     int
	RunwayInUse     nmea.Optional[int]
	VHFFrequencyMHz nmea.Optional[float64]
	AltimeterQNH    nmea.Optional[int]
	AirportStatus   nmea.Optional[int]
}

type PFLAMAirportWeatherMessage struct {
	WindDirection      int
	WindSpeedKnots     int
	WindGustsKnots     nmea.Optional[int]
	WindVariationBelow nmea.Optional[int]
	WindVariationAbove nmea.Optional[int]
	Visibility         int
	SkyCondition       nmea.Optional[string]
	BaseHeight         nmea.Optional[int]
	Temperature        int
	DewPoint           int
	PresentWeather     nmea.Optional[string]
}

type PFLAMOpenBroadcastMessage struct {
	Data []byte
}

type PFLAMOpenUnicastMessage struct {
	ID     int
	IDType int
	Data   []byte
}

type PFLAMPilotNameMessage struct {
	Name string
}

type PFLAMSensorMeasurementsMessage struct {
	IAS         nmea.Optional[int]
	Altimeter   nmea.Optional[int]
	Vario       nmea.Optional[float64]
	Temperature nmea.Optional[float64]
}

type PFLAMTeamNameMessage struct {
	Name string
}

type PFLAMVHFRadioFrequencyMessage struct {
	FrequencyAMHz float64
	FrequencyBMHz nmea.Optional[float64]
	FrequencyCMHz nmea.Optional[float64]
	FrequencyDMHz nmea.Optional[float64]
}

type PFLAMVersionMessage struct {
	Version              string
	DevType              int
	Region               int
	Hardware             string
	ObstName             nmea.Optional[string]
	ObstYear             nmea.Optional[string]
	TransponderInstalled int
	ModeSAlt             int
	OwnModeC             int
	PCASCal              int
}

type ParsePFLAMMessageFunc func(sp *nmea.Tokenizer) (PFLAMMessage, error)

var DefaultPFLAMMessageParsers = map[string]ParsePFLAMMessageFunc{
	"ACALL": MakeReturnPFLAMMessage(ParsePFLAMAircraftCallsignMessage),
	"AIRPT": MakeReturnPFLAMMessage(ParsePFLAMAirportInformationMessage),
	"AREG":  MakeReturnPFLAMMessage(ParsePFLAMAircraftRegistrationMessage),
	"ATYPE": MakeReturnPFLAMMessage(ParsePFLAMAircraftTypeMessage),
	"BCST":  MakeReturnPFLAMMessage(ParsePFLAMOpenBroadcastMessage),
	"METAR": MakeReturnPFLAMMessage(ParsePFLAMAirportWeatherMessage),
	"PNAME": MakeReturnPFLAMMessage(ParsePFLAMPilotNameMessage),
	"SENS":  MakeReturnPFLAMMessage(ParsePFLAMSensorMeasurementsMessage),
	"TEAM":  MakeReturnPFLAMMessage(ParsePFLAMTeamNameMessage),
	"UCST":  MakeReturnPFLAMMessage(ParsePFLAMOpenUnicastMessage),
	"VER":   MakeReturnPFLAMMessage(ParsePFLAMVersionMessage),
	"VHF":   MakeReturnPFLAMMessage(ParsePFLAMVHFRadioFrequencyMessage),
}

func MakeReturnPFLAMMessage[T PFLAMMessage](f func(sp *nmea.Tokenizer) (T, error)) func(sp *nmea.Tokenizer) (PFLAMMessage, error) {
	return func(tok *nmea.Tokenizer) (PFLAMMessage, error) {
		return f(tok)
	}
}

func ParsePFLAM(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("ARU")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		return ParsePFLAMAnswer(addr, tok)
	case 'R':
		return ParsePFLAMResponse(addr, tok)
	case 'U':
		return ParsePFLAMUnsolicited(addr, tok)
	default:
		return nil, fmt.Errorf("%c: unknown query type", queryType)
	}
}

func ParsePFLAMAnswer(addr string, tok *nmea.Tokenizer) (*PFLAMAnswer, error) {
	var pflamAnswer PFLAMAnswer
	pflamAnswer.Address = nmea.NewAddress(addr)
	pflamAnswer.Result = tok.CommaString()
	switch {
	case tok.Err() != nil:
		return nil, tok.Err()
	case pflamAnswer.Result == "OK":
		pflamAnswer.MessageType = tok.CommaString()
		if pflamMessageParser, ok := DefaultPFLAMMessageParsers[pflamAnswer.MessageType]; ok {
			pflamAnswer.Message, _ = pflamMessageParser(tok)
		} else {
			// FIXME consider returning unknown payload struct instead of error
			return &pflamAnswer, fmt.Errorf("%s: unknown payload type", pflamAnswer.MessageType)
		}
		return &pflamAnswer, tok.Err()
	case pflamAnswer.Result == "ERROR":
		pflamAnswer.Error = tok.CommaString()
		tok.AtEndOfData()
		return &pflamAnswer, tok.Err()
	default:
		return nil, fmt.Errorf("%s: unknown result", pflamAnswer.Result)
	}
}

func ParsePFLAMResponse(addr string, tok *nmea.Tokenizer) (*PFLAMResponse, error) {
	var r PFLAMResponse
	r.Address = nmea.NewAddress(addr)
	r.Scheduled = tok.CommaUnsignedInt()
	r.Transmitted = tok.CommaUnsignedInt()
	r.FreeSlots = tok.CommaUnsignedInt()
	return &r, tok.Err()
}

func ParsePFLAMUnsolicited(addr string, tok *nmea.Tokenizer) (*PFLAMUnsolicited, error) {
	var u PFLAMUnsolicited
	u.Address = nmea.NewAddress(addr)
	u.IDType = tok.CommaInt()
	u.ID = tok.CommaHex()
	u.MessageType = tok.CommaString()
	if pflamMessageParser, ok := DefaultPFLAMMessageParsers[u.MessageType]; ok {
		u.Message, _ = pflamMessageParser(tok)
	} else {
		// FIXME consider returning unknown payload struct instead of error
		return &u, fmt.Errorf("%s: unknown payload type", u.MessageType)
	}
	return &u, tok.Err()
}

func ParsePFLAMAircraftCallsignMessage(tok *nmea.Tokenizer) (*PFLAMAircraftCallsignMessage, error) {
	var m PFLAMAircraftCallsignMessage
	m.Name = string(tok.CommaHexBytes())
	tok.EndOfData()
	return &m, tok.Err()
}

func (p PFLAMAircraftCallsignMessage) MessageType() string {
	return "ACALL"
}

func ParsePFLAMAircraftRegistrationMessage(tok *nmea.Tokenizer) (*PFLAMAircraftRegistrationMessage, error) {
	var m PFLAMAircraftRegistrationMessage
	m.Name = string(tok.CommaHexBytes())
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMAircraftRegistrationMessage) MessageType() string {
	return "AREG"
}

func ParsePFLAMAircraftTypeMessage(tok *nmea.Tokenizer) (*PFLAMAircraftTypeMessage, error) {
	var m PFLAMAircraftTypeMessage
	m.Name = string(tok.CommaHexBytes())
	tok.EndOfData()
	return &m, tok.Err()
}

func (p PFLAMAircraftTypeMessage) MessageType() string {
	return "ATYPE"
}

func ParsePFLAMAirportInformationMessage(tok *nmea.Tokenizer) (*PFLAMAirportInformationMessage, error) {
	var m PFLAMAirportInformationMessage
	m.ICAOCode = tok.CommaString()
	m.Lat = tok.CommaFloat()
	m.Lon = tok.CommaFloat()
	m.AltAMSLFeet = tok.CommaInt()
	m.RunwayInUse = tok.CommaOptionalUnsignedInt()
	m.VHFFrequencyMHz = tok.CommaOptionalUnsignedFloat()
	m.AltimeterQNH = tok.CommaOptionalUnsignedInt()
	m.AirportStatus = tok.CommaOptionalUnsignedInt()
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMAirportInformationMessage) MessageType() string {
	return "AIRPT"
}

func ParsePFLAMAirportWeatherMessage(tok *nmea.Tokenizer) (*PFLAMAirportWeatherMessage, error) {
	var m PFLAMAirportWeatherMessage
	m.WindDirection = tok.CommaUnsignedInt()
	m.WindSpeedKnots = tok.CommaUnsignedInt()
	m.WindGustsKnots = tok.CommaOptionalUnsignedInt()
	m.WindVariationBelow = tok.CommaOptionalUnsignedInt()
	m.WindVariationAbove = tok.CommaOptionalUnsignedInt()
	m.Visibility = tok.CommaUnsignedInt()
	m.SkyCondition = tok.CommaOptionalString()
	m.BaseHeight = tok.CommaOptionalUnsignedInt()
	m.Temperature = tok.CommaInt()
	m.DewPoint = tok.CommaInt()
	m.PresentWeather = tok.CommaOptionalString()
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMAirportWeatherMessage) MessageType() string {
	return "METAR"
}

func ParsePFLAMOpenBroadcastMessage(tok *nmea.Tokenizer) (*PFLAMOpenBroadcastMessage, error) {
	var p PFLAMOpenBroadcastMessage
	p.Data = tok.CommaHexBytes()
	tok.EndOfData()
	return &p, tok.Err()
}

func (m PFLAMOpenBroadcastMessage) MessageType() string {
	return "BCST"
}

func ParsePFLAMOpenUnicastMessage(tok *nmea.Tokenizer) (*PFLAMOpenUnicastMessage, error) {
	var m PFLAMOpenUnicastMessage
	m.IDType = tok.CommaUnsignedInt()
	m.ID = tok.CommaHex()
	m.Data = tok.CommaHexBytes()
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMOpenUnicastMessage) MessageType() string {
	return "UCST"
}

func ParsePFLAMPilotNameMessage(tok *nmea.Tokenizer) (*PFLAMPilotNameMessage, error) {
	var m PFLAMPilotNameMessage
	m.Name = string(tok.CommaHexBytes())
	tok.EndOfData()
	return &m, tok.Err()
}

func (p PFLAMPilotNameMessage) MessageType() string {
	return "PNAME"
}

func ParsePFLAMSensorMeasurementsMessage(tok *nmea.Tokenizer) (*PFLAMSensorMeasurementsMessage, error) {
	var m PFLAMSensorMeasurementsMessage
	m.IAS = tok.CommaOptionalInt()
	m.Altimeter = tok.CommaOptionalInt()
	m.Vario = tok.CommaOptionalFloat()
	m.Temperature = tok.CommaOptionalFloat()
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMSensorMeasurementsMessage) MessageType() string {
	return "SENS"
}

func ParsePFLAMTeamNameMessage(tok *nmea.Tokenizer) (*PFLAMTeamNameMessage, error) {
	var m PFLAMTeamNameMessage
	m.Name = string(tok.CommaHexBytes())
	tok.EndOfData()
	return &m, tok.Err()
}

func (m PFLAMTeamNameMessage) MessageType() string {
	return "TEAM"
}

func ParsePFLAMVHFRadioFrequencyMessage(tok *nmea.Tokenizer) (*PFLAMVHFRadioFrequencyMessage, error) {
	var m PFLAMVHFRadioFrequencyMessage
	// FIXME use integer Hz for frequencies, not float64 MHz
	m.FrequencyAMHz = tok.CommaFloat()
	m.FrequencyBMHz = tok.CommaOptionalUnsignedFloat()
	m.FrequencyCMHz = tok.CommaOptionalUnsignedFloat()
	m.FrequencyDMHz = tok.CommaOptionalUnsignedFloat()
	tok.EndOfData()
	return &m, tok.Err()
}

func (p PFLAMVHFRadioFrequencyMessage) MessageType() string {
	return "VHF"
}

func ParsePFLAMVersionMessage(tok *nmea.Tokenizer) (*PFLAMVersionMessage, error) {
	var m PFLAMVersionMessage
	m.Version = tok.CommaString()
	m.DevType = tok.CommaInt()
	m.Region = tok.CommaInt()
	m.Hardware = tok.CommaString()
	m.ObstName = tok.CommaOptionalString()
	m.ObstYear = tok.CommaOptionalString()
	m.TransponderInstalled = tok.CommaInt()
	m.ModeSAlt = tok.CommaInt()
	m.OwnModeC = tok.CommaInt()
	m.PCASCal = tok.CommaInt()
	tok.EndOfData()
	return &m, tok.Err()
}

func (p PFLAMVersionMessage) MessageType() string {
	return "VER"
}
