package flarm_test

import (
	"testing"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/flarm"
	"github.com/twpayne/go-nmea/nmeatest"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineNever),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(flarm.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$PFLAU,3,1,2,1,2,-30,2,-32,755*",
				Expected: &flarm.PFLAU{
					Address:          nmea.NewAddress("PFLAU"),
					RX:               3,
					TX:               1,
					GPS:              2,
					Power:            1,
					AlarmLevel:       2,
					RelativeBearing:  nmea.NewOptional(-30),
					AlarmType:        2,
					RelativeVertical: nmea.NewOptional(-32),
					RelativeDistance: nmea.NewOptional(755),
				},
			},
			{
				S: "$PFLAU,2,1,1,1,0,,0,,,*",
				Expected: &flarm.PFLAU{
					Address:    nmea.NewAddress("PFLAU"),
					RX:         2,
					TX:         1,
					GPS:        1,
					Power:      1,
					AlarmLevel: 0,
					AlarmType:  0,
				},
			},
			{
				S: "$PFLAU,2,1,2,1,1,-45,2,50,75,1A304C*",
				Expected: &flarm.PFLAU{
					Address:          nmea.NewAddress("PFLAU"),
					RX:               2,
					TX:               1,
					GPS:              2,
					Power:            1,
					AlarmLevel:       1,
					RelativeBearing:  nmea.NewOptional(-45),
					AlarmType:        2,
					RelativeVertical: nmea.NewOptional(50),
					RelativeDistance: nmea.NewOptional(75),
					ID:               nmea.NewOptional(0x1A304C),
				},
			},
			{
				S: "$PFLAU,2,1,2,1,1,0,41,0,0,A25703*",
				Expected: &flarm.PFLAU{
					Address:          nmea.NewAddress("PFLAU"),
					RX:               2,
					TX:               1,
					GPS:              2,
					Power:            1,
					AlarmLevel:       1,
					RelativeBearing:  nmea.NewOptional(0),
					AlarmType:        41,
					RelativeVertical: nmea.NewOptional(0),
					RelativeDistance: nmea.NewOptional(0),
					ID:               nmea.NewOptional(0xA25703),
				},
			},
			{
				S: "$PFLAA,0,-1234,1234,220,2,DD8F12,180,,30,-1.4,1*",
				Expected: &flarm.PFLAA{
					Address:          nmea.NewAddress("PFLAA"),
					RelativeNorth:    -1234,
					RelativeEast:     1234,
					RelativeVertical: 220,
					IDType:           nmea.NewOptional(2),
					ID:               nmea.NewOptional(0xDD8F12),
					Track:            180,
					GroundSpeed:      nmea.NewOptional(30),
					ClimbRate:        nmea.NewOptional(-1.4),
					AircraftType:     1,
				},
			},
		})
}
