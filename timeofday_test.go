package nmea_test

import (
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-nmea"
)

func TestParseTimeOfDay(t *testing.T) {
	for _, tc := range []struct {
		s        string
		expected nmea.TimeOfDay
	}{
		{
			s: "010203",
			expected: nmea.TimeOfDay{
				Hour:   1,
				Minute: 2,
				Second: 3,
			},
		},
		{
			s: "010203.",
			expected: nmea.TimeOfDay{
				Hour:   1,
				Minute: 2,
				Second: 3,
			},
		},
		{
			s: "010203.4",
			expected: nmea.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 400000000,
			},
		},
		{
			s: "010203.456789123",
			expected: nmea.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 456789123,
			},
		},
		{
			s: "010203.4567891234",
			expected: nmea.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 456789123,
			},
		},
		{
			s: "010203.4567891235",
			expected: nmea.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 456789124,
			},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			tok := nmea.NewTokenizer([]byte(tc.s))
			assert.Equal(t, tc.expected, nmea.ParseTimeOfDay(tok))
			assert.NoError(t, tok.Err())
		})
	}
}

func TestTimeOfDay(t *testing.T) {
	for _, tc := range []struct {
		timeOfDay             nmea.TimeOfDay
		expectedInvalid       bool
		expectedString        string
		expectedSinceMidnight time.Duration
	}{
		{
			timeOfDay:             nmea.TimeOfDay{},
			expectedString:        "00:00:00.000000000",
			expectedSinceMidnight: 0,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Hour:       1,
				Minute:     2,
				Second:     3,
				Nanosecond: 456789000,
			},
			expectedString:        "01:02:03.456789000",
			expectedSinceMidnight: 1*time.Hour + 2*time.Minute + 3*time.Second + 456789000*time.Nanosecond,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Hour:       23,
				Minute:     59,
				Second:     59,
				Nanosecond: 999999999,
			},
			expectedString:        "23:59:59.999999999",
			expectedSinceMidnight: 24*time.Hour - time.Nanosecond,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Hour: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Hour: 24,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Minute: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Minute: 60,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Second: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Second: 61,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
				Nanosecond: -1,
			},
			expectedInvalid: true,
		},
		{
			timeOfDay: nmea.TimeOfDay{
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
