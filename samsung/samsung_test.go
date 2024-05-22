package samsung_test

import (
	"testing"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeatest"
	"github.com/twpayne/go-nmea/samsung"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(samsung.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$PSAMCLK,2315,30430300,29,-77797,-17518,304303000*76",
				Expected: &samsung.PSAMCLK{
					Address:  nmea.NewAddress("PSAMCLK"),
					Unknown1: 2315,
					Unknown2: 30430300,
					Unknown3: 29,
					Unknown4: -77797,
					Unknown5: -17518,
					Unknown6: 304303000,
				},
			},
			{
				S: "$PSAMDLOK,87,0x000,0x00F*56",
				Expected: &samsung.PSAMDLOK{
					Address:  nmea.NewAddress("PSAMDLOK"),
					Unknown1: 87,
					Unknown2: 0x000,
					Unknown3: 0x00F,
				},
			},
			{
				S: "$PSAMID,SLL_SPOTNAV_4.7.2_9,*47",
				Expected: &samsung.PSAMID{
					Address:  nmea.NewAddress("PSAMID"),
					Unknown1: "SLL_SPOTNAV_4.7.2_9",
				},
			},
			{
				S: "$PSAMSA,123125.000,4710.9456,N,00831.2559,E,1,M,,*63",
				Expected: &samsung.PSAMSA{
					Address: nmea.NewAddress("PSAMSA"),
					TimeOfDay: nmea.NewOptional(nmea.TimeOfDay{
						Hour:   12,
						Minute: 31,
						Second: 25,
					}),
					Lat:      nmea.NewOptional(47.182426666666665),
					Lon:      nmea.NewOptional(8.520931666666666),
					Unknown1: nmea.NewOptional(1),
				},
			},
			{
				S: "$PSAMSA,,,,,,0,,,,,,,,*2D",
				Expected: &samsung.PSAMSA{
					Address:  nmea.NewAddress("PSAMSA"),
					Unknown1: nmea.NewOptional(0),
				},
			},
		},
	)
}
