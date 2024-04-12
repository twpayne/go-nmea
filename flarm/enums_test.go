package flarm_test

import (
	"strconv"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea/flarm"
)

func TestDescribeHex(t *testing.T) {
	for i, tc := range []struct {
		value    int
		m        map[int]string
		expected string
	}{
		{
			value:    0x2,
			m:        flarm.AircraftTypes,
			expected: "0x2 (tow plane/tug plane)",
		},
		{
			value:    0x10,
			m:        flarm.AircraftTypes,
			expected: "0x10",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := flarm.DescribeHex(tc.value, tc.m)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDescribeInt(t *testing.T) {
	for i, tc := range []struct {
		value    int
		m        map[int]string
		expected string
	}{
		{
			value:    0,
			m:        flarm.IDTypes,
			expected: "0 (random ID)",
		},
		{
			value:    -1,
			m:        flarm.IDTypes,
			expected: "-1",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := flarm.DescribeInt(tc.value, tc.m)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
