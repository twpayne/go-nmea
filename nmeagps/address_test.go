package nmeagps_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea/nmeagps"
)

func TestAddress(t *testing.T) {
	for _, tc := range []struct {
		address               nmeagps.Address
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
		t.Run(tc.address.String(), func(t *testing.T) {
			assert.Equal(t, tc.expectedConstellation, tc.address.Constellation())
			assert.Equal(t, tc.expectedFormatter, tc.address.Formatter())
			assert.False(t, tc.address.Proprietary())
			assert.Equal(t, tc.expectedTalker, tc.address.Talker())
			assert.Equal(t, tc.expectedValid, tc.address.Valid())
		})
	}
}
