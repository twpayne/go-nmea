package ublox_test

import (
	"testing"
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeatest"
	"github.com/twpayne/go-nmea/ublox"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(ublox.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$PUBX,00,081350.00,4717.113210,N,00833.915187,E,546.589,G3,2.1,2.0,0.007,77.52,0.007,,0.92,1.19,0.77,9,0,0*5F",
				Expected: &ublox.Position{
					Address: nmea.NewAddress("PUBX"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   8,
						Minute: 13,
						Second: 50,
					},
					Lat:                47.28522016666667,
					Lon:                8.565253116666666,
					AltRef:             546.589,
					NavStat:            "G3",
					HorizAcc:           2.1,
					VertAcc:            2,
					SpeedOverGroundKPH: 0.007,
					CourseOverGround:   77.52,
					VertVel:            0.007,
					HDOP:               0.92,
					VDOP:               1.19,
					TDOP:               0.77,
					NumSVs:             9,
				},
			},
			{
				S: "$PUBX,40,GLL,1,0,0,0,0,0*5D",
				Expected: &ublox.Rate{
					Address: nmea.NewAddress("PUBX"),
					MsgID:   "GLL",
					RDDC:    1,
					RUS1:    0,
					RUS2:    0,
					RUSB:    0,
					RSPI:    0,
				},
			},
			{
				S: "$PUBX,03,11,23,-,,,45,010,29,-,,,46,013,07,-,,,42,015,08,U,067,31,42,025,10,U,195,33,46,026,18,U,326,08,39,026,17,-,,,32,015,26,U,306,66,48,025,27,U,073,10,36,026,28,U,089,61,46,024,15,-,,,39,014*0D",
				Expected: &ublox.Status{
					Address: nmea.NewAddress("PUBX"),
					N:       11,
					SatelliteStatuses: []ublox.SatelliteStatus{
						{
							SVID:   23,
							Status: 45,
							CNO:    45,
							Lck:    10,
						},
						{
							SVID:   29,
							Status: 45,
							CNO:    46,
							Lck:    13,
						},
						{
							SVID:   7,
							Status: 45,
							CNO:    42,
							Lck:    15,
						},
						{
							SVID:   8,
							Status: 85,
							Az:     nmea.NewOptional(67),
							El:     nmea.NewOptional(31),
							CNO:    42,
							Lck:    25,
						},
						{
							SVID:   10,
							Status: 85,
							Az:     nmea.NewOptional(195),
							El:     nmea.NewOptional(33),
							CNO:    46,
							Lck:    26,
						},
						{
							SVID:   18,
							Status: 85,
							Az:     nmea.NewOptional(326),
							El:     nmea.NewOptional(8),
							CNO:    39,
							Lck:    26,
						},
						{
							SVID:   17,
							Status: 45,
							CNO:    32,
							Lck:    15,
						},
						{
							SVID:   26,
							Status: 85,
							Az:     nmea.NewOptional(306),
							El:     nmea.NewOptional(66),
							CNO:    48,
							Lck:    25,
						},
						{
							SVID:   27,
							Status: 85,
							Az:     nmea.NewOptional(73),
							El:     nmea.NewOptional(10),
							CNO:    36,
							Lck:    26,
						},
						{
							SVID:   28,
							Status: 85,
							Az:     nmea.NewOptional(89),
							El:     nmea.NewOptional(61),
							CNO:    46,
							Lck:    24,
						},
						{
							SVID:   15,
							Status: 45,
							CNO:    39,
							Lck:    14,
						},
					},
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$PUBX,04,073731.00,091202,113851.00,1196,15D,1930035,-2660.664,43,*3C",
				Expected: &ublox.Time{
					Address:              nmea.NewAddress("PUBX"),
					Time:                 time.Date(2002, time.December, 9, 7, 37, 31, 0, time.UTC),
					UTCTimeOfWeek:        113851,
					UTCWeek:              1196,
					LeapSeconds:          15,
					LeapSecondsDefault:   true,
					ClockBias:            1930035,
					ClockDrift:           -2660.664,
					TimePulseGranularity: 43,
				},
			},
		})
}
