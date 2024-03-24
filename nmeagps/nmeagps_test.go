package nmeagps

// FIXME fix all skips

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea"
)

func TestParseSentence(t *testing.T) {
	for _, tc := range []struct {
		skip        string
		options     []nmea.ParserOption
		s           string
		expectedErr error
		expected    nmea.Sentence
	}{
		// u-blox examples from
		// https://content.u-blox.com/sites/default/files/products/documents/u-blox8-M8_ReceiverDescrProtSpec_UBX-13003221.pdf
		{
			s: "$GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F",
			expected: &DTM{
				address:  NewAddress("GPDTM"),
				Datum:    "W84",
				RefDatum: "W84",
			},
		},
		{
			s: "$GPGBS,235503.00,1.6,1.4,3.2,,,,,,*40",
			expected: &GBS{
				address: NewAddress("GPGBS"),
				TimeOfDay: TimeOfDay{
					Hour:   23,
					Minute: 55,
					Second: 3,
				},
				ErrLat: 1.6,
				ErrLon: 1.4,
				ErrAlt: 3.2,
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPGBS,235458.00,1.4,1.3,3.1,03,,-21.4,3.8,1,0*5B",
			expected: &GBS{
				address: NewAddress("GPGBS"),
				TimeOfDay: TimeOfDay{
					Hour:   23,
					Minute: 54,
					Second: 58,
				},
				ErrLat:   1.4,
				ErrLon:   1.3,
				ErrAlt:   3.1,
				SVID:     nmea.NewOptional(3),
				Bias:     nmea.NewOptional(-21.4),
				StdDev:   nmea.NewOptional(3.8),
				SystemID: nmea.NewOptional(1),
				SignalID: nmea.NewOptional(0),
			},
		},
		{
			s: "$GPGGA,092725.00,4717.11399,N,00833.91590,E,1,08,1.01,499.6,M,48.0,M,,*5B",
			expected: &GGA{
				address: NewAddress("GPGGA"),
				TimeOfDay: TimeOfDay{
					Hour:   9,
					Minute: 27,
					Second: 25,
				},
				Lat:                              47.285233166666664,
				Lon:                              8.565265,
				FixQuality:                       1,
				NumberOfSatellites:               8,
				HDOP:                             1.01,
				Alt:                              499.6,
				HeightOfGeoidAboveWGS84Ellipsoid: 48,
			},
		},
		{
			s: "$GPGLL,4717.11364,N,00833.91565,E,092321.00,A,A*60",
			expected: &GLL{
				address: NewAddress("GPGLL"),
				Lat:     47.28522733333333,
				Lon:     8.565260833333333,
				TimeOfDay: TimeOfDay{
					Hour:   9,
					Minute: 23,
					Second: 21,
				},
				Status:  'A',
				PosMode: 'A',
			},
		},
		{
			s: "$GNGNS,103600.01,5114.51176,N,00012.29380,W,ANNN,07,1.18,111.5,45.6,,,V*00",
			expected: &GNS{
				address: NewAddress("GNGNS"),
				TimeOfDay: TimeOfDay{
					Hour:       10,
					Minute:     36,
					Nanosecond: 10000000,
				},
				Lat:       nmea.NewOptional(51.24186266666667),
				Lon:       nmea.NewOptional(-0.20489666666666664),
				PosMode:   []byte{'A', 'N', 'N', 'N'},
				NumSV:     7,
				HDOP:      nmea.NewOptional(1.18),
				Alt:       nmea.NewOptional(111.5),
				Sep:       nmea.NewOptional(45.6),
				NavStatus: 'V',
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GNGNS,122310.2,3722.425671,N,12258.856215,W,DAAA,14,0.9,1005.543,6.5,,,V*0E",
			expected: &GNS{
				address: NewAddress("GNGNS"),
				TimeOfDay: TimeOfDay{
					Hour:       12,
					Minute:     23,
					Second:     10,
					Nanosecond: 200000000,
				},
				Lat:       nmea.NewOptional(37.373761183333336),
				Lon:       nmea.NewOptional(-122.98093691666666),
				PosMode:   []byte{'D', 'A', 'A', 'A'},
				NumSV:     14,
				HDOP:      nmea.NewOptional(0.9),
				Alt:       nmea.NewOptional(1005.543),
				Sep:       nmea.NewOptional(6.5),
				NavStatus: 'V',
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPGNS,122310.2,,,,,,07,,,,5.2,23,V*02",
			expected: &GNS{
				address: NewAddress("GPGNS"),
				TimeOfDay: TimeOfDay{
					Hour:       12,
					Minute:     23,
					Second:     10,
					Nanosecond: 200000000,
				},
				NumSV:       7,
				DiffAge:     nmea.NewOptional(5.2),
				DiffStation: nmea.NewOptional(23),
				NavStatus:   'V',
			},
		},
		{
			s: "$GNGRS,104148.00,1,2.6,2.2,-1.6,-1.1,-1.7,-1.5,5.8,1.7,,,,,1,1*52",
			expected: &GRS{
				address: NewAddress("GNGRS"),
				TimeOfDay: TimeOfDay{
					Hour:   10,
					Minute: 41,
					Second: 48,
				},
				Mode: 1,
				Residuals: []nmea.Optional[float64]{
					nmea.NewOptional(2.6),
					nmea.NewOptional(2.2),
					nmea.NewOptional(-1.6),
					nmea.NewOptional(-1.1),
					nmea.NewOptional(-1.7),
					nmea.NewOptional(-1.5),
					nmea.NewOptional(5.8),
					nmea.NewOptional(1.7),
					{},
					{},
					{},
					{},
				},
				SystemID: 1,
				SignalID: 1,
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GNGRS,104148.00,1,,0.0,2.5,0.0,,2.8,,,,,,,1,5*52",
			expected: &GRS{
				address: NewAddress("GNGRS"),
				TimeOfDay: TimeOfDay{
					Hour:   10,
					Minute: 41,
					Second: 48,
				},
				Mode: 1,
				Residuals: []nmea.Optional[float64]{
					{},
					nmea.NewOptional(0.0),
					nmea.NewOptional(2.5),
					nmea.NewOptional(0.0),
					{},
					nmea.NewOptional(2.8),
					{},
					{},
					{},
					{},
					{},
					{},
				},
				SystemID: 1,
				SignalID: 5,
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPGSA,A,3,23,29,07,08,09,18,26,28,,,,,1.94,1.18,1.54,1*0D",
			expected: &GSA{
				address: NewAddress("GPGSA"),
				OpMode:  'A',
				NavMode: 3,
				SVIDs: []nmea.Optional[int]{
					nmea.NewOptional(23),
					nmea.NewOptional(29),
					nmea.NewOptional(7),
					nmea.NewOptional(8),
					nmea.NewOptional(9),
					nmea.NewOptional(18),
					nmea.NewOptional(26),
					nmea.NewOptional(28),
					{},
					{},
					{},
					{},
				},
				PDOP:     nmea.NewOptional(1.94),
				HDOP:     nmea.NewOptional(1.18),
				VDOP:     nmea.NewOptional(1.54),
				SystemID: nmea.NewOptional(1),
			},
		},
		{
			s: "$GPGST,082356.00,1.8,,,,1.7,1.3,2.2*7E",
			expected: &GST{
				address: NewAddress("GPGST"),
				TimeOfDay: TimeOfDay{
					Hour:   8,
					Minute: 23,
					Second: 56,
				},
				RangeRMS:  1.8,
				LatStdDev: 1.7,
				LonStdDev: 1.3,
				AltStdDev: 2.2,
			},
		},
		{
			s: "$GPGSV,3,1,09,09,,,17,10,,,40,12,,,49,13,,,35,1*6F",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  3,
				MsgNum:  1,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 9,
						CNO:  nmea.NewOptional(17),
					},
					{
						SVID: 10,
						CNO:  nmea.NewOptional(40),
					},
					{
						SVID: 12,
						CNO:  nmea.NewOptional(49),
					},
					{
						SVID: 13,
						CNO:  nmea.NewOptional(35),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			s: "$GPGSV,3,2,09,15,,,44,17,,,45,19,,,44,24,,,50,1*64",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  3,
				MsgNum:  2,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 15,
						CNO:  nmea.NewOptional(44),
					},
					{
						SVID: 17,
						CNO:  nmea.NewOptional(45),
					},
					{
						SVID: 19,
						CNO:  nmea.NewOptional(44),
					},
					{
						SVID: 24,
						CNO:  nmea.NewOptional(50),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			s: "$GPGSV,3,3,09,25,,,40,1*6E",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  3,
				MsgNum:  3,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 25,
						CNO:  nmea.NewOptional(40),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			s: "$GPGSV,1,1,03,12,,,42,24,,,47,32,,,37,5*66",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  1,
				MsgNum:  1,
				NumSV:   3,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 12,
						CNO:  nmea.NewOptional(42),
					},
					{
						SVID: 24,
						CNO:  nmea.NewOptional(47),
					},
					{
						SVID: 32,
						CNO:  nmea.NewOptional(37),
					},
				},
				SignalID: nmea.NewOptional(5),
			},
		},
		{
			s: "$GAGSV,1,1,00,2*76",
			expected: &GSV{
				address:  NewAddress("GAGSV"),
				NumMsg:   1,
				MsgNum:   1,
				NumSV:    0,
				SignalID: nmea.NewOptional(2),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPRMC,083559.00,A,4717.11437,N,00833.91522,E,0.004,77.52,091202,,,A,V*57",
			expected: &RMC{
				address:           NewAddress("GPRMC"),
				Time:              time.Date(2002, time.December, 9, 8, 35, 59, 0, time.UTC),
				Status:            'A',
				Lat:               47.2852395,
				Lon:               8.565253666666667,
				SpeedOverGroundKN: 0.004,
				CourseOverGround:  nmea.NewOptional(77.52),
				ModeIndicator:     'A',
				NavStatus:         'V',
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPTHS,77.52,E*32",
			expected: &THS{
				address:       NewAddress("GPTHS"),
				TrueHeading:   77.52,
				ModeIndicator: 'E',
			},
		},
		{
			s: "$GPTXT,01,01,02,u-blox ag - www.u-blox.com*50",
			expected: &TXT{
				address: NewAddress("GPTXT"),
				NumMsg:  1,
				MsgNum:  1,
				MsgType: 2,
				Text:    "u-blox ag - www.u-blox.com",
			},
		},
		{
			s: "$GPTXT,01,01,02,ANTARIS ATR0620 HW 00000040*67",
			expected: &TXT{
				address: NewAddress("GPTXT"),
				NumMsg:  1,
				MsgNum:  1,
				MsgType: 2,
				Text:    "ANTARIS ATR0620 HW 00000040",
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPVLW,,N,,N,15.8,N,1.2,N*06",
			expected: &VLW{
				address:               NewAddress("GPVLW"),
				TotalGroundDistanceNM: nmea.NewOptional(15.8),
				GroundDistanceNM:      nmea.NewOptional(1.2),
			},
		},
		{
			s: "$GPVTG,77.52,T,,M,0.004,N,0.008,K,A*06",
			expected: &VTG{
				address:              NewAddress("GPVTG"),
				TrueCourseOverGround: 77.52,
				SpeedOverGroundKN:    0.004,
				SpeedOverGroundKPH:   0.008,
				ModeIndicator:        'A',
			},
		},
		{
			s: "$GPZDA,082710.00,16,09,2002,00,00*64",
			expected: &ZDA{
				address:              NewAddress("GPZDA"),
				Time:                 time.Date(2002, time.September, 16, 8, 27, 10, 0, time.UTC),
				LocalTimeZoneHours:   0,
				LocalTimeZoneMinutes: 0,
			},
		},

		// sparkfun examples from https://www.sparkfun.com/datasheets/GPS/NMEA%20Reference%20Manual-Rev2.1-Dec07.pdf
		{
			s: "$GPGGA,002153.000,3342.6618,N,11751.3858,W,1,10,1.2,27.0,M,-34.2,M,,0000*5E",
			expected: &GGA{
				address: NewAddress("GPGGA"),
				TimeOfDay: TimeOfDay{
					Hour:   0,
					Minute: 21,
					Second: 53,
				},
				Lat:                              33.71103,
				Lon:                              -117.85643,
				FixQuality:                       1,
				NumberOfSatellites:               10,
				HDOP:                             1.2,
				Alt:                              27,
				HeightOfGeoidAboveWGS84Ellipsoid: -34.2,
				DGPSReferenceStationID:           "0000",
			},
		},
		{
			s: "$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,A*41",
			expected: &GLL{
				address: NewAddress("GPGLL"),
				Lat:     37.387458333333335,
				Lon:     -121.97236,
				TimeOfDay: TimeOfDay{
					Hour:       16,
					Minute:     12,
					Second:     29,
					Nanosecond: 487000000,
				},
				Status:  'A',
				PosMode: 'A',
			},
		},
		{
			s: "$GPGSA,A,3,07,02,26,27,09,04,15,,,,,,1.8,1.0,1.5*33",
			expected: &GSA{
				address: NewAddress("GPGSA"),
				OpMode:  'A',
				NavMode: 3,
				SVIDs: []nmea.Optional[int]{
					nmea.NewOptional(7),
					nmea.NewOptional(2),
					nmea.NewOptional(26),
					nmea.NewOptional(27),
					nmea.NewOptional(9),
					nmea.NewOptional(4),
					nmea.NewOptional(15),
					{},
					{},
					{},
					{},
					{},
				},
				PDOP: nmea.NewOptional(1.8),
				HDOP: nmea.NewOptional(1.0),
				VDOP: nmea.NewOptional(1.5),
			},
		},
		{
			s: "$GPGSV,2,1,07,07,79,048,42,02,51,062,43,26,36,256,42,27,27,138,42*71",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  2,
				MsgNum:  1,
				NumSV:   7,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 7,
						Elv:  nmea.NewOptional(79),
						Az:   nmea.NewOptional(48),
						CNO:  nmea.NewOptional(42),
					},
					{
						SVID: 2,
						Elv:  nmea.NewOptional(51),
						Az:   nmea.NewOptional(62),
						CNO:  nmea.NewOptional(43),
					},
					{
						SVID: 26,
						Elv:  nmea.NewOptional(36),
						Az:   nmea.NewOptional(256),
						CNO:  nmea.NewOptional(42),
					},
					{
						SVID: 27,
						Elv:  nmea.NewOptional(27),
						Az:   nmea.NewOptional(138),
						CNO:  nmea.NewOptional(42),
					},
				},
			},
		},
		{
			s: "$GPGSV,2,2,07,09,23,313,42,04,19,159,41,15,12,041,42*41",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  2,
				MsgNum:  2,
				NumSV:   7,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 9,
						Elv:  nmea.NewOptional(23),
						Az:   nmea.NewOptional(313),
						CNO:  nmea.NewOptional(42),
					},
					{
						SVID: 4,
						Elv:  nmea.NewOptional(19),
						Az:   nmea.NewOptional(159),
						CNO:  nmea.NewOptional(41),
					},
					{
						SVID: 15,
						Elv:  nmea.NewOptional(12),
						Az:   nmea.NewOptional(41),
						CNO:  nmea.NewOptional(42),
					},
				},
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPMSS,55,27,318.0,100,1,*57",
			expected: &MSS{
				address:            NewAddress("GPMSS"),
				SignalStrength:     55,
				SignalToNoiseRatio: 27,
				BeaconFrequency:    318000,
				BeaconBitRate:      100,
				ChannelNumber:      1,
			},
		},
		{
			s: "$GPRMC,161229.487,A,3723.2475,N,12158.3416,W,0.13,309.62,120598,,*10",
			expected: &RMC{
				address:           NewAddress("GPRMC"),
				Time:              time.Date(1998, time.May, 12, 16, 12, 29, 487000000, time.UTC),
				Status:            65,
				Lat:               37.387458333333335,
				Lon:               -121.97236,
				SpeedOverGroundKN: 0.13,
				CourseOverGround:  nmea.NewOptional(309.62),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
			},
			s: "$GPVTG,309.62,T,,M,0.13,N,0.2,K,A*23",
			expected: &VTG{
				address:              NewAddress("GPVTG"),
				TrueCourseOverGround: 309.62,
				SpeedOverGroundKN:    0.13,
				SpeedOverGroundKPH:   0.2,
				ModeIndicator:        'A',
			},
		},
		{
			s: "$GPZDA,181813,14,10,2003,00,00*4F",
			expected: &ZDA{
				address: NewAddress("GPZDA"),
				Time:    time.Date(2003, time.October, 14, 18, 18, 13, 0, time.UTC),
			},
		},

		// GNSSDO from https://ww1.microchip.com/downloads/aemDocuments/documents/VOP/ProductDocuments/ReferenceManuals/GNSSDO_NMEA_Reference_Manual_RevA.pdf
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGGA,020418.127,4048.4894,N,7720.2754,W,1,8,1.5,42.0,M,33.8,M,,*XX",
			expected: &GGA{
				address: NewAddress("GPGGA"),
				TimeOfDay: TimeOfDay{
					Hour:       2,
					Minute:     4,
					Second:     18,
					Nanosecond: 127000000,
				},
				Lat:                              40.80815666666667,
				Lon:                              -772.00459, // FIXME this is wrong. NMEA sentence is missing leading zeros.
				FixQuality:                       1,
				NumberOfSatellites:               8,
				HDOP:                             1.5,
				Alt:                              42,
				HeightOfGeoidAboveWGS84Ellipsoid: 33.8,
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGLL,4048.4894,N,7720.2754,W,020418.127,A,A*XX",
			expected: &GLL{
				address: NewAddress("GPGLL"),
				Lat:     40.80815666666667,
				Lon:     -772.00459,
				TimeOfDay: TimeOfDay{
					Hour:       2,
					Minute:     4,
					Second:     18,
					Nanosecond: 127000000,
				},
				Status:  'A',
				PosMode: 'A',
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GNGNS,020418.127,4048.4894,N,7720.2754,W,AAN,18,1.5,42.0,33.8,,,V*XX",
			expected: &GNS{
				address: NewAddress("GNGNS"),
				TimeOfDay: TimeOfDay{
					Hour:       2,
					Minute:     4,
					Second:     18,
					Nanosecond: 127000000,
				},
				Lat:       nmea.NewOptional(40.80815666666667),
				Lon:       nmea.NewOptional(-772.00459), // FIXME this is wrong
				PosMode:   []byte{'A', 'A', 'N'},
				NumSV:     18,
				HDOP:      nmea.NewOptional(1.5),
				Alt:       nmea.NewOptional(42.0),
				Sep:       nmea.NewOptional(33.8),
				NavStatus: 'V',
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GNGSA,A,3,09,15,26,05,24,21,08,02,29,28,18,10,0.8,0.5,0.5,1*XX",
			expected: &GSA{
				address: NewAddress("GNGSA"),
				OpMode:  'A',
				NavMode: 3,
				SVIDs: []nmea.Optional[int]{
					nmea.NewOptional(9),
					nmea.NewOptional(15),
					nmea.NewOptional(26),
					nmea.NewOptional(5),
					nmea.NewOptional(24),
					nmea.NewOptional(21),
					nmea.NewOptional(8),
					nmea.NewOptional(2),
					nmea.NewOptional(29),
					nmea.NewOptional(28),
					nmea.NewOptional(18),
					nmea.NewOptional(10),
				},
				PDOP:     nmea.NewOptional(0.8),
				HDOP:     nmea.NewOptional(0.5),
				VDOP:     nmea.NewOptional(0.5),
				SystemID: nmea.NewOptional(1),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGSV,4,1,14,15,67,319,52,09,63,068,53,26,45,039,50,05,44,104,49,1*XX",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  4,
				MsgNum:  1,
				NumSV:   14,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 15,
						Elv:  nmea.NewOptional(67),
						Az:   nmea.NewOptional(319),
						CNO:  nmea.NewOptional(52),
					},
					{
						SVID: 9,
						Elv:  nmea.NewOptional(63),
						Az:   nmea.NewOptional(68),
						CNO:  nmea.NewOptional(53),
					},
					{
						SVID: 26,
						Elv:  nmea.NewOptional(45),
						Az:   nmea.NewOptional(39),
						CNO:  nmea.NewOptional(50),
					},
					{
						SVID: 5,
						Elv:  nmea.NewOptional(44),
						Az:   nmea.NewOptional(104),
						CNO:  nmea.NewOptional(49),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGSV,4,2,14,24,42,196,47,21,34,302,46,18,12,305,43,28,11,067,41,1*XX",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  4,
				MsgNum:  2,
				NumSV:   14,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 24,
						Elv:  nmea.NewOptional(42),
						Az:   nmea.NewOptional(196),
						CNO:  nmea.NewOptional(47),
					},
					{
						SVID: 21,
						Elv:  nmea.NewOptional(34),
						Az:   nmea.NewOptional(302),
						CNO:  nmea.NewOptional(46),
					},
					{
						SVID: 18,
						Elv:  nmea.NewOptional(12),
						Az:   nmea.NewOptional(305),
						CNO:  nmea.NewOptional(43),
					},
					{
						SVID: 28,
						Elv:  nmea.NewOptional(11),
						Az:   nmea.NewOptional(67),
						CNO:  nmea.NewOptional(41),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGSV,4,3,14,08,07,035,38,29,04,237,39,02,02,161,40,50,47,163,44,1*XX",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  4,
				MsgNum:  3,
				NumSV:   14,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 8,
						Elv:  nmea.NewOptional(7),
						Az:   nmea.NewOptional(35),
						CNO:  nmea.NewOptional(38),
					},
					{
						SVID: 29,
						Elv:  nmea.NewOptional(4),
						Az:   nmea.NewOptional(237),
						CNO:  nmea.NewOptional(39),
					},
					{
						SVID: 2,
						Elv:  nmea.NewOptional(2),
						Az:   nmea.NewOptional(161),
						CNO:  nmea.NewOptional(40),
					},
					{
						SVID: 50,
						Elv:  nmea.NewOptional(47),
						Az:   nmea.NewOptional(163),
						CNO:  nmea.NewOptional(44),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			skip: "FIXME handle empty satellites in view",
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPGSV,4,4,14,42,48,171,44,93,65,191,48,,,,,,,,,1*XX",
			expected: &GSV{
				address: NewAddress("GPGSV"),
				NumMsg:  4,
				MsgNum:  4,
				NumSV:   14,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 42,
						Elv:  nmea.NewOptional(48),
						Az:   nmea.NewOptional(171),
						CNO:  nmea.NewOptional(44),
					},
					{
						SVID: 93,
						Elv:  nmea.NewOptional(65),
						Az:   nmea.NewOptional(191),
						CNO:  nmea.NewOptional(48),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GLGSV,3,1,09,79,66,099,50,69,55,019,53,80,33,176,46,68,28,088,45,1*XX",
			expected: &GSV{
				address: NewAddress("GLGSV"),
				NumMsg:  3,
				MsgNum:  1,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 79,
						Elv:  nmea.NewOptional(66),
						Az:   nmea.NewOptional(99),
						CNO:  nmea.NewOptional(50),
					},
					{
						SVID: 69,
						Elv:  nmea.NewOptional(55),
						Az:   nmea.NewOptional(19),
						CNO:  nmea.NewOptional(53),
					},
					{
						SVID: 80,
						Elv:  nmea.NewOptional(33),
						Az:   nmea.NewOptional(176),
						CNO:  nmea.NewOptional(46),
					},
					{
						SVID: 68,
						Elv:  nmea.NewOptional(28),
						Az:   nmea.NewOptional(88),
						CNO:  nmea.NewOptional(45),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GLGSV,3,2,09,70,25,315,46,78,24,031,42,85,18,293,44,84,16,246,41,1*XX",
			expected: &GSV{
				address: NewAddress("GLGSV"),
				NumMsg:  3,
				MsgNum:  2,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 70,
						Elv:  nmea.NewOptional(25),
						Az:   nmea.NewOptional(315),
						CNO:  nmea.NewOptional(46),
					},
					{
						SVID: 78,
						Elv:  nmea.NewOptional(24),
						Az:   nmea.NewOptional(31),
						CNO:  nmea.NewOptional(42),
					},
					{
						SVID: 85,
						Elv:  nmea.NewOptional(18),
						Az:   nmea.NewOptional(293),
						CNO:  nmea.NewOptional(44),
					},
					{
						SVID: 84,
						Elv:  nmea.NewOptional(16),
						Az:   nmea.NewOptional(246),
						CNO:  nmea.NewOptional(41),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			skip: "FIXME handle empty satellites in view",
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GLGSV,3,3,09,86,02,338,,,,,,,,,,,,,,1*XX",
			expected: &GSV{
				address: NewAddress("GLGSV"),
				NumMsg:  3,
				MsgNum:  3,
				NumSV:   9,
				SatellitesInView: []SatelliteInView{
					{
						SVID: 86,
						Elv:  nmea.NewOptional(2),
						Az:   nmea.NewOptional(338),
						CNO:  nmea.NewOptional(46),
					},
				},
				SignalID: nmea.NewOptional(1),
			},
		},
		{
			skip: "FIXME handle missing status",
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GNRMC,020418.127,4048.4894,N,7720.2754,W,0.00,0.00,180116,,,A,V*XX",
			expected: &RMC{
				address: NewAddress("GNRMC"),
			},
		},
		{
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GNVTG,0.00,T,,M,0.00,N,0.00,K,D*XX",
			expected: &VTG{
				address:       NewAddress("GNVTG"),
				ModeIndicator: 'D',
			},
		},
		{
			skip: "FIXME handle + sign in hours",
			options: []nmea.ParserOption{
				nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			},
			s: "$GPZDA,014811.000,13,09,2013,+00,00*XX",
			expected: &ZDA{
				address: NewAddress("GPZDA"),
				Time:    time.Date(2013, time.September, 13, 1, 48, 11, 0, time.UTC),
			},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			if tc.skip != "" {
				t.Skip(tc.skip)
			}
			actual, err := newParser(tc.options...).ParseString(tc.s)
			if tc.expectedErr != nil {
				assert.IsError(t, err, tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}

}

func newParser(options ...nmea.ParserOption) *nmea.Parser {
	options = append([]nmea.ParserOption{
		nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
		nmea.WithSentenceParserFunc(SentenceParser),
	}, options...)
	return nmea.NewParser(options...)
}
