package lk8000_test

import (
	"testing"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/lk8000"
	"github.com/twpayne/go-nmea/nmeatest"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(lk8000.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$LK8EX1,101523,99999,-2,18.7,999,*18",
				Expected: &lk8000.LK8EX1{
					Address:     nmea.NewAddress("LK8EX1"),
					RawPressure: nmea.NewOptional(1015.23),
					VarioCMS:    nmea.NewOptional(-2.0),
					Temperature: nmea.NewOptional(18.7),
				},
			},
			{
				S: "$LK8EX1,99860,99999,9999,25,1000,*12",
				Expected: &lk8000.LK8EX1{
					Address:        nmea.NewAddress("LK8EX1"),
					RawPressure:    nmea.NewOptional(998.6),
					Temperature:    nmea.NewOptional(25.0),
					BatteryPercent: nmea.NewOptional(0),
				},
			},
			{
				S: "$LK8EX1,93459,676,13,20,999,*2f",
				Expected: &lk8000.LK8EX1{
					Address:     nmea.NewAddress("LK8EX1"),
					RawPressure: nmea.NewOptional(934.59),
					Alt:         nmea.NewOptional(676.0),
					VarioCMS:    nmea.NewOptional(13.0),
					Temperature: nmea.NewOptional(20.0),
				},
			},
			{
				S: "$LK8EX1,98684,99999,-4,28,1100,*02",
				Expected: &lk8000.LK8EX1{
					Address:        nmea.NewAddress("LK8EX1"),
					RawPressure:    nmea.NewOptional(986.84),
					VarioCMS:       nmea.NewOptional(-4.0),
					Temperature:    nmea.NewOptional(28.0),
					BatteryPercent: nmea.NewOptional(100),
				},
			},
			{
				S: "$LK8EX1,99545,149,1,26,5.10*18",
				Expected: &lk8000.LK8EX1{
					Address:        nmea.NewAddress("LK8EX1"),
					RawPressure:    nmea.NewOptional(995.45),
					Alt:            nmea.NewOptional(149.0),
					VarioCMS:       nmea.NewOptional(1.0),
					Temperature:    nmea.NewOptional(26.0),
					BatteryVoltage: nmea.NewOptional(5.1),
				},
			},
		},
	)
}
