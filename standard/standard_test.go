package standard_test

import (
	"testing"
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeatest"
	"github.com/twpayne/go-nmea/standard"
)

func TestUblox(t *testing.T) {
	// From https://content.u-blox.com/sites/default/files/products/documents/u-blox8-M8_ReceiverDescrProtSpec_UBX-13003221.pdf.
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F",
				Expected: &standard.DTM{
					Address:  nmea.NewAddress("GPDTM"),
					Datum:    "W84",
					RefDatum: "W84",
				},
			},
			{
				S: "$GPGBS,235503.00,1.6,1.4,3.2,,,,,,*40",
				Expected: &standard.GBS{
					Address: nmea.NewAddress("GPGBS"),
					TimeOfDay: nmea.TimeOfDay{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPGBS,235458.00,1.4,1.3,3.1,03,,-21.4,3.8,1,0*5B",
				Expected: &standard.GBS{
					Address: nmea.NewAddress("GPGBS"),
					TimeOfDay: nmea.TimeOfDay{
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
				S: "$GPGGA,092725.00,4717.11399,N,00833.91590,E,1,08,1.01,499.6,M,48.0,M,,*5B",
				Expected: &standard.GGA{
					Address: nmea.NewAddress("GPGGA"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   9,
						Minute: 27,
						Second: 25,
					},
					Lat:                              nmea.NewOptional(47.285233166666664),
					Lon:                              nmea.NewOptional(8.565265),
					FixQuality:                       1,
					NumberOfSatellites:               8,
					HDOP:                             1.01,
					Alt:                              nmea.NewOptional(499.6),
					HeightOfGeoidAboveWGS84Ellipsoid: nmea.NewOptional(48.0),
				},
			},
			{
				S: "$GPGLL,4717.11364,N,00833.91565,E,092321.00,A,A*60",
				Expected: &standard.GLL{
					Address: nmea.NewAddress("GPGLL"),
					Lat:     nmea.NewOptional(47.28522733333333),
					Lon:     nmea.NewOptional(8.565260833333333),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   9,
						Minute: 23,
						Second: 21,
					},
					Status:  'A',
					PosMode: 'A',
				},
			},
			{
				S: "$GNGNS,103600.01,5114.51176,N,00012.29380,W,ANNN,07,1.18,111.5,45.6,,,V*00",
				Expected: &standard.GNS{
					Address: nmea.NewAddress("GNGNS"),
					TimeOfDay: nmea.TimeOfDay{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GNGNS,122310.2,3722.425671,N,12258.856215,W,DAAA,14,0.9,1005.543,6.5,,,V*0E",
				Expected: &standard.GNS{
					Address: nmea.NewAddress("GNGNS"),
					TimeOfDay: nmea.TimeOfDay{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPGNS,122310.2,,,,,,07,,,,5.2,23,V*02",
				Expected: &standard.GNS{
					Address: nmea.NewAddress("GPGNS"),
					TimeOfDay: nmea.TimeOfDay{
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
				S: "$GNGRS,104148.00,1,2.6,2.2,-1.6,-1.1,-1.7,-1.5,5.8,1.7,,,,,1,1*52",
				Expected: &standard.GRS{
					Address: nmea.NewAddress("GNGRS"),
					TimeOfDay: nmea.TimeOfDay{
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
					SystemID: nmea.NewOptional(1),
					SignalID: nmea.NewOptional(1),
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GNGRS,104148.00,1,,0.0,2.5,0.0,,2.8,,,,,,,1,5*52",
				Expected: &standard.GRS{
					Address: nmea.NewAddress("GNGRS"),
					TimeOfDay: nmea.TimeOfDay{
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
					SystemID: nmea.NewOptional(1),
					SignalID: nmea.NewOptional(5),
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPGSA,A,3,23,29,07,08,09,18,26,28,,,,,1.94,1.18,1.54,1*0D",
				Expected: &standard.GSA{
					Address: nmea.NewAddress("GPGSA"),
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
				S: "$GPGST,082356.00,1.8,,,,1.7,1.3,2.2*7E",
				Expected: &standard.GST{
					Address: nmea.NewAddress("GPGST"),
					TimeOfDay: nmea.TimeOfDay{
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
				S: "$GPGSV,3,1,09,09,,,17,10,,,40,12,,,49,13,,,35,1*6F",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  3,
					MsgNum:  1,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
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
				S: "$GPGSV,3,2,09,15,,,44,17,,,45,19,,,44,24,,,50,1*64",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  3,
					MsgNum:  2,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
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
				S: "$GPGSV,3,3,09,25,,,40,1*6E",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  3,
					MsgNum:  3,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
						{
							SVID: 25,
							CNO:  nmea.NewOptional(40),
						},
					},
					SignalID: nmea.NewOptional(1),
				},
			},
			{
				S: "$GPGSV,1,1,03,12,,,42,24,,,47,32,,,37,5*66",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  1,
					MsgNum:  1,
					NumSV:   3,
					SatellitesInView: []standard.SatelliteInView{
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
				S: "$GAGSV,1,1,00,2*76",
				Expected: &standard.GSV{
					Address:  nmea.NewAddress("GAGSV"),
					NumMsg:   1,
					MsgNum:   1,
					NumSV:    0,
					SignalID: nmea.NewOptional(2),
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPRMC,083559.00,A,4717.11437,N,00833.91522,E,0.004,77.52,091202,,,A,V*57",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GPRMC"),
					Time:              time.Date(2002, time.December, 9, 8, 35, 59, 0, time.UTC),
					Status:            'A',
					Lat:               nmea.NewOptional(47.2852395),
					Lon:               nmea.NewOptional(8.565253666666667),
					SpeedOverGroundKN: nmea.NewOptional(0.004),
					CourseOverGround:  nmea.NewOptional(77.52),
					ModeIndicator:     'A',
					NavStatus:         'V',
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPTHS,77.52,E*32",
				Expected: &standard.THS{
					Address:       nmea.NewAddress("GPTHS"),
					TrueHeading:   77.52,
					ModeIndicator: 'E',
				},
			},
			{
				S: "$GPTXT,01,01,02,u-blox ag - www.u-blox.com*50",
				Expected: &standard.TXT{
					Address: nmea.NewAddress("GPTXT"),
					NumMsg:  1,
					MsgNum:  1,
					MsgType: 2,
					Text:    "u-blox ag - www.u-blox.com",
				},
			},
			{
				S: "$GPTXT,01,01,02,ANTARIS ATR0620 HW 00000040*67",
				Expected: &standard.TXT{
					Address: nmea.NewAddress("GPTXT"),
					NumMsg:  1,
					MsgNum:  1,
					MsgType: 2,
					Text:    "ANTARIS ATR0620 HW 00000040",
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPVLW,,N,,N,15.8,N,1.2,N*06",
				Expected: &standard.VLW{
					Address:               nmea.NewAddress("GPVLW"),
					TotalGroundDistanceNM: nmea.NewOptional(15.8),
					GroundDistanceNM:      nmea.NewOptional(1.2),
				},
			},
			{
				S: "$GPVTG,77.52,T,,M,0.004,N,0.008,K,A*06",
				Expected: &standard.VTG{
					Address:              nmea.NewAddress("GPVTG"),
					TrueCourseOverGround: nmea.NewOptional(77.52),
					SpeedOverGroundKN:    nmea.NewOptional(0.004),
					SpeedOverGroundKPH:   nmea.NewOptional(0.008),
					ModeIndicator:        'A',
				},
			},
			{
				S: "$GPZDA,082710.00,16,09,2002,00,00*64",
				Expected: &standard.ZDA{
					Address:              nmea.NewAddress("GPZDA"),
					Time:                 time.Date(2002, time.September, 16, 8, 27, 10, 0, time.UTC),
					LocalTimeZoneHours:   nmea.NewOptional(0),
					LocalTimeZoneMinutes: nmea.NewOptional(0),
				},
			},
		},
	)
}

func TestSparkfun(t *testing.T) {
	// From https://www.sparkfun.com/datasheets/GPS/NMEA%20Reference%20Manual-Rev2.1-Dec07.pdf.
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$GPGGA,002153.000,3342.6618,N,11751.3858,W,1,10,1.2,27.0,M,-34.2,M,,0000*5E",
				Expected: &standard.GGA{
					Address: nmea.NewAddress("GPGGA"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   0,
						Minute: 21,
						Second: 53,
					},
					Lat:                              nmea.NewOptional(33.71103),
					Lon:                              nmea.NewOptional(-117.85643),
					FixQuality:                       1,
					NumberOfSatellites:               10,
					HDOP:                             1.2,
					Alt:                              nmea.NewOptional(27.0),
					HeightOfGeoidAboveWGS84Ellipsoid: nmea.NewOptional(-34.2),
					DGPSReferenceStationID:           "0000",
				},
			},
			{
				S: "$GPGLL,3723.2475,N,12158.3416,W,161229.487,A,A*41",
				Expected: &standard.GLL{
					Address: nmea.NewAddress("GPGLL"),
					Lat:     nmea.NewOptional(37.387458333333335),
					Lon:     nmea.NewOptional(-121.97236),
					TimeOfDay: nmea.TimeOfDay{
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
				S: "$GPGSA,A,3,07,02,26,27,09,04,15,,,,,,1.8,1.0,1.5*33",
				Expected: &standard.GSA{
					Address: nmea.NewAddress("GPGSA"),
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
				S: "$GPGSV,2,1,07,07,79,048,42,02,51,062,43,26,36,256,42,27,27,138,42*71",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  2,
					MsgNum:  1,
					NumSV:   7,
					SatellitesInView: []standard.SatelliteInView{
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
				S: "$GPGSV,2,2,07,09,23,313,42,04,19,159,41,15,12,041,42*41",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  2,
					MsgNum:  2,
					NumSV:   7,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPMSS,55,27,318.0,100,1*57",
				Expected: &standard.MSS{
					Address:            nmea.NewAddress("GPMSS"),
					SignalStrength:     55,
					SignalToNoiseRatio: 27,
					BeaconFrequencyKHz: 318,
					BeaconBitRate:      100,
					ChannelNumber:      nmea.NewOptional(1),
				},
			},
			{
				S: "$GPRMC,161229.487,A,3723.2475,N,12158.3416,W,0.13,309.62,120598,,*10",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GPRMC"),
					Time:              time.Date(1998, time.May, 12, 16, 12, 29, 487000000, time.UTC),
					Status:            'A',
					Lat:               nmea.NewOptional(37.387458333333335),
					Lon:               nmea.NewOptional(-121.97236),
					SpeedOverGroundKN: nmea.NewOptional(0.13),
					CourseOverGround:  nmea.NewOptional(309.62),
				},
			},
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineRequire),
				},
				S: "$GPVTG,309.62,T,,M,0.13,N,0.2,K,A*23",
				Expected: &standard.VTG{
					Address:              nmea.NewAddress("GPVTG"),
					TrueCourseOverGround: nmea.NewOptional(309.62),
					SpeedOverGroundKN:    nmea.NewOptional(0.13),
					SpeedOverGroundKPH:   nmea.NewOptional(0.2),
					ModeIndicator:        'A',
				},
			},
			{
				S: "$GPZDA,181813,14,10,2003,00,00*4F",
				Expected: &standard.ZDA{
					Address:              nmea.NewAddress("GPZDA"),
					Time:                 time.Date(2003, time.October, 14, 18, 18, 13, 0, time.UTC),
					LocalTimeZoneHours:   nmea.NewOptional(0),
					LocalTimeZoneMinutes: nmea.NewOptional(0),
				},
			},
		},
	)
}

func TestGNSSDO(t *testing.T) {
	// From https://ww1.microchip.com/downloads/aemDocuments/documents/VOP/ProductDocuments/ReferenceManuals/GNSSDO_NMEA_Reference_Manual_RevA.pdf.
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GNGSA,A,3,09,15,26,05,24,21,08,02,29,28,18,10,0.8,0.5,0.5,1*XX",
				Expected: &standard.GSA{
					Address: nmea.NewAddress("GNGSA"),
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GPGSV,4,1,14,15,67,319,52,09,63,068,53,26,45,039,50,05,44,104,49,1*XX",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  4,
					MsgNum:  1,
					NumSV:   14,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GPGSV,4,2,14,24,42,196,47,21,34,302,46,18,12,305,43,28,11,067,41,1*XX",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  4,
					MsgNum:  2,
					NumSV:   14,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GPGSV,4,3,14,08,07,035,38,29,04,237,39,02,02,161,40,50,47,163,44,1*XX",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  4,
					MsgNum:  3,
					NumSV:   14,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GLGSV,3,1,09,79,66,099,50,69,55,019,53,80,33,176,46,68,28,088,45,1*XX",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GLGSV"),
					NumMsg:  3,
					MsgNum:  1,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GLGSV,3,2,09,70,25,315,46,78,24,031,42,85,18,293,44,84,16,246,41,1*XX",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GLGSV"),
					NumMsg:  3,
					MsgNum:  2,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
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
				Options: []nmea.ParserOption{
					nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
				},
				S: "$GNVTG,0.00,T,,M,0.00,N,0.00,K,D*XX",
				Expected: &standard.VTG{
					Address:              nmea.NewAddress("GNVTG"),
					TrueCourseOverGround: nmea.NewOptional(0.0),
					SpeedOverGroundKN:    nmea.NewOptional(0.0),
					SpeedOverGroundKPH:   nmea.NewOptional(0.0),
					ModeIndicator:        'D',
				},
			},
		},
	)
}

func TestMiscellaneous(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$GPGGA,102039.00,,,,,0,00,99.99,,,,,,*6F",
				Expected: &standard.GGA{
					Address: nmea.NewAddress("GPGGA"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   10,
						Minute: 20,
						Second: 39,
					},
					HDOP: 99.99,
				},
			},
			{
				S: "$GNGLL,,,,,095943.00,V,N*56",
				Expected: &standard.GLL{
					Address: nmea.NewAddress("GNGLL"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   9,
						Minute: 59,
						Second: 43,
					},
					Status:  'V',
					PosMode: 'N',
				},
			},
			{
				S: "$GPMSS,0,0,0.000000,0,*58",
				Expected: &standard.MSS{
					Address: nmea.NewAddress("GPMSS"),
				},
			},
			{
				S: "$GPMSS,0,0,0.000000,200,*5A",
				Expected: &standard.MSS{
					Address:       nmea.NewAddress("GPMSS"),
					BeaconBitRate: 200,
				},
			},
			{
				S: "$GPMSS,55,27,318.0,100,*66",
				Expected: &standard.MSS{
					Address:            nmea.NewAddress("GPMSS"),
					SignalStrength:     55,
					SignalToNoiseRatio: 27,
					BeaconFrequencyKHz: 318,
					BeaconBitRate:      100,
				},
			},
			{
				S: "$GPRMC,102042.00,V,,,,,,,110324,,,N*7D",
				Expected: &standard.RMC{
					Address:       nmea.NewAddress("GPRMC"),
					Time:          time.Date(2024, time.March, 11, 10, 20, 42, 0, time.UTC),
					Status:        'V',
					ModeIndicator: 'N',
				},
			},
			{
				S: "$GPRMC,085131,A,4652.8560,N,00821.9500,E,14.8,90.0,070424,,,A*76",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GPRMC"),
					Time:              time.Date(2024, time.April, 7, 8, 51, 31, 0, time.UTC),
					Status:            65,
					Lat:               nmea.NewOptional(46.88093333333333),
					Lon:               nmea.NewOptional(8.365833333333333),
					SpeedOverGroundKN: nmea.NewOptional(14.8),
					CourseOverGround:  nmea.NewOptional(90.0),
					ModeIndicator:     'A',
				},
			},
			{
				S: "$GPGSV,2,2,07,32,-1,222,08,36,34,160,,49,36,185,*58",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GPGSV"),
					NumMsg:  2,
					MsgNum:  2,
					NumSV:   7,
					SatellitesInView: []standard.SatelliteInView{
						{
							SVID: 32,
							Elv:  nmea.NewOptional(-1),
							Az:   nmea.NewOptional(222),
							CNO:  nmea.NewOptional(8),
						},
						{
							SVID: 36,
							Elv:  nmea.NewOptional(34),
							Az:   nmea.NewOptional(160),
						},
						{
							SVID: 49,
							Elv:  nmea.NewOptional(36),
							Az:   nmea.NewOptional(185),
						},
					},
				},
			},
		},
	)
}

func TestTheNMEA0813InformationSheetIssue4(t *testing.T) {
	// From https://actisense.com/wp-content/uploads/2020/01/NMEA-0183-Information-sheet-issue-4-1-1.pdf.
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineIgnore),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$SDDBT,8.1,f,2.4,M,1.3,F*0B",
				Expected: &standard.DBT{
					Address:      nmea.NewAddress("SDDBT"),
					DepthFeet:    8.1,
					Depth:        2.4,
					DepthFathoms: 1.3,
				},
			},
			{
				S: "$SDDPT,76.1,0.0,100*00",
				Expected: &standard.DPT{
					Address: nmea.NewAddress("SDDPT"),
					Depth:   76.1,
					Offset:  nmea.NewOptional(0.0),
					Maximum: nmea.NewOptional(100.0),
				},
			},
			{
				S: "$SDDPT,2.4,,*7F",
				Expected: &standard.DPT{
					Address: nmea.NewAddress("SDDPT"),
					Depth:   2.4,
				},
			},
			{
				S: "$YXMTW,17.75,C*5D",
				Expected: &standard.MTW{
					Address:     nmea.NewAddress("YXMTW"),
					Temperature: 17.75,
				},
			},
			{
				S: "$VWVHW,,,,,0.0,N,0.0,K*4D",
				Expected: &standard.VHW{
					Address:    nmea.NewAddress("VWVHW"),
					SpeedKnots: nmea.NewOptional(0.0),
					SpeedKPH:   nmea.NewOptional(0.0),
				},
			},
			{
				S: "$VWVLW,2.8,N,2.8,N*4C",
				Expected: &standard.VLW{
					Address:              nmea.NewAddress("VWVLW"),
					TotalWaterDistanceNM: nmea.NewOptional(2.8),
					WaterDistanceNM:      nmea.NewOptional(2.8),
				},
			},
		},
	)
}

func TestNovatel(t *testing.T) {
	// From https://docs.novatel.com/OEM7/Content/Logs/Core_Logs.htm.
	// FIXME add more test cases
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(standard.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$GPALM,30,01,01,2210,00,617b,0f,1da7,fd70,a10d0a,24de91,6fe696,16263f,17c,ffe*7B",
				Expected: &standard.ALM{
					Address:                nmea.NewAddress("GPALM"),
					NumMsg:                 30,
					MsgNum:                 1,
					PRN:                    1,
					GPSWeek:                2210,
					SVHealth:               0,
					Eccentricity:           0x617b,
					AlmanacReferenceTime:   0x0f,
					InclinationAngle:       0x1da7,
					OmegaDot:               0xfd70,
					RootAxis:               0xa10d0a,
					Omega:                  0x24de91,
					AscensionNodeLongitude: 0x6fe696,
					MeanAnomaly:            0x16263f,
					AF0:                    0x17c,
					AF1:                    0xffe,
				},
			},
			{
				S: "$GPGGA,202530.00,5109.0262,N,11401.8407,W,5,40,0.5,1097.36,M,-17.00,M,18,TSTR*61",
				Expected: &standard.GGA{
					Address: nmea.NewAddress("GPGGA"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   20,
						Minute: 25,
						Second: 30,
					},
					Lat:                              nmea.NewOptional(51.150436666666664),
					Lon:                              nmea.NewOptional(-114.03067833333333),
					FixQuality:                       5,
					NumberOfSatellites:               40,
					HDOP:                             0.5,
					Alt:                              nmea.NewOptional(1097.36),
					HeightOfGeoidAboveWGS84Ellipsoid: nmea.NewOptional(-17.0),
					TimeSinceLastDGPSUpdate:          nmea.NewOptional(18),
					DGPSReferenceStationID:           "TSTR",
				},
			},
			{
				S: "$GPGLL,5109.0262317,N,11401.8407304,W,202725.00,A,D*79",
				Expected: &standard.GLL{
					Address: nmea.NewAddress("GPGLL"),
					Lat:     nmea.NewOptional(51.150437195),
					Lon:     nmea.NewOptional(-114.03067884),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   20,
						Minute: 27,
						Second: 25,
					},
					Status:  'A',
					PosMode: 'D',
				},
			},
			{
				S: "$GNGLL,5109.0262321,N,11401.8407167,W,174738.00,A,D*6B",
				Expected: &standard.GLL{
					Address: nmea.NewAddress("GNGLL"),
					Lat:     nmea.NewOptional(51.15043720166667),
					Lon:     nmea.NewOptional(-114.03067861166667),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   17,
						Minute: 47,
						Second: 38,
					},
					Status:  'A',
					PosMode: 'D',
				},
			},
			{
				S: "$GNGRS,174837.00,1,-0.1,0.7,-0.2,0.1,0.3,0.5,-0.7,-0.5,-0.3,0.3,,*72",
				Expected: &standard.GRS{
					Address: nmea.NewAddress("GNGRS"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   17,
						Minute: 48,
						Second: 37,
					},
					Mode: 1,
					Residuals: []nmea.Optional[float64]{
						nmea.NewOptional(-0.1),
						nmea.NewOptional(0.7),
						nmea.NewOptional(-0.2),
						nmea.NewOptional(0.1),
						nmea.NewOptional(0.3),
						nmea.NewOptional(0.5),
						nmea.NewOptional(-0.7),
						nmea.NewOptional(-0.5),
						nmea.NewOptional(-0.3),
						nmea.NewOptional(0.3),
						{},
						{},
					},
				},
			},
			{
				S: "$GPRMC,203522.00,A,5109.0262308,N,11401.8407342,W,0.004,133.4,130522,0.0,E,D*2B",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GPRMC"),
					Time:              time.Date(2022, time.May, 13, 20, 35, 22, 0, time.UTC),
					Status:            'A',
					Lat:               nmea.NewOptional(51.15043718),
					Lon:               nmea.NewOptional(-114.03067890333334),
					SpeedOverGroundKN: nmea.NewOptional(0.004),
					CourseOverGround:  nmea.NewOptional(133.4),
					MagneticVariation: nmea.NewOptional(0.0),
					ModeIndicator:     'D',
				},
			},
			{
				S: "$GNRMC,204520.00,A,5109.0262239,N,11401.8407338,W,0.004,102.3,130522,0.0,E,D*3B",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GNRMC"),
					Time:              time.Date(2022, time.May, 13, 20, 45, 20, 0, time.UTC),
					Status:            'A',
					Lat:               nmea.NewOptional(51.150437065),
					Lon:               nmea.NewOptional(-114.03067889666667),
					SpeedOverGroundKN: nmea.NewOptional(0.004),
					CourseOverGround:  nmea.NewOptional(102.3),
					MagneticVariation: nmea.NewOptional(0.0),
					ModeIndicator:     'D',
				},
			},
			{
				S: "$GPGSA,M,3,05,02,31,06,19,29,20,12,24,25,,,0.9,0.5,0.7*35",
				Expected: &standard.GSA{
					Address: nmea.NewAddress("GPGSA"),
					OpMode:  'M',
					NavMode: 3,
					SVIDs: []nmea.Optional[int]{
						nmea.NewOptional(5),
						nmea.NewOptional(2),
						nmea.NewOptional(31),
						nmea.NewOptional(6),
						nmea.NewOptional(19),
						nmea.NewOptional(29),
						nmea.NewOptional(20),
						nmea.NewOptional(12),
						nmea.NewOptional(24),
						nmea.NewOptional(25),
						{},
						{},
					},
					PDOP: nmea.NewOptional(0.9),
					HDOP: nmea.NewOptional(0.5),
					VDOP: nmea.NewOptional(0.7),
				},
			},
			{
				S: "$GPGST,203017.00,1.25,0.02,0.01,-16.7566,0.02,0.01,0.03*7D",
				Expected: &standard.GST{
					Address: nmea.NewAddress("GPGST"),
					TimeOfDay: nmea.TimeOfDay{
						Hour:   20,
						Minute: 30,
						Second: 17,
					},
					RangeRMS:    1.25,
					MajorStdDev: nmea.NewOptional(0.02),
					MinorStdDev: nmea.NewOptional(0.01),
					Orientation: nmea.NewOptional(-16.7566),
					LatStdDev:   0.02,
					LonStdDev:   0.01,
					AltStdDev:   0.03,
				},
			},
			{
				S: "$GAGSV,3,1,09,34,72,231,53,30,65,251,53,36,51,059,51,02,36,170,49*62",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GAGSV"),
					NumMsg:  3,
					MsgNum:  1,
					NumSV:   9,
					SatellitesInView: []standard.SatelliteInView{
						{
							SVID: 34,
							Elv:  nmea.NewOptional(72),
							Az:   nmea.NewOptional(231),
							CNO:  nmea.NewOptional(53),
						},
						{
							SVID: 30,
							Elv:  nmea.NewOptional(65),
							Az:   nmea.NewOptional(251),
							CNO:  nmea.NewOptional(53),
						},
						{
							SVID: 36,
							Elv:  nmea.NewOptional(51),
							Az:   nmea.NewOptional(59),
							CNO:  nmea.NewOptional(51),
						},
						{
							SVID: 2,
							Elv:  nmea.NewOptional(36),
							Az:   nmea.NewOptional(170),
							CNO:  nmea.NewOptional(49),
						},
					},
				},
			},
			{
				S: "$GQGSV,1,1,01,02,08,309,37*4D",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GQGSV"),
					NumMsg:  1,
					MsgNum:  1,
					NumSV:   1,
					SatellitesInView: []standard.SatelliteInView{
						{
							SVID: 2,
							Elv:  nmea.NewOptional(8),
							Az:   nmea.NewOptional(309),
							CNO:  nmea.NewOptional(37),
						},
					},
				},
			},
			{
				S: "$GIGSV,1,1,00,,,,*60",
				Expected: &standard.GSV{
					Address: nmea.NewAddress("GIGSV"),
					NumMsg:  1,
					MsgNum:  1,
					NumSV:   0,
				},
			},
			{
				S: "$GPHDT,75.5664,T*36",
				Expected: &standard.HDT{
					Address:     nmea.NewAddress("GPHDT"),
					HeadingTrue: 75.5664,
				},
			},
			{
				S: "$GPRMB,A,4.32,L,FROM,TO,5109.7578000,N,11409.0960000,W,4.6,279.2,0.0,V,D*4A",
				Expected: &standard.RMB{
					Address:                nmea.NewAddress("GPRMB"),
					DataStatus:             'A',
					CrossTrackError:        4.32,
					CrossTrackErrorDir:     76,
					OriginWaypointID:       "FROM",
					DestinationWaypointID:  "TO",
					DestinationWaypointLat: 51.16263,
					DestinationWaypointLon: -114.1516,
					RangeNM:                4.6,
					BearingTrue:            279.2,
					ArrivalStatus:          'V',
					ModeIndicator:          'D',
				},
			},
			{
				S: "$GPRMC,203522.00,A,5109.0262308,N,11401.8407342,W,0.004,133.4,130522,0.0,E,D*2B",
				Expected: &standard.RMC{
					Address:           nmea.NewAddress("GPRMC"),
					Time:              time.Date(2022, time.May, 13, 20, 35, 22, 0, time.UTC),
					Status:            'A',
					Lat:               nmea.NewOptional(51.15043718),
					Lon:               nmea.NewOptional(-114.03067890333334),
					SpeedOverGroundKN: nmea.NewOptional(0.004),
					CourseOverGround:  nmea.NewOptional(133.4),
					MagneticVariation: nmea.NewOptional(0.0),
					ModeIndicator:     'D',
				},
			},
			{
				S: "$GNVTG,139.969,T,139.969,M,0.007,N,0.013,K,D*3D",
				Expected: &standard.VTG{
					Address:                  nmea.NewAddress("GNVTG"),
					TrueCourseOverGround:     nmea.NewOptional(139.969),
					MagneticCourseOverGround: nmea.NewOptional(139.969),
					SpeedOverGroundKN:        nmea.NewOptional(0.007),
					SpeedOverGroundKPH:       nmea.NewOptional(0.013),
					ModeIndicator:            'D',
				},
			},
			{
				S: "$GPZDA,204007.00,13,05,2022,,*62",
				Expected: &standard.ZDA{
					Address: nmea.NewAddress("GPZDA"),
					Time:    time.Date(2022, time.May, 13, 20, 40, 7, 0, time.UTC),
				},
			},
			{
				S: "$GNVTG,,,,,,,,,N*2E",
				Expected: &standard.VTG{
					Address:       nmea.NewAddress("GNVTG"),
					ModeIndicator: 'N',
				},
			},
		},
	)
}
