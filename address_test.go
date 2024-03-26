package nmea

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestAddress(t *testing.T) {
	for _, tc := range []struct {
		address             string
		expectedFormatter   string
		expectedProprietary bool
		expectedTalker      string
	}{
		{
			address:             "",
			expectedFormatter:   "",
			expectedProprietary: false,
			expectedTalker:      "",
		},
		{
			address:             "GPGGA",
			expectedFormatter:   "GGA",
			expectedProprietary: false,
			expectedTalker:      "GP",
		},
		{
			address:             "PGRMA",
			expectedFormatter:   "PGRMA",
			expectedProprietary: true,
			expectedTalker:      "PGRM",
		},
	} {
		t.Run(tc.address, func(t *testing.T) {
			address := NewAddress(tc.address)
			assert.Equal(t, tc.expectedProprietary, address.Proprietary())
			assert.Equal(t, tc.expectedTalker, address.Talker())
		})
	}
}
