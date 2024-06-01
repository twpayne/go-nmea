package xctracer_test

import (
	"testing"
	"time"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/nmeatest"
	"github.com/twpayne/go-nmea/xctracer"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(xctracer.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$XCTRC,2015,1,5,16,34,33,36,46.947508,7.453117,540.32,12.35,270.4,2.78,,,,964.93,98*79",
				Expected: &xctracer.XCTRC{
					Address:          nmea.NewAddress("XCTRC"),
					Time:             time.Date(2015, 1, 5, 16, 34, 33, 36e7, time.UTC),
					Lat:              46.947508,
					Lon:              7.453117,
					Alt:              540.32,
					SpeedOverGround:  12.35,
					CourseOverGround: 270.4,
					ClimbRate:        2.78,
					RawPressure:      964.93,
					BatteryPercent:   98,
				},
			},
			{
				S: "$XCTRC,2015,8,11,10,56,23,80,48.62825,8.104885,129.4,0.01,322.76,-0.05,,,,997.79,77*66",
				Expected: &xctracer.XCTRC{
					Address:          nmea.NewAddress("XCTRC"),
					Time:             time.Date(2015, 8, 11, 10, 56, 23, 80e7, time.UTC),
					Lat:              48.62825,
					Lon:              8.104885,
					Alt:              129.4,
					SpeedOverGround:  0.01,
					CourseOverGround: 322.76,
					ClimbRate:        -0.05,
					RawPressure:      997.79,
					BatteryPercent:   77,
				},
			},
		},
	)
}
