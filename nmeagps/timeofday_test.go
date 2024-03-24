package nmeagps_test

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea/nmeagps"
)

func TestTimeOfDay(t *testing.T) {
	for _, tc := range []struct {
		timeOfDay             nmeagps.TimeOfDay
		expectedInvalid       bool
		expectedString        string
		expectedSinceMidnight time.Duration
	}{
		{
			timeOfDay:             nmeagps.TimeOfDay{},
			expectedString:        "00:00:00.000000000",
			expectedSinceMidnight: 0,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 456789000,
			},
			expectedString:        "01:02:03.456789000",
			expectedSinceMidnight: 1*time.Hour + 2*time.Minute + 3*time.Second + 456789000*time.Nanosecond,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Hour:       23,
				Minute:     59,
				Second:     59,
				Nanosecond: 999999999,
			},
			expectedString:        "23:59:59.999999999",
			expectedSinceMidnight: 24*time.Hour - time.Nanosecond,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Hour: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Hour: 24,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Minute: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Minute: 60,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Second: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Second: 61,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Nanosecond: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmeagps.TimeOfDay{
				Nanosecond: 1000000000,
			},
			expectedInvalid: true,
		},
	} {
		t.Run(tc.timeOfDay.String(), func(t *testing.T) {
			actualValid := tc.timeOfDay.Valid()
			if tc.expectedInvalid {
				assert.False(t, actualValid)
			} else {
				assert.Equal(t, tc.expectedString, tc.timeOfDay.String())
				assert.Equal(t, tc.expectedSinceMidnight, tc.timeOfDay.SinceMidnight())
			}
		})
	}
}
