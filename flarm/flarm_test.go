package flarm_test

import (
	"testing"
	"time"

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
			{
				S: "$PFLAE,A,0,0*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  0,
					ErrorCode: 0,
				},
			},
			{
				S: "$PFLAE,A,2,81*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  2,
					ErrorCode: 0x81,
				},
			},
			{
				S: "$PFLAE,A,3,11,Software expiry*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  3,
					ErrorCode: 0x11,
					Message:   nmea.NewOptional("Software expiry"),
				},
			},
			{
				S: "$PFLAV,A,2.00,5.00,alps20110221_*",
				Expected: &flarm.PFLAVAnswer{
					Address:         nmea.NewAddress("PFLAV"),
					HardwareVersion: "2.00",
					SoftwareVersion: "5.00",
					ObstacleVersion: nmea.NewOptional("alps20110221_"),
				},
			},
			{
				S: "$PFLAV,A,2.00,5.00,*",
				Expected: &flarm.PFLAVAnswer{
					Address:         nmea.NewAddress("PFLAV"),
					HardwareVersion: "2.00",
					SoftwareVersion: "5.00",
				},
			},
			{
				S: "$PFLAQ,OBST,,10*",
				Expected: &flarm.PFLAQ{
					Address:   nmea.NewAddress("PFLAQ"),
					Operation: "OBST",
					Info:      nmea.NewOptional(""),
					Progress:  10,
				},
			},
			{
				S: "$PFLAQ,IGC,2A8GJ7K1.IGC,55*",
				Expected: &flarm.PFLAQ{
					Address:   nmea.NewAddress("PFLAQ"),
					Operation: "IGC",
					Info:      nmea.NewOptional("2A8GJ7K1.IGC"),
					Progress:  55,
				},
			},
			{
				S: "$PFLAQ,IGC,2A8GJ7K1.IGC,65*",
				Expected: &flarm.PFLAQ{
					Address:   nmea.NewAddress("PFLAQ"),
					Operation: "IGC",
					Info:      nmea.NewOptional("2A8GJ7K1.IGC"),
					Progress:  65,
				},
			},
			{
				S: "$PFLAQ,IGC,25*",
				Expected: &flarm.PFLAQ{
					Address:   nmea.NewAddress("PFLAQ"),
					Operation: "IGC",
					Progress:  25,
				},
			},
			{
				S: "$PFLAQ,IGC,60*",
				Expected: &flarm.PFLAQ{
					Address:   nmea.NewAddress("PFLAQ"),
					Operation: "IGC",
					Progress:  60,
				},
			},
			{
				S: "$PFLAO,1,1,471122335,85577812,2000,100,4550,1432832400,DF4738,2,41*",
				Expected: &flarm.PFLAO{
					Address:       nmea.NewAddress("PFLAO"),
					AlarmLevel:    1,
					Inside:        1,
					Lat:           471122335,
					Lon:           85577812,
					Radius:        2000,
					Bottom:        100,
					Top:           4550,
					ActivityLimit: time.Date(2015, time.May, 28, 17, 0, 0, 0, time.UTC),
					ID:            0xDF4738,
					IDType:        2,
					ZoneType:      0x41,
				},
			},
			{
				S: "$PFLAI,IGCREADOUT,OK*",
				Expected: &flarm.PFLAI{
					Address: nmea.NewAddress("PFLAI"),
					Value:   "IGCREADOUT",
					Result:  "OK",
				},
			},
			{
				S: "$PFLAI,PILOTEVENT,OK*",
				Expected: &flarm.PFLAI{
					Address: nmea.NewAddress("PFLAI"),
					Value:   "PILOTEVENT",
					Result:  "OK",
				},
			},
			{
				S: "$PFLAC,A,ERROR*",
				Expected: &flarm.PFLACError{
					Address: nmea.NewAddress("PFLAC"),
				},
			},
		})
}
