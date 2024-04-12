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
					Rx:               3,
					Tx:               1,
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
					Rx:         2,
					Tx:         1,
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
					Rx:               2,
					Tx:               1,
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
					Rx:               2,
					Tx:               1,
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
					RelativeEast:     nmea.NewOptional(1234),
					RelativeVertical: 220,
					IDType:           nmea.NewOptional(2),
					ID:               nmea.NewOptional(0xDD8F12),
					Track:            nmea.NewOptional(180),
					GroundSpeed:      nmea.NewOptional(30),
					ClimbRate:        nmea.NewOptional(-1.4),
					AircraftType:     1,
				},
			},
			{
				S: "$PFLAE,A,0,0*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  nmea.NewOptional(0),
					ErrorCode: nmea.NewOptional(0),
				},
			},
			{
				S: "$PFLAE,A*",
				Expected: &flarm.PFLAEAnswer{
					Address: nmea.NewAddress("PFLAE"),
				},
			},
			{
				S: "$PFLAE,A,2,81*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  nmea.NewOptional(2),
					ErrorCode: nmea.NewOptional(0x81),
				},
			},
			{
				S: "$PFLAE,A,3,11,Software expiry*",
				Expected: &flarm.PFLAEAnswer{
					Address:   nmea.NewAddress("PFLAE"),
					Severity:  nmea.NewOptional(3),
					ErrorCode: nmea.NewOptional(0x11),
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
			{
				S: "$PFLAC,A,ID,FFFFFF*",
				Expected: &flarm.PFLACAnswer{
					Address:           nmea.NewAddress("PFLAC"),
					ConfigurationItem: "ID",
					Values:            []string{"FFFFFF"},
				},
			},
			{
				S: "$PFLAC,A,RADIOID,2,DF202B*",
				Expected: &flarm.PFLACAnswer{
					Address:           nmea.NewAddress("PFLAC"),
					ConfigurationItem: "RADIOID",
					Values:            []string{"2", "DF202B"},
				},
			},
			{
				S: "$PFLAJ,A,1,1,0*",
				Expected: &flarm.PFLAJAnswer{
					Address:              nmea.NewAddress("PFLAJ"),
					FlightState:          1,
					FlightRecorderState:  1,
					TISBADSRClientStatus: nmea.NewOptional(0),
				},
			},
			{
				S: "$PFLAN,A,RANGE,RFTOP,A,5600,4800,3600,2400,1200,1200,*",
				Expected: &flarm.PFLANRangeStatisticAnswer{
					Address:       nmea.NewAddress("PFLAN"),
					StatisticType: "RFTOP",
					Channel:       'A',
					Values: []nmea.Optional[int]{
						nmea.NewOptional(5600),
						nmea.NewOptional(4800),
						nmea.NewOptional(3600),
						nmea.NewOptional(2400),
						nmea.NewOptional(1200),
						nmea.NewOptional(1200),
						{},
					},
				},
			},
			{
				S: "$PFLAN,A,RANGE,RFCNT,A,54,121,65,41,87,98,*",
				Expected: &flarm.PFLANRangeStatisticAnswer{
					Address:       nmea.NewAddress("PFLAN"),
					StatisticType: "RFCNT",
					Channel:       'A',
					Values: []nmea.Optional[int]{
						nmea.NewOptional(54),
						nmea.NewOptional(121),
						nmea.NewOptional(65),
						nmea.NewOptional(41),
						nmea.NewOptional(87),
						nmea.NewOptional(98),
						{},
					},
				},
			},
			{
				S: "$PFLAN,A,RANGE,RFDEV,B,1800,1100,1500,900,1200,1300,*",
				Expected: &flarm.PFLANRangeStatisticAnswer{
					Address:       nmea.NewAddress("PFLAN"),
					StatisticType: "RFDEV",
					Channel:       'B',
					Values: []nmea.Optional[int]{
						nmea.NewOptional(1800),
						nmea.NewOptional(1100),
						nmea.NewOptional(1500),
						nmea.NewOptional(900),
						nmea.NewOptional(1200),
						nmea.NewOptional(1300),
						{},
					},
				},
			},
			{
				S: "$PFLAN,A,RANGE,STATS,5000*",
				Expected: &flarm.PFLANRangeStatsAnswer{
					Address:           nmea.NewAddress("PFLAN"),
					NumberOfPointsTop: 5000,
				},
			},
			{
				S: "$PFLAN,A,RANGE*",
				Expected: &flarm.PFLANRangeAnswer{
					Address: nmea.NewAddress("PFLAN"),
				},
			},
			{
				S: "$PFLAN,A,RESET*",
				Expected: &flarm.PFLANResetAnswer{
					Address: nmea.NewAddress("PFLAN"),
				},
			},
			{
				S: "$PFLAF,A,1*",
				Expected: &flarm.PFLAFAnswer{
					Address:        nmea.NewAddress("PFLAF"),
					ScenarioNumber: 1,
				},
			},
			{
				S: "$PFLAF,A,ERROR,COMMAND*",
				Expected: &flarm.PFLAFError{
					Address:   nmea.NewAddress("PFLAF"),
					ErrorType: "COMMAND",
				},
			},
			{
				S: "$PFLAL,12224002NbWFCFcMN?lknsqrbser;NAKELu[*",
				Expected: &flarm.PFLAL{
					Address:      nmea.NewAddress("PFLAL"),
					DebugMessage: "12224002NbWFCFcMN?lknsqrbser;NAKELu[",
				},
			},
			{
				S: "$PFLAL,122242GPS 7 39*",
				Expected: &flarm.PFLAL{
					Address:      nmea.NewAddress("PFLAL"),
					DebugMessage: "122242GPS 7 39",
				},
			},
			{
				S: "$PFLAA,0,4964,,0,1,123456,,,,,0,0,6,*",
				Expected: &flarm.PFLAA{
					Address:          nmea.NewAddress("PFLAA"),
					AlarmLevel:       0,
					RelativeNorth:    4964,
					RelativeVertical: 0,
					IDType:           nmea.NewOptional(1),
					ID:               nmea.NewOptional(0x123456),
					NoTrack:          nmea.NewOptional(0),
					Source:           nmea.NewOptional(6),
				},
			},
		})
}
