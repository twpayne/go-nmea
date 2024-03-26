package gps_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea/gps"
)

func TestAddress(t *testing.T) {
	for _, tc := range []struct {
		address               string
		expectedConstellation byte
		expectedFormatter     string
		expectedTalker        string
		expectedValid         bool
	}{
		{
			address:               "",
			expectedConstellation: 0,
			expectedFormatter:     "",
			expectedTalker:        "",
			expectedValid:         false,
		},
		{
			address:               "GPGGA",
			expectedConstellation: 'P',
			expectedFormatter:     "GGA",
			expectedTalker:        "GP",
			expectedValid:         true,
		},
		{
			address:               "GNRMC",
			expectedConstellation: 'N',
			expectedFormatter:     "RMC",
			expectedTalker:        "GN",
			expectedValid:         true,
		},
		{
			address:               "A",
			expectedConstellation: 0,
			expectedFormatter:     "",
			expectedTalker:        "A",
			expectedValid:         false,
		},
		{
			address:               "ABCDEFG",
			expectedConstellation: 'B',
			expectedFormatter:     "CDEFG",
			expectedTalker:        "AB",
			expectedValid:         false,
		},
	} {
		t.Run(tc.address, func(t *testing.T) {
			address := gps.NewAddress(tc.address)
			assert.Equal(t, tc.expectedConstellation, address.Constellation())
			assert.Equal(t, tc.expectedFormatter, address.Formatter())
			assert.False(t, address.Proprietary())
			assert.Equal(t, tc.expectedTalker, address.Talker())
			assert.Equal(t, tc.expectedValid, address.Valid())
		})
	}
}
