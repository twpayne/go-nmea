package flarm

import (
	"fmt"
	"time"

	"github.com/twpayne/go-nmea"
)

type PFLANRangeAnswer struct {
	nmea.Address
}

type PFLANRangeStatisticAnswer struct {
	nmea.Address
	StatisticType string
	Channel       byte
	Values        []nmea.Optional[int]
}

type PFLANRangeStatsAnswer struct {
	nmea.Address
	NumberOfPointsTop int
}

type PFLANRangeTimeSpanAnswer struct {
	nmea.Address
	Start time.Time
	Stop  time.Time
}

type PFLANResetAnswer struct {
	nmea.Address
}

func ParsePFLAN(addr string, tok *nmea.Tokenizer) (nmea.Sentence, error) {
	queryType := tok.CommaOneByteOf("A")
	if err := tok.Err(); err != nil {
		return nil, err
	}
	switch queryType {
	case 'A':
		switch s := tok.CommaString(); s {
		case "RANGE":
			if tok.AtEndOfData() {
				var pflanRangeAnswer PFLANRangeAnswer
				pflanRangeAnswer.Address = nmea.NewAddress(addr)
				tok.EndOfData()
				return &pflanRangeAnswer, tok.Err()
			}
			switch s := tok.CommaString(); s {
			case "RFTOP", "RFCNT", "RFDEV":
				var pflanRangeStatisticsAnswer PFLANRangeStatisticAnswer
				pflanRangeStatisticsAnswer.Address = nmea.NewAddress(addr)
				pflanRangeStatisticsAnswer.StatisticType = s
				pflanRangeStatisticsAnswer.Channel = tok.CommaOneByteOf("AB")
				for !tok.AtEndOfData() {
					value := tok.CommaOptionalInt()
					pflanRangeStatisticsAnswer.Values = append(pflanRangeStatisticsAnswer.Values, value)
				}
				tok.EndOfData()
				return &pflanRangeStatisticsAnswer, tok.Err()
			case "STATS":
				var pflanRangeStatsAnswer PFLANRangeStatsAnswer
				pflanRangeStatsAnswer.Address = nmea.NewAddress(addr)
				pflanRangeStatsAnswer.NumberOfPointsTop = tok.CommaUnsignedInt()
				tok.EndOfData()
				return &pflanRangeStatsAnswer, tok.Err()
			case "TIMESPAN":
				var pflanRangeTimeSpanAnswer PFLANRangeTimeSpanAnswer
				pflanRangeTimeSpanAnswer.Address = nmea.NewAddress(addr)
				pflanRangeTimeSpanAnswer.Start = time.Unix(int64(tok.CommaUnsignedInt()), 0)
				pflanRangeTimeSpanAnswer.Stop = time.Unix(int64(tok.CommaUnsignedInt()), 0)
				tok.EndOfData()
				return &pflanRangeTimeSpanAnswer, tok.Err()
			default:
				return nil, fmt.Errorf("%s: unexpected string", s)
			}
		case "RESET":
			var pflanResetAnswer PFLANResetAnswer
			pflanResetAnswer.Address = nmea.NewAddress(addr)
			tok.EndOfData()
			return &pflanResetAnswer, tok.Err()
		default:
			return nil, fmt.Errorf("%s: unexpected string", s)
		}
	default:
		return nil, &UnknownQueryTypeError{
			QueryType: queryType,
		}
	}
}
