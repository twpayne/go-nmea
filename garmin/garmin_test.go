package garmin_test

import (
	"testing"
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/garmin"
	"github.com/twpayne/go-nmea/nmeatest"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(garmin.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$PGRMB,0.0,200,,,,K,,N,N*31",
				Expected: &garmin.PGRMB{
					Address:       nmea.NewAddress("PGRMB"),
					BeaconBitRate: 200,
					DGPSFixSource: 'N',
					DGPSMode:      'N',
				},
			},
			{
				S: "$PGRME,4.4,M,5.5,M,7.1,M*28",
				Expected: &garmin.PGRME{
					Address:                 nmea.NewAddress("PGRME"),
					HorizontalPositionError: 4.4,
					VerticalPositionError:   5.5,
					PositionError:           7.1,
				},
			},
			{
				S: "$PGRMF,290,293895,160305,093802,13,5213.1439,N,02100.6511,E,A,2,0,226,2,1*11",
				Expected: &garmin.PGRMF{
					Address:          nmea.NewAddress("PGRMF"),
					GPSWeekNumber:    290,
					GPSSeconds:       293895,
					Time:             time.Date(2005, time.March, 16, 9, 38, 2, 0, time.UTC),
					LeapSeconds:      13,
					Lat:              52.219065,
					Lon:              21.010851666666667,
					Mode:             'A',
					FixType:          2,
					CourseOverGround: 226,
					PDOP:             2,
					TDOP:             1,
				},
			},
			{
				S: "$PGRMM,WGS 84*06",
				Expected: &garmin.PGRMM{
					Address: nmea.NewAddress("PGRMM"),
					Datum:   "WGS 84",
				},
			},
			{
				S: "$PGRMT,GPS 17-HVS Ver. 2.80,P,P,R,R,P,,37,R*0F",
				Expected: &garmin.PGRMT{
					Address:                        nmea.NewAddress("PGRMT"),
					ProductModelAndSoftwareVersion: "GPS 17-HVS Ver. 2.80",
					ROMChecksumTest:                nmea.NewOptional[byte]('P'),
					ReceiverFailureDiscrete:        nmea.NewOptional[byte]('P'),
					StoredDataLost:                 nmea.NewOptional[byte]('R'),
					RealTimeClockLost:              nmea.NewOptional[byte]('R'),
					OscillatorDriftDiscrete:        nmea.NewOptional[byte]('P'),
					GPSSensorTemperature:           nmea.NewOptional(37),
					GPSSensorConfigurationData:     nmea.NewOptional[byte]('R'),
				},
			},
			{
				S: "$PGRMT,GPS17x Software Version 2.30,,,,,,,,*00",
				Expected: &garmin.PGRMT{
					Address:                        nmea.NewAddress("PGRMT"),
					ProductModelAndSoftwareVersion: "GPS17x Software Version 2.30",
				},
			},
			{
				S: "$PGRMT,GPS 25-LVS VER 2.50 ,P,P,R,R,P,,27,R*08",
				Expected: &garmin.PGRMT{
					Address:                        nmea.NewAddress("PGRMT"),
					ProductModelAndSoftwareVersion: "GPS 25-LVS VER 2.50 ",
					ROMChecksumTest:                nmea.NewOptional[byte]('P'),
					ReceiverFailureDiscrete:        nmea.NewOptional[byte]('P'),
					StoredDataLost:                 nmea.NewOptional[byte]('R'),
					RealTimeClockLost:              nmea.NewOptional[byte]('R'),
					OscillatorDriftDiscrete:        nmea.NewOptional[byte]('P'),
					GPSSensorTemperature:           nmea.NewOptional(27),
					GPSSensorConfigurationData:     nmea.NewOptional[byte]('R'),
				},
			},
			{
				S: "$PGRMV,-2.5,-1.1,0.3*58",
				Expected: &garmin.PGRMV{
					Address:           nmea.NewAddress("PGRMV"),
					TrueEastVelocity:  -2.5,
					TrueNorthVelocity: -1.1,
					UpVelocity:        0.3,
				},
			},
			{
				S: "$PGRMZ,5584,F,2*06",
				Expected: &garmin.PGRMZ{
					Address: nmea.NewAddress("PGRMZ"),
					AltFeet: 5584,
					FixType: 2,
				},
			},
			{
				S: "$PGRMZ,2062,f,3*2D",
				Expected: &garmin.PGRMZ{
					Address: nmea.NewAddress("PGRMZ"),
					AltFeet: 2062,
					FixType: 3,
				},
			},
		})
}
