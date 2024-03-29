package garmin

import (
	"testing"
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeatesting"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatesting.TestSentenceParserFunc(t, SentenceParserFunc, []nmeatesting.TestCase{
		{
			S: "$PGRMB,0.0,200,,,,K,,N,N*31",
			Expected: &PGRMB{
				address:       nmea.NewAddress("PGRMB"),
				BeaconBitRate: 200,
				DGPSFixSource: 'N',
				DGPSMode:      'N',
			},
		},
		{
			S: "$PGRME,4.4,M,5.5,M,7.1,M*28",
			Expected: &PGRME{
				address:                 nmea.NewAddress("PGRME"),
				HorizontalPositionError: 4.4,
				VerticalPositionError:   5.5,
				PositionError:           7.1,
			},
		},
		{
			S: "$PGRMF,290,293895,160305,093802,13,5213.1439,N,02100.6511,E,A,2,0,226,2,1*11",
			Expected: &PGRMF{
				address:          nmea.NewAddress("PGRMF"),
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
			Expected: &PGRMM{
				address: nmea.NewAddress("PGRMM"),
				Datum:   "WGS 84",
			},
		},
		{
			S: "$PGRMT,GPS 17-HVS Ver. 2.80,P,P,R,R,P,,37,R*0F",
			Expected: &PGRMT{
				address:                        nmea.NewAddress("PGRMT"),
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
			Expected: &PGRMT{
				address:                        nmea.NewAddress("PGRMT"),
				ProductModelAndSoftwareVersion: "GPS17x Software Version 2.30",
			},
		},
		{
			S: "$PGRMT,GPS 25-LVS VER 2.50 ,P,P,R,R,P,,27,R*08",
			Expected: &PGRMT{
				address:                        nmea.NewAddress("PGRMT"),
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
			Expected: &PGRMV{
				address:           nmea.NewAddress("PGRMV"),
				TrueEastVelocity:  -2.5,
				TrueNorthVelocity: -1.1,
				UpVelocity:        0.3,
			},
		},
		{
			S: "$PGRMZ,5584,F,2*06",
			Expected: &PGRMZ{
				address: nmea.NewAddress("PGRMZ"),
				AltFeet: 5584,
				FixType: 2,
			},
		},
		{
			S: "$PGRMZ,2062,f,3*2D",
			Expected: &PGRMZ{
				address: nmea.NewAddress("PGRMZ"),
				AltFeet: 2062,
				FixType: 3,
			},
		},
	})
}
