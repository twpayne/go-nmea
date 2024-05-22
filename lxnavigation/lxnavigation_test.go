package lxnavigation_test

import (
	"testing"

	"github.com/twpayne/go-nmea"
	"github.com/twpayne/go-nmea/lxnavigation"
	"github.com/twpayne/go-nmea/nmeatest"
)

func TestSentenceParserFunc(t *testing.T) {
	nmeatest.TestSentenceParserFunc(t,
		[]nmea.ParserOption{
			nmea.WithChecksumDiscipline(nmea.ChecksumDisciplineStrict),
			nmea.WithLineEndingDiscipline(nmea.LineEndingDisciplineNever),
			nmea.WithSentenceParserFunc(lxnavigation.SentenceParserFunc),
		},
		[]nmeatest.TestCase{
			{
				S: "$LXWP0,Y,119.4,1717.6,0.02,0.02,0.02,0.02,0.02,0.02,,000,107.2*5b",
				Expected: &lxnavigation.LXWP0{
					Address:         nmea.NewAddress("LXWP0"),
					IsLoggerRunning: true,
					TrueAirspeedKPH: nmea.NewOptional(119.4),
					Alt:             nmea.NewOptional(1717.6),
					Vario: []float64{
						0.02,
						0.02,
						0.02,
						0.02,
						0.02,
						0.02,
					},
					WindDirection: nmea.NewOptional(0.0),
					WindSpeedKMH:  nmea.NewOptional(107.2),
				},
			},
			{
				S: "$LXWP0,N,,405.74,0.01,,,,,,218,,*55",
				Expected: &lxnavigation.LXWP0{
					Address:         nmea.NewAddress("LXWP0"),
					IsLoggerRunning: false,
					Alt:             nmea.NewOptional(405.74),
					Vario: []float64{
						0.01,
					},
					Heading: nmea.NewOptional(218),
				},
			},
			{
				S: "$LXWP0,Y,222.3,1665.5,1.71,,,,,,239,174,10.1*47",
				Expected: &lxnavigation.LXWP0{
					Address:         nmea.NewAddress("LXWP0"),
					IsLoggerRunning: true,
					TrueAirspeedKPH: nmea.NewOptional(222.3),
					Alt:             nmea.NewOptional(1665.5),
					Vario: []float64{
						1.71,
					},
					Heading:       nmea.NewOptional(239),
					WindDirection: nmea.NewOptional(174.0),
					WindSpeedKMH:  nmea.NewOptional(10.1),
				},
			},
			{
				S: "$LXWP0,Y,185.4,12149.5,-1.47,,,,,,263,267,145.8*6B",
				Expected: &lxnavigation.LXWP0{
					Address:         nmea.NewAddress("LXWP0"),
					IsLoggerRunning: true,
					TrueAirspeedKPH: nmea.NewOptional(185.4),
					Alt:             nmea.NewOptional(12149.5),
					Vario: []float64{
						-1.47,
					},
					Heading:       nmea.NewOptional(263),
					WindDirection: nmea.NewOptional(267.0),
					WindSpeedKMH:  nmea.NewOptional(145.8),
				},
			},
			{
				S: "$LXWP1,LX Eos,34949,1.5,1.4*7d",
				Expected: &lxnavigation.LXWP1{
					Address:         nmea.NewAddress("LXWP1"),
					DeviceName:      "LX Eos",
					SerialNumber:    34949,
					SoftwareVersion: "1.5",
					HardwareVersion: "1.4",
				},
			},
			{
				S: "$LXWP2,1.5,1.11,13,2.96,-3.03,1.35,45*02",
				Expected: &lxnavigation.LXWP2{
					Address:         nmea.NewAddress("LXWP2"),
					MacCreadyFactor: 1.5,
					LoadFactor:      1.11,
					BugsPercent:     13,
					PolarA:          2.96,
					PolarB:          -3.03,
					PolarC:          1.35,
					VolumePercent:   45,
				},
			},
			{
				S: "$LXWP3,47.76,0,2.0,5.0,15,30,2.5,1.0,0,100,0.1,,0*08",
				Expected: &lxnavigation.LXWP3{
					Address:         nmea.NewAddress("LXWP3"),
					AltOffset:       47.76,
					Mode:            0,
					Filter:          2.0,
					TELevel:         15,
					IntegrationTime: 30,
					Range:           2.5,
					Silence:         1,
					SwitchMode:      0,
					Speed:           100,
					PolarName:       "",
				},
			},
			{
				S: "$LXWP3,0,2,5.0,0,29,20,10.0,1.3,1,120,0,KA6e,0*74",
				Expected: &lxnavigation.LXWP3{
					Address:         nmea.NewAddress("LXWP3"),
					AltOffset:       0,
					Mode:            2,
					Filter:          5.0,
					TELevel:         29,
					IntegrationTime: 20,
					Range:           10.0,
					Silence:         1.3,
					SwitchMode:      1,
					Speed:           120,
					Unknown:         struct{}{},
					PolarName:       "KA6e",
				},
			},
			/*
				{
					S: "$GPRMB,A,0.00,R,,CELJE,4614.367,N,01513.482,E,1.7,273.8,0.0,A*7f",
					Expected: &standard.RMB{
						Address: nmea.NewAddress("GPRMB"),
					},
				},
			*/
		},
	)
}
