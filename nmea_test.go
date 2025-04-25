package nmea_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea"
)

func TestFrameString(t *testing.T) {
	for _, tc := range []struct {
		data     string
		expected string
	}{
		{
			data:     "GPDTM,W84,,0.0,N,0.0,E,0.0,W84",
			expected: "$GPDTM,W84,,0.0,N,0.0,E,0.0,W84*6F\r\n",
		},
	} {
		assert.Equal(t, tc.expected, nmea.FrameString(tc.data))
	}
}
